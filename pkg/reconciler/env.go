package reconciler

// Common environment variables propagated to adapters.
const (
	EnvName      = "NAME"
	EnvNamespace = "NAMESPACE"

	envSink                  = "K_SINK"
	envComponent             = "K_COMPONENT"
	envSinkTimeout           = "K_SINK_TIMEOUT"
	envMetricsPrometheusPort = "METRICS_PROMETHEUS_PORT"

	// Overrides for CloudEvents context attributes (only supported by a subset of components)
	EnvCESource = "CE_SOURCE"
	EnvCEType   = "CE_TYPE"
)
