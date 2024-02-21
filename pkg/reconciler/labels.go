package reconciler

// Kubernetes recommended labels
// https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/
const (
	// appNameLabel is the name of the application.
	appNameLabel = "app.kubernetes.io/name"
	// appInstanceLabel is a unique name identifying the instance of an application.
	appInstanceLabel = "app.kubernetes.io/instance"
	// appComponentLabel is the component within the architecture.
	appComponentLabel = "app.kubernetes.io/component"
	// appPartOfLabel is the name of a higher level application this one is part of.
	appPartOfLabel = "app.kubernetes.io/part-of"
	// appManagedByLabel is the tool being used to manage the operation of an application.
	appManagedByLabel = "app.kubernetes.io/managed-by"
)

// Common label values
const (
	partOf           = "typhoon"
	managedBy        = "typhoon-controller"
	componentAdapter = "adapter"
)

// labelsPropagationList is a list of labels that, if present on the parent
// object, should be propagated to the adapters.
var labelsPropagationList = []string{
	"bridges.typhoon.zeiss.com/id",
	"flow.typhoon.zeiss.com/created-by",
}
