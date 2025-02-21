package webhooksource

import (
	"context"

	sourcesalphav1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	webhookreconciler "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/sources/v1alpha1/webhooksource"

	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/kmeta"
	"knative.dev/pkg/reconciler"
)

// newWebhookSourceSkipped makes a new reconciler event with event type Normal, and
// reason WebhookSourceNotReady
func newWebhookSourceSkipped() reconciler.Event {
	return reconciler.NewEvent(corev1.EventTypeNormal, "WebhookSourceSkipped", "WebhookSource is not ready")
}

// newWebhookSourceNotReady makes a new reconciler event with event type Normal, and
// reason WebhookSourceNotReady
func newWebhookSourceSynchronized() reconciler.Event {
	return reconciler.NewEvent(corev1.EventTypeNormal, "WebhookSourceSynchronized", "WebhookSource adapter is synchronized")
}

// Reconciler reconciles WebhookSources
type Reconciler struct {
	mtadapter MTAdapter
}

// Check that our Reconciler implements ReconcileKind.
var _ webhookreconciler.Interface = (*Reconciler)(nil)

func (r *Reconciler) ReconcileKind(ctx context.Context, source *sourcesalphav1.WebhookSource) reconciler.Event {
	if !source.Status.IsReady() {
		return newWebhookSourceSkipped()
	}

	// Update the adapter state
	r.mtadapter.Update(ctx, source)

	return newWebhookSourceSynchronized()
}

func (r *Reconciler) deleteFunc(obj interface{}) {
	if obj == nil {
		return
	}
	acc, err := kmeta.DeletionHandlingAccessor(obj)
	if err != nil {
		return
	}
	WebhookSource, ok := acc.(*sourcesalphav1.WebhookSource)
	if !ok || WebhookSource == nil {
		return
	}

	r.mtadapter.Remove(WebhookSource)
}
