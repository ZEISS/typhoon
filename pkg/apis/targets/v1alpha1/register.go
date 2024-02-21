package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/targets"
)

var (
	// SchemeGroupVersion contains the group and version used to register types for this custom API.
	SchemeGroupVersion = schema.GroupVersion{Group: targets.GroupName, Version: "v1alpha1"}
	// SchemeBuilder creates a Scheme builder that is used to register types for this custom API.
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme registers the types stored in SchemeBuilder.
	AddToScheme = SchemeBuilder.AddToScheme
)

// AllTypes is a list of all the types defined in this package.
var AllTypes = []v1alpha1.GroupObject{
	{Single: &CloudEventsTarget{}, List: &CloudEventsTargetList{}},
	{Single: &HTTPTarget{}, List: &HTTPTargetList{}},
	{Single: &KafkaTarget{}, List: &KafkaTargetList{}},
	{Single: &LogzMetricsTarget{}, List: &LogzMetricsTargetList{}},
	{Single: &LogzTarget{}, List: &LogzTargetList{}},
}

// addKnownTypes adds all this custom API's types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	for _, t := range AllTypes {
		scheme.AddKnownTypes(SchemeGroupVersion, t.Single, t.List)
	}
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind.
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
