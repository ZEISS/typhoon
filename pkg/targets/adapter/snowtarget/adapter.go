package snowtarget

import (
	"context"

	"github.com/zeiss/snow-go/push"
	"github.com/zeiss/typhoon/pkg/apis/targets"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	snowgo "github.com/zeiss/snow-go"
	"go.uber.org/zap"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	instance string `envconfig:"SNOW_INSTANCE"`
	user     string `envconfig:"SNOW_BASICAUTH_USER"`
	password string `envconfig:"SNOW_BASICAUTH_PASSWORD"`
	source   string `envconfig:"SNOW_SOURCE"`
}

var _ pkgadapter.Adapter = (*snowAdapter)(nil)

type snowAdapter struct {
	client   cloudevents.Client
	instance string
	sc       *snowgo.Client

	logger *zap.SugaredLogger
	mt     *pkgadapter.MetricTag
}

// NewTarget adapter implementation
func NewTarget(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, client cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: targets.KafkaTargetResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	env := envAcc.(*envAccessor)

	basicAuth, err := snowgo.NewBasicAuth(env.user, env.password)
	if err != nil {
		logger.Panicw("failed to create authenticator", zap.Error(err))
	}
	sc := snowgo.New(env.instance, snowgo.WithRequestEditorFn(basicAuth.Intercept))

	return &snowAdapter{
		mt:       mt,
		sc:       sc,
		instance: env.instance,
	}
}

// Start is the main entrypoint for the adapter
func (a *snowAdapter) Start(ctx context.Context) error {
	a.logger.Info("starting SNOW adapter")

	return a.client.StartReceiver(ctx, a.dispatch)
}

func (a *snowAdapter) dispatch(event cloudevents.Event) cloudevents.Result {
	url := push.NewPushConnectorUrl(a.instance, "typhoon")

	req := push.NewRequest(url, event)
	res := &push.Response{}

	err := a.sc.Do(context.Background(), req, res)
	if err != nil {
		return err
	}

	return cloudevents.ResultACK
}

func (a *snowAdapter) errorHTTPResult(statusCode int, message string, args ...interface{}) cloudevents.Result {
	r := cloudevents.NewHTTPResult(statusCode, message, args...)
	a.logger.Error(r.Error())

	return r
}
