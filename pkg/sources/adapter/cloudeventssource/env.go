package cloudeventssource

import (
	"encoding/json"

	cereconciler "github.com/zeiss/typhoon/pkg/sources/reconciler/cloudeventssource"
	"knative.dev/eventing/pkg/adapter/v2"
)

// NewEnvConfig satisfies pkgadapter.EnvConfigConstructor.
func NewEnvConfig() adapter.EnvConfigAccessor {
	return &envAccessor{}
}

// KeyMountedValues contains a set of file mounted values
// by their name.
type KeyMountedValues []cereconciler.KeyMountedValue

// Decode an array of KeyMountedValues
func (is *KeyMountedValues) Decode(value string) error {
	if err := json.Unmarshal([]byte(value), is); err != nil {
		return err
	}
	return nil
}

type envAccessor struct {
	adapter.EnvConfig

	Path              string           `envconfig:"CLOUDEVENTS_PATH"`
	BasicAuths        KeyMountedValues `envconfig:"CLOUDEVENTS_BASICAUTH_CREDENTIALS"`
	RequestsPerSecond uint64           `envconfig:"CLOUDEVENTS_RATELIMITER_RPS"`
}
