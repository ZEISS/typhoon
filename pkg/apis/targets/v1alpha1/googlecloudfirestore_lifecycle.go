

package v1alpha1

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

// Managed event types
const (
	EventTypeGoogleCloudFirestoreWriteResponse = "io.triggermesh.google.firestore.write.response"
	EventTypeGoogleCloudFirestoreWrite         = "io.triggermesh.google.firestore.write"

	EventTypeGoogleCloudFirestoreQueryTablesResponse = "io.triggermesh.google.firestore.query.tables.response"
	EventTypeGoogleCloudFirestoreQueryTables         = "io.triggermesh.google.firestore.query.tables"

	EventTypeGoogleCloudFirestoreQueryTableResponse = "io.triggermesh.google.firestore.query.table.response"
	EventTypeGoogleCloudFirestoreQueryTable         = "io.triggermesh.google.firestore.query.table"
)
github.com/zeiss/typhoon
// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*GoogleCloudFirestoreTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("GoogleCloudFirestoreTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*GoogleCloudFirestoreTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *GoogleCloudFirestoreTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *GoogleCloudFirestoreTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*GoogleCloudFirestoreTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeGoogleCloudFirestoreWrite,
		EventTypeGoogleCloudFirestoreQueryTables,
		EventTypeGoogleCloudFirestoreQueryTable,
	}
}

// GetEventTypes implements EventSource.
func (*GoogleCloudFirestoreTarget) GetEventTypes() []string {
	return []string{
		EventTypeGoogleCloudFirestoreWriteResponse,
		EventTypeGoogleCloudFirestoreWrite,
		EventTypeGoogleCloudFirestoreQueryTablesResponse,
		EventTypeGoogleCloudFirestoreQueryTables,
		EventTypeGoogleCloudFirestoreQueryTableResponse,
		EventTypeGoogleCloudFirestoreQueryTable,
	}
}

// AsEventSource implements EventSource.
func (t *GoogleCloudFirestoreTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "io.triggermesh." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *GoogleCloudFirestoreTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// WantsOwnServiceAccount implements ServiceAccountProvider.
func (t *GoogleCloudFirestoreTarget) WantsOwnServiceAccount() bool {
	return t.Spec.Auth != nil && t.Spec.Auth.WantsOwnServiceAccount()
}

// ServiceAccountOptions implements ServiceAccountProvider.
func (t *GoogleCloudFirestoreTarget) ServiceAccountOptions() []resource.ServiceAccountOption {
	if t.Spec.Auth == nil {
		return []resource.ServiceAccountOption{}
	}
	return t.Spec.Auth.ServiceAccountOptions()
}

// SetDefaults implements apis.Defaultable
func (t *GoogleCloudFirestoreTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *GoogleCloudFirestoreTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
