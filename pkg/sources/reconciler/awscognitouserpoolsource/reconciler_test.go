

package awscognitouserpoolsource

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	appsv1 "k8s.io/api/apps/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	rt "knative.dev/pkg/reconciler/testing"

	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	fakeinjectionclient "github.com/zeiss/typhoon/pkg/client/generated/injection/client/fake"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/sources/v1alpha1/awscognitouserpoolsource"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	. "github.com/zeiss/typhoon/pkg/reconciler/testing"
)

func TestReconcileSource(t *testing.T) {
	adapterCfg := &adapterConfig{
		Image:   "registry/image:tag",
		configs: &source.EmptyVarsGenerator{},
	}

	ctor := reconcilerCtor(adapterCfg)
	src := newEventSource()
	ab := adapterBuilder(github.com/zeiss/typhoon
github.com/zeiss/typhoon
	TestRecogithub.com/zeiss/typhoon
}github.com/zeiss/typhoon

// reconcilerCtor returns a Ctor for a AWSCognitoUserPoolSource Reconciler.
func reconcilerCtor(cfg *adapterConfig) Ctor {
	return func(t *testing.T, ctx context.Context, _ *rt.TableRow, ls *Listers) controller.Reconciler {
		r := &Reconciler{
			adapterCfg: cfg,
		}

		r.base = NewTestDeploymentReconciler[*v1alpha1.AWSCognitoUserPoolSource](ctx, ls,
			ls.GetAWSCognitoUserPoolSourceLister().AWSCognitoUserPoolSources,
		)

		return reconcilerv1alpha1.NewReconciler(ctx, logging.FromContext(ctx),
			fakeinjectionclient.Get(ctx), ls.GetAWSCognitoUserPoolSourceLister(),
			controller.GetEventRecorder(ctx), r)
	}
}

// newEventSource returns a test source object with a minimal set of pre-filled attributes.
func newEventSource(skipCEAtrributes ...interface{}) *v1alpha1.AWSCognitoUserPoolSource {
	src := &v1alpha1.AWSCognitoUserPoolSource{
		Spec: v1alpha1.AWSCognitoUserPoolSourceSpec{
			ARN: NewARN(cognitoidentityprovider.ServiceName, "userpool/triggermeshtest"),
		},
	}

	Populate(src)

	return src
}

// adapterBuilder returns a slim Reconciler containing only the fields accessed
// by r.BuildAdapter().
func adapterBuilder(cfg *adapterConfig) common.AdapterBuilder[*appsv1.Deployment] {
	return &Reconciler{
		adapterCfg: cfg,
	}
}
