package cloudeventssource

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/adapter/fs"
	"github.com/zeiss/typhoon/pkg/apis/sources"
	"github.com/zeiss/typhoon/pkg/sources/adapter/cloudeventssource/ratelimiter"
)

// NewAdapter satisfies pkgadapter.AdapterConstructor.
func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: sources.CloudEventsSourceResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	env := envAcc.(*envAccessor)

	cfw, err := fs.NewCachedFileWatcher(logger)
	if err != nil {
		logger.Panic("Could not create a file watcher", zap.Error(err))
	}

	for _, as := range env.BasicAuths {
		if err := cfw.Add(as.MountedValueFile); err != nil {
			logger.Panic(
				fmt.Sprintf("Authentication secret at %q could not be watched", as.MountedValueFile),
				zap.Error(err))
		}
	}

	ceh := &cloudEventsHandler{
		basicAuths: env.BasicAuths,

		cfw:      cfw,
		ceClient: ceClient,
		logger:   logger,
		mt:       mt,
	}

	// prepare CE server options
	options := []cehttp.Option{}

	if env.Path != "" {
		options = append(options, cehttp.WithPath(env.Path))
	}
	if len(env.BasicAuths) != 0 {
		options = append(options, cehttp.WithMiddleware(ceh.handleAuthentication))
	}

	if env.RequestsPerSecond != 0 {
		rl, err := ratelimiter.New(env.RequestsPerSecond)
		if err != nil {
			logger.Panic("Could not create rate limiter", zap.Error(err))
		}
		options = append(options, cehttp.WithRateLimiter(rl))
	}

	ceServer, err := cloudevents.NewClientHTTP(options...)
	if err != nil {
		logger.Panic("Error creating CloudEvents client", zap.Error(err))
	}

	ceh.ceServer = ceServer
	return ceh
}

var _ pkgadapter.Adapter = (*cloudEventsHandler)(nil)
