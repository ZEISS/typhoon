

package googlecloudpubsubtarget

import (
	"strconv"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
	"github.com/zeiss/typhoon/pkg/targets/reconciler"
)

const (
	envEventsPayloadPolicy = "EVENTS_PAYLOAD_POLICY"
	envDiscardCEContext    = "DISCARD_CE_CONTEXT"
)

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
type adapterConfig struct {
	/github.com/zeiss/typhoonng/metrics/tracing
	obsConfigithub.com/zeiss/typhoon
	/github.com/zeiss/typhoon
	Igithub.com/zeiss/typhoonermesh/googlecloudpubsubtarget-adapter"`
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*servingv1.Service] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, _ *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.GoogleCloudPubSubTarget)

	return common.NewAdapterKnService(trg, nil,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
	), nil
}

func MakeAppEnv(o *v1alpha1.GoogleCloudPubSubTarget) []corev1.EnvVar {
	env := []corev1.EnvVar{
		{
			Name:  "GCLOUD_PUBSUB_TOPIC",
			Value: o.Spec.Topic.String(),
		}, {
			Name:  envDiscardCEContext,
			Value: strconv.FormatBool(o.Spec.DiscardCloudEventContext),
		},
	}

	env = append(env, reconciler.MakeGCPAuthEnvVars(o.Spec.ServiceAccountKey, o.Spec.Auth)...)

	if o.Spec.EventOptions != nil && o.Spec.EventOptions.PayloadPolicy != nil {
		env = append(env, corev1.EnvVar{
			Name:  envEventsPayloadPolicy,
			Value: string(*o.Spec.EventOptions.PayloadPolicy),
		})
	}

	return env
}
