

package zendesktarget

import (
	"context"
	"testing"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	rt "knative.dev/pkg/reconciler/testing"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	fakeinjectionclient "github.com/zeiss/typhoon/pkg/client/generated/injection/client/fake"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/targets/v1alpha1/zendesktarget"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	. "github.com/zeiss/typhoon/pkg/reconciler/testing"
)

func TestReconcile(t *testing.T) {
	adapterCfg := &adapterConfig{
		Image:     "registry/image:tag",
		obsConfig: &source.EmptyVarsGenerator{},
	}

	ctor := reconcilerCtor(adapterCfg)
	trg := newTarget()
	ab := adapterBuilder(github.com/zeiss/typhoon
github.com/zeiss/typhoon
	TestRecogithub.com/zeiss/typhoon
}github.com/zeiss/typhoon

// reconcilerCtor returns a Ctor for a ZendeskTarget Reconciler.
func reconcilerCtor(cfg *adapterConfig) Ctor {
	return func(t *testing.T, ctx context.Context, _ *rt.TableRow, ls *Listers) controller.Reconciler {
		r := &Reconciler{
			adapterCfg: cfg,
		}

		r.base = NewTestServiceReconciler[*v1alpha1.ZendeskTarget](ctx, ls,
			ls.GetZendeskTargetLister().ZendeskTargets,
		)

		return reconcilerv1alpha1.NewReconciler(ctx, logging.FromContext(ctx),
			fakeinjectionclient.Get(ctx), ls.GetZendeskTargetLister(),
			controller.GetEventRecorder(ctx), r)
	}
}

// newTarget returns a populated target object.
func newTarget() *v1alpha1.ZendeskTarget {
	trg := &v1alpha1.ZendeskTarget{
		Spec: v1alpha1.ZendeskTargetSpec{
			Subject:   "tmTickets0",
			Subdomain: "triggermesh",
			Email:     "woodford@example.com",
		},
	}

	Populate(trg)

	return trg
}

// adapterBuilder returns a slim Reconciler containing only the fields accessed
// by r.BuildAdapter().
func adapterBuilder(cfg *adapterConfig) common.AdapterBuilder[*servingv1.Service] {
	return &Reconciler{
		adapterCfg: cfg,
	}
}
