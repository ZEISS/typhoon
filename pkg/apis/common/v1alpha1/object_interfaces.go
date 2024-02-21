package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"

	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

// Reconcilable is implemented by all components.
type Reconcilable interface {
	metav1.Object
	runtime.Object
	// OwnerRefable is used to construct a generic reconciler for each
	// component type, and convert custom objects to owner references.
	kmeta.OwnerRefable
	// KRShaped is used by generated reconcilers to perform pre and
	// post-reconcile status updates.
	duckv1.KRShaped

	// GetStatusManager returns a manager for the component's status.
	GetStatusManager() *StatusManager
}

// AdapterConfigurable is implemented by types that can override the default
// configuration of their receive adapter.
type AdapterConfigurable interface {
	// GetAdapterOverrides returns the adapter overrides.
	GetAdapterOverrides() *AdapterOverrides
}

// EventSource is implemented by types that emit events, either by sending them
// to a sink or by replying to incoming event requests.
type EventSource interface {
	// GetEventTypes returns the event types generated by the component.
	GetEventTypes() []string
	// AsEventSource returns a unique reference to the component suitable
	// for use as a CloudEvent 'source' attribute.
	AsEventSource() string
}

// EventSender is implemented by types that send events to a sink.
type EventSender interface {
	// GetSink returns the component's event sink.
	GetSink() *duckv1.Destination
}

// EventReceiver is implemented by types that receive and process events.
type EventReceiver interface {
	// AcceptedEventTypes returns the event types accepted by the target.
	AcceptedEventTypes() []string
}

// MultiTenant is implemented by all multi-tenant component types.
type MultiTenant interface {
	IsMultiTenant() bool
}

// IsMultiTenant returns whether the given component type is multi-tenant.
func IsMultiTenant(r Reconcilable) bool {
	mt, ok := r.(MultiTenant)
	return ok && mt.IsMultiTenant()
}

// ServiceAccountProvider is implemented by types which are able to influence
// the shape of the ServiceAccount used by their own receive adapter.
type ServiceAccountProvider interface {
	WantsOwnServiceAccount() bool
	ServiceAccountOptions() []resource.ServiceAccountOption
}

// WantsOwnServiceAccount returns whether the given component instance should
// have a dedicated ServiceAccount associated with its receive adapter.
func WantsOwnServiceAccount(r Reconcilable) bool {
	saProvider, ok := r.(ServiceAccountProvider)
	return ok && saProvider.WantsOwnServiceAccount()
}

// ServiceAccountOptions returns functional options for mutating the
// ServiceAccount associated with a given component instance.
func ServiceAccountOptions(r Reconcilable) []resource.ServiceAccountOption {
	saProvider, ok := r.(ServiceAccountProvider)
	if !ok {
		return nil
	}

	return saProvider.ServiceAccountOptions()
}
