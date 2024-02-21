

package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*GoogleCloudStorageSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("GoogleCloudStorageSource")
}

// GetConditionSet implements duckv1.KRShaped.
func (*GoogleCloudStorageSource) GetConditionSet() apis.ConditionSet {
	return googleCloudStorageSourceConditionSet
}

// GetStatus implements duckv1.KRShaped.
fugithub.com/zeiss/typhoonetStatus() *duckv1.Status {
	return &s.Status.Status.Status
}

// GetSink implements EventSender.
func (s *GoogleCloudStorageSource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetStatusManager implements Reconcilable.
func (s *GoogleCloudStorageSource) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: s.GetConditionSet(),
		Status:       &s.Status.Status,
	}
}

// AsEventSource implements EventSource.
func (s *GoogleCloudStorageSource) AsEventSource() string {
	return "gs://" + s.Spec.Bucket
}

// GetAdapterOverrides implements AdapterConfigurable.
func (s *GoogleCloudStorageSource) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return s.Spec.AdapterOverrides
}

// WantsOwnServiceAccount implements ServiceAccountProvider.
func (s *GoogleCloudStorageSource) WantsOwnServiceAccount() bool {
	return s.Spec.Auth.WantsOwnServiceAccount()
}

// ServiceAccountOptions implements ServiceAccountProvider.
func (s *GoogleCloudStorageSource) ServiceAccountOptions() []resource.ServiceAccountOption {
	return s.Spec.Auth.ServiceAccountOptions()
}

// Supported event types
const (
	GoogleCloudStorageGenericEventType = "com.google.cloud.storage.notification"

	GoogleCloudStorageFinalizeEventType = "com.google.cloud.storage.objectfinalize"
	GoogleCloudStorageUpdateEventType   = "com.google.cloud.storage.objectmetadataupdate"
	GoogleCloudStorageDeleteEventType   = "com.google.cloud.storage.objectdelete"
	GoogleCloudStorageArchiveEventType  = "com.google.cloud.storage.objectarchive"
)

// GetEventTypes returns the event types generated by the source.
func (s *GoogleCloudStorageSource) GetEventTypes() []string {
	var eventTypes []string
	if len(s.Spec.EventTypes) == 0 {
		eventTypes = []string{
			GoogleCloudStorageFinalizeEventType,
			GoogleCloudStorageUpdateEventType,
			GoogleCloudStorageDeleteEventType,
			GoogleCloudStorageArchiveEventType,
		}
	}
	for _, eventType := range s.Spec.EventTypes {
		eventTypes = append(eventTypes, PubSubAttributeToCEType(eventType))
	}
	return eventTypes
}

// PubSubAttributeToCEType translates Pub/Sub event type to CloudEvents type attribute.
func PubSubAttributeToCEType(eventType string) string {
	switch eventType {
	case "OBJECT_FINALIZE":
		return GoogleCloudStorageFinalizeEventType
	case "OBJECT_METADATA_UPDATE":
		return GoogleCloudStorageUpdateEventType
	case "OBJECT_DELETE":
		return GoogleCloudStorageDeleteEventType
	case "OBJECT_ARCHIVE":
		return GoogleCloudStorageArchiveEventType
	}
	return GoogleCloudStorageGenericEventType
}

// Status conditions
const (
	// GoogleCloudStorageConditionSubscribed has status True when the source has subscribed to a topic.
	GoogleCloudStorageConditionSubscribed apis.ConditionType = "Subscribed"
)

// googleCloudStorageSourceConditionSet is a set of conditions for
// GoogleCloudStorageSource objects.
var googleCloudStorageSourceConditionSet = v1alpha1.NewConditionSet(
	GoogleCloudStorageConditionSubscribed,
)

// MarkSubscribed sets the Subscribed condition to True.
func (s *GoogleCloudStorageSourceStatus) MarkSubscribed() {
	googleCloudStorageSourceConditionSet.Manage(s).MarkTrue(GoogleCloudStorageConditionSubscribed)
}

// MarkNotSubscribed sets the Subscribed condition to False with the given
// reason and message.
func (s *GoogleCloudStorageSourceStatus) MarkNotSubscribed(reason, msg string) {
	googleCloudStorageSourceConditionSet.Manage(s).MarkFalse(GoogleCloudStorageConditionSubscribed, reason, msg)
}

// SetDefaults implements apis.Defaultable
func (s *GoogleCloudStorageSource) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (s *GoogleCloudStorageSource) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
