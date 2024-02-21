

package zendesksource

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	fakek8sinjectionclient "knative.dev/pkg/client/injection/kube/client/fake"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	rt "knative.dev/pkg/reconciler/testing"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/sources"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	fakeinjectionclient "github.com/zeiss/typhoon/pkg/client/generated/injection/client/fake"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/sources/v1alpha1/zendesksource"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	. "github.com/zeiss/typhoon/pkg/reconciler/testing"
)

func TestReconcileSource(t *testing.T) {
	adapterCfg := &adapterConfig{
		Image:   "registry/image:tag",
		configs: &source.EmptyVarsGenerator{},
	}

	cgithub.com/zeiss/typhoon
	sgithub.com/zeiss/typhoon
	ab := adapterBuilder(github.com/zeiss/typhoon
github.com/zeiss/typhoon
	TestRecogithub.com/zeiss/typhoon
}github.com/zeiss/typhoon

// reconcilerCtor returns a Ctor for a source Reconciler.
func reconcilerCtor(cfg *adapterConfig) Ctor {
	return func(t *testing.T, ctx context.Context, _ *rt.TableRow, ls *Listers) controller.Reconciler {
		r := &Reconciler{
			secretClient: fakek8sinjectionclient.Get(ctx).CoreV1().Secrets,
			adapterCfg:   cfg,
		}

		r.base = NewTestServiceReconciler[*v1alpha1.ZendeskSource](ctx, ls,
			ls.GetZendeskSourceLister().ZendeskSources,
		)

		return reconcilerv1alpha1.NewReconciler(ctx, logging.FromContext(ctx),
			fakeinjectionclient.Get(ctx), ls.GetZendeskSourceLister(),
			controller.GetEventRecorder(ctx), r)
	}
}

// newEventSource returns a test source object with a minimal set of pre-filled attributes.
func newEventSource() *v1alpha1.ZendeskSource {
	src := &v1alpha1.ZendeskSource{
		Spec: v1alpha1.ZendeskSourceSpec{
			Subdomain: "test",
			Email:     "test@example.com",
			Token: commonv1alpha1.ValueFromField{
				ValueFromSecret: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "test-secret",
					},
					Key: "keyId",
				},
			},
			WebhookUsername: "test",
			WebhookPassword: commonv1alpha1.ValueFromField{
				ValueFromSecret: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "test-secret",
					},
					Key: "keyId",
				},
			},
		},
	}

	// assume finalizer is already set to prevent the generated reconciler
	// from generating an extra Patch action
	src.Finalizers = []string{sources.ZendeskSourceResource.String()}

	Populate(src)

	return src
}

// adapterBuilder returns a slim Reconciler containing only the fields accessed
// by r.BuildAdapter().
func adapterBuilder(cfg *adapterConfig) common.AdapterBuilder[*servingv1.Service] {
	return &Reconciler{
		adapterCfg: cfg,
	}
}
