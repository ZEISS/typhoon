

package oracletarget

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
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/targets/v1alpha1/oracletarget"
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

// reconcilerCtor returns a Ctor for a OracleTarget Reconciler.
func reconcilerCtor(cfg *adapterConfig) Ctor {
	return func(t *testing.T, ctx context.Context, _ *rt.TableRow, ls *Listers) controller.Reconciler {
		r := &Reconciler{
			adapterCfg: cfg,
		}

		r.base = NewTestServiceReconciler[*v1alpha1.OracleTarget](ctx, ls,
			ls.GetOracleTargetLister().OracleTargets,
		)

		return reconcilerv1alpha1.NewReconciler(ctx, logging.FromContext(ctx),
			fakeinjectionclient.Get(ctx), ls.GetOracleTargetLister(),
			controller.GetEventRecorder(ctx), r)
	}
}

// newTarget returns a populated target object.
func newTarget() *v1alpha1.OracleTarget {
	trg := &v1alpha1.OracleTarget{
		Spec: v1alpha1.OracleTargetSpec{
			Tenancy: "ocid1.tenancy.oc1..aaaaaaaaav23f45mqyxmwu4x3s2uhuh4rb2bwdpgb5kbpjqvwiiqufhsq6za",
			User:    "ocid1.user.oc1..aaaaaaaacaxtveoy4zx7rsg7lanexmouxjxay6godthrfsocpl6ggrfpbiuq",
			Region:  "us-phoenix-1",
			OracleFunctionSpec: &v1alpha1.OracleFunctionSpecSpec{
				Function: "ocid1.fnfunc.oc1.phx.aaaaaaaaaajrgy4on66e6krko73h2im5qaiiagecg5hmbcqib2kpbzlcy3bq",
			},
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
