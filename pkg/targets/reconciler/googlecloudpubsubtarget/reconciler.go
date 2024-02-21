

package googlecloudpubsubtarget

import (
	"context"

	"knative.dev/pkg/reconciler"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/targets/v1alpha1/googlecloudpubsubtarget"
	listersv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/targets/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// Reconciler implements controller.Reconciler for the event target type.
type Reconciler struct {
	base       common.GenericServiceReconciler[*v1alpha1.GoogleCloudPubSubTarget, listersv1alpha1.GoogleCloudPubSubTargetNamespaceLister]
	adapterCfg *adapterConfig
}

// Check that our Reconciler implements Interface
var _ reconcilerv1alpha1.Interface = (*Reconciler)(nil)
github.com/zeiss/typhoon
// ReconcileKind implgithub.com/zeiss/typhoon
func (r *Reconcilegithub.com/zeiss/typhoonxt, trg *v1alpha1.GoogleCloudPubSubTarget) reconciler.Event {
	// injecgithub.com/zeiss/typhoon reconciliation logic
	ctx = commonv1alpha1.WithReconcilable(ctx, trg)

	return r.base.ReconcileAdapter(ctx, r)
}
