package v1alpha1

// Reasons for status conditions
const (
	// GCloudReasonNoClient is set on a status condition when a Google Cloud API client cannot be obtained.
	GCloudReasonNoClient = "NoClient"
	// GCloudReasonAPIError is set on a status condition when a Google Cloud API returns an error.
	GCloudReasonAPIError = "APIError"
)
