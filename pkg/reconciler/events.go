package reconciler

// Reasons for API Events
const (
	// ReasonRBACCreate indicates that an RBAC object was successfully created.
	ReasonRBACCreate = "CreateRBAC"
	// ReasonRBACUpdate indicates that an RBAC object was successfully updated.
	ReasonRBACUpdate = "UpdateRBAC"
	// ReasonFailedRBACCreate indicates that the creation of an RBAC object failed.
	ReasonFailedRBACCreate = "FailedRBACCreate"
	// ReasonFailedRBACUpdate indicates that the update of an RBAC object failed.
	ReasonFailedRBACUpdate = "FailedRBACUpdate"

	// ReasonAdapterCreate indicates that an adapter object was successfully created.
	ReasonAdapterCreate = "CreateAdapter"
	// ReasonAdapterUpdate indicates that an adapter object was successfully updated.
	ReasonAdapterUpdate = "UpdateAdapter"
	// ReasonFailedAdapterCreate indicates that the creation of an adapter object failed.
	ReasonFailedAdapterCreate = "FailedAdapterCreate"
	// ReasonFailedAdapterUpdate indicates that the update of an adapter object failed.
	ReasonFailedAdapterUpdate = "FailedAdapterUpdate"

	// ReasonBadSinkURI indicates that the URI of a sink can't be determined.
	ReasonBadSinkURI = "BadSinkURI"

	// ReasonInvalidSpec indicates that spec of a reconciled object is invalid.
	ReasonInvalidSpec = "InvalidSpec"
)
