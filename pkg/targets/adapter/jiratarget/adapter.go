package jiratarget

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/google/uuid"
	"go.uber.org/zap"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/andygrunwald/go-jira"

	"github.com/zeiss/typhoon/pkg/apis/targets"
	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	"github.com/zeiss/typhoon/pkg/metrics"
)

// Method performed on the API URI
type Method string

// Available actions at the JIRA API
const (
	MethodCreate Method = http.MethodPost
	MethodPut    Method = http.MethodPut
	MethodPatch  Method = http.MethodPatch
	MethodGet    Method = http.MethodGet
	MethodDelete Method = http.MethodDelete
)

// JiraAPIRequest contains common parameters used for
// interacting with Jira using the API.
type JiraAPIRequest struct {
	Method  Method            `json:"method"`
	Path    string            `json:"path"`
	Query   map[string]string `json:"query"`
	Payload json.RawMessage   `json:"payload"`
}

// IssueGetRequest contains parameters for issue retrieval
type IssueGetRequest struct {
	ID      string               `json:"id"`
	Options jira.GetQueryOptions `json:"options"`
}

// NewTarget creates a Jira target adapter
func NewTarget(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: targets.JiraTargetResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	metrics.MustRegisterEventProcessingStatsView()

	env := envAcc.(*envAccessor)

	basicAuth := jira.BasicAuthTransport{
		Username: env.JiraBasicAuthUser,
		Password: env.JiraBasicAuthToken,
	}

	jiraClient, err := jira.NewClient(basicAuth.Client(), env.JiraURL)
	if err != nil {
		logger.Panicw("Could not create the Jira client", zap.Error(err))
	}

	return &jiraAdapter{
		ceClient: ceClient,
		logger:   logger,

		jiraClient: jiraClient,
		baseURL:    env.JiraURL,
		resSource:  env.Namespace + "/" + env.Name + ": " + env.JiraURL,

		sr: metrics.MustNewEventProcessingStatsReporter(mt),
	}
}

type jiraAdapter struct {
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger

	baseURL    string
	jiraClient *jira.Client
	resSource  string

	sr *metrics.EventProcessingStatsReporter
}

var _ pkgadapter.Adapter = (*jiraAdapter)(nil)

func (a *jiraAdapter) Start(ctx context.Context) error {
	a.logger.Info("Starting Jira adapter")

	if err := a.ceClient.StartReceiver(ctx, a.dispatch); err != nil {
		return err
	}
	return nil
}

func (a *jiraAdapter) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	switch event.Type() {
	case v1alpha1.EventTypeJiraIssueCreate:
		return a.jiraIssueCreate(ctx, event)
	case v1alpha1.EventTypeJiraIssueGet:
		return a.jiraIssueGet(ctx, event)
	case v1alpha1.EventTypeJiraCustom:
		return a.jiraCustomRequest(ctx, event)
	}

	a.logger.Errorf("Event type %q is not supported", event.Type())
	return nil, cloudevents.ResultNACK
}

func (a *jiraAdapter) jiraIssueCreate(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	j := &jira.Issue{}
	if err := event.DataAs(j); err != nil {
		a.logger.Error("Error processing incoming event data as Jira Issue", zap.Error(err))
		return nil, cloudevents.ResultACK
	}

	issue, res, err := a.jiraClient.Issue.CreateWithContext(ctx, j)
	if err != nil {
		respErr := jira.NewJiraError(res, err)
		a.logger.Error("Error requesting Jira API", zap.Error(respErr))
		return nil, cloudevents.ResultACK
	}

	out := cloudevents.NewEvent()
	if err := out.SetData(cloudevents.ApplicationJSON, issue); err != nil {
		a.logger.Error("Error generating response event", zap.Error(err))
		return nil, cloudevents.ResultACK
	}

	out.SetID(uuid.New().String())
	out.SetType(v1alpha1.EventTypeJiraIssue)
	out.SetSource(a.resSource)

	return &out, cloudevents.ResultACK
}

func (a *jiraAdapter) jiraIssueGet(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	j := &IssueGetRequest{}
	if err := event.DataAs(j); err != nil {
		a.logger.Error("Error processing incoming event data as IssueGetRequest", zap.Error(err))
		return nil, cloudevents.ResultACK
	}

	issue, res, err := a.jiraClient.Issue.GetWithContext(ctx, j.ID, &j.Options)
	if err != nil {
		respErr := jira.NewJiraError(res, err)
		a.logger.Error("Error requesting Jira API", zap.Error(respErr))
		return nil, cloudevents.ResultACK
	}

	out := cloudevents.NewEvent()
	if err := out.SetData(cloudevents.ApplicationJSON, issue); err != nil {
		a.logger.Error("Error generating response event", zap.Error(err))
		return nil, cloudevents.ResultACK
	}

	out.SetID(uuid.New().String())
	out.SetType(v1alpha1.EventTypeJiraIssue)
	out.SetSource(a.resSource)

	return &out, cloudevents.ResultACK
}

func (a *jiraAdapter) jiraCustomRequest(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	j := &JiraAPIRequest{}
	if err := event.DataAs(j); err != nil {
		a.logger.Error("Error processing incoming event data as generic Jira API request", zap.Error(err))
		return nil, cloudevents.ResultACK
	}

	u, err := url.Parse(a.baseURL)
	if err != nil {
		a.logger.Error("Error parsing base URL", zap.Error(err))
		return nil, cloudevents.ResultACK
	}
	u.Path = path.Join(u.Path, j.Path)

	if len(j.Query) > 0 {
		q := url.Values{}
		for k, v := range j.Query {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}

	req, err := a.jiraClient.NewRequestWithContext(ctx, string(j.Method), u.String(), j.Payload)
	if err != nil {
		a.logger.Error("Error creating request", zap.Error(err))
		return nil, cloudevents.ResultACK
	}

	res, err := a.jiraClient.Do(req, nil)
	if err != nil {
		respErr := jira.NewJiraError(res, err)
		a.logger.Error("Error requesting Jira API", zap.Error(respErr))
		return nil, cloudevents.ResultACK
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		a.logger.Error("Error reading response from Jira API", zap.Error(err))
		return nil, cloudevents.ResultACK
	}

	out := cloudevents.NewEvent()
	if err := out.SetData(cloudevents.ApplicationJSON, resBody); err != nil {
		a.logger.Error("Error generating response event", zap.Error(err))
		return nil, cloudevents.ResultACK
	}

	out.SetID(uuid.New().String())
	out.SetType(event.Type() + ".response")
	out.SetSource(a.resSource)

	return &out, cloudevents.ResultACK
}
