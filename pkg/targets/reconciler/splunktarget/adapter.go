/*
Copyright 2022 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package splunktarget

import (
	"strconv"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	commonv1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/common/v1alpha1"
	"github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	common "github.com/triggermesh/triggermesh/pkg/reconciler"
	"github.com/triggermesh/triggermesh/pkg/reconciler/resource"
)

const (
	envHECEndpoint   = "SPLUNK_HEC_ENDPOINT"
	envHECToken      = "SPLUNK_HEC_TOKEN"
	envIndex         = "SPLUNK_INDEX"
	envSkipTLSVerify = "SPLUNK_SKIP_TLS_VERIFY"
)

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
type adapterConfig struct {
	// Configuration accessor for logging/metrics/tracing
	obsConfig source.ConfigAccessor
	// Container image
	Image string `default:"gcr.io/triggermesh/splunktarget-adapter"`
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*servingv1.Service] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, _ *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.SplunkTarget)

	return common.NewAdapterKnService(trg, nil,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.SplunkTarget) []corev1.EnvVar {
	hecURL := apis.URL{
		Scheme: o.Spec.Endpoint.Scheme,
		Host:   o.Spec.Endpoint.Host,
	}

	env := []corev1.EnvVar{
		{
			Name:  envHECEndpoint,
			Value: hecURL.String(),
		},
		{
			Name:  "DISCARD_CE_CONTEXT",
			Value: strconv.FormatBool(o.Spec.DiscardCEContext),
		},
	}

	env = common.MaybeAppendValueFromEnvVar(env, envHECToken, o.Spec.Token)

	if idx := o.Spec.Index; idx != nil && *idx != "" {
		env = append(env, corev1.EnvVar{
			Name:  envIndex,
			Value: *idx,
		})
	}

	if o.Spec.SkipTLSVerify != nil {
		env = append(env, corev1.EnvVar{
			Name:  envSkipTLSVerify,
			Value: strconv.FormatBool(*o.Spec.SkipTLSVerify),
		})
	}

	return env
}
