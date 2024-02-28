package httppollersource

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/sources"
)

// NewAdapter satisfies pkgadapter.AdapterConstructor.
func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: sources.HTTPPollerSourceResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	env := envAcc.(*envAccessor)

	t := &http.Transport{
		TLSClientConfig: &tls.Config{ // #nosec G402
			InsecureSkipVerify: env.SkipVerify,
		},
	}

	if env.CACertificate != "" {
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM([]byte(env.CACertificate)) {
			logger.Panicf("Failed adding certificate to pool: %s", env.CACertificate)
		}

		t.TLSClientConfig = &tls.Config{
			RootCAs:    certPool,
			MinVersion: tls.VersionTLS12,
		}
	}

	httpClient := &http.Client{Transport: t}

	httpRequest, err := http.NewRequest(env.Method, env.Endpoint, nil)
	if err != nil {
		logger.Panicw("Cannot build request", zap.Error(err))
	}

	for k, v := range env.Headers {
		httpRequest.Header.Set(k, v)
	}

	if env.BasicAuthUsername != "" || env.BasicAuthPassword != "" {
		httpRequest.SetBasicAuth(env.BasicAuthUsername, env.BasicAuthPassword)
	}

	return &httpPoller{
		eventType:   env.EventType,
		eventSource: env.EventSource,
		interval:    env.Interval,

		httpClient:  httpClient,
		httpRequest: httpRequest,

		ceClient: ceClient,
		logger:   logger,
		mt:       mt,
	}
}
