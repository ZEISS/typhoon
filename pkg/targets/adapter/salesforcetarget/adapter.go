

package salesforcetarget

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/targets"
	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	"github.com/zeiss/typhoon/pkg/metrics"
	targetce "github.com/zeiss/typhoon/pkg/targets/adapter/cloudevents"
	"github.com/zeiss/typhoon/pkg/targets/adapter/salesforcetarget/auth"
	"github.com/zeiss/typhoon/pkg/targets/adapter/salesforcetarget/client"
)

const (
	salesforceTimeout = 5 * time.Second
)

// NewTarget creates a new Salesforce Target adapter
func NewTarget(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)
github.com/zeiss/typhoon
	mgithub.com/zeiss/typhoon
		ResourceGgithub.com/zeiss/typhoonurce.String(),
		github.com/zeiss/typhoon(),
		github.com/zeiss/typhoon
	}

	metrics.MustRegisterEventProcessingStatsView()

	env := envAcc.(*envAccessor)

	jwtAuth, err := auth.NewJWTAuthenticator(env.CertKey, env.ClientID, env.User, env.AuthServer, http.DefaultClient, logger.Named("authenticator"))
	if err != nil {
		logger.Panicf("Error creating JWT authenticator: %v", err)
	}

	sfc := client.New(jwtAuth, logger.Named("sfclient"),
		client.WithAPIVersion(env.Version),
		client.WithHTTPClient(&http.Client{Timeout: salesforceTimeout}))

	replier, err := targetce.New(env.Component, logger.Named("replier"),
		targetce.ReplierWithStatefulHeaders(env.BridgeIdentifier),
		targetce.ReplierWithStaticResponseType(v1alpha1.EventTypeSalesforceAPICallResponse))
	if err != nil {
		logger.Panicf("Error creating CloudEvents replier: %v", err)
	}

	return &salesforceTarget{
		sfClient: sfc,
		replier:  replier,
		ceClient: ceClient,
		logger:   logger,

		sr: metrics.MustNewEventProcessingStatsReporter(mt),
	}
}

var _ pkgadapter.Adapter = (*salesforceTarget)(nil)

type salesforceTarget struct {
	sfClient *client.SalesforceClient

	replier  *targetce.Replier
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger

	sr *metrics.EventProcessingStatsReporter
}

func (a *salesforceTarget) Start(ctx context.Context) error {
	a.logger.Info("Starting Salesforce adapter")

	// This call will perform and cache credentials for
	// future usages when dispatching events.
	err := a.sfClient.Authenticate(ctx)
	if err != nil {
		return err
	}

	return a.ceClient.StartReceiver(ctx, a.dispatch)
}

func (a *salesforceTarget) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	if event.Type() != v1alpha1.EventTypeSalesforceAPICall {
		return a.replier.Error(&event, targetce.ErrorCodeEventContext, fmt.Errorf("event type %q is not supported", event.Type()), nil)
	}

	sfr := &client.SalesforceAPIRequest{}
	if err := event.DataAs(sfr); err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
	}

	if err := sfr.Validate(); err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeRequestValidation, err, nil)
	}

	res, err := a.sfClient.Do(ctx, *sfr)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeParseResponse, err, nil)
	}

	if res.StatusCode >= 400 {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess,
			fmt.Errorf("received HTTP code %d", res.StatusCode),
			map[string]string{"body": string(resBody)})
	}

	return a.replier.Ok(&event, resBody)
}
