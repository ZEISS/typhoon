package googlecloudstoragesource

const (
	// ReasonSubscribed indicates that a source subscribed to change
	// notifications from a Cloud Storage bucket.
	ReasonSubscribed = "Subscribed"
	// ReasonUnsubscribed indicates that a source unsubscribed from change
	// notifications from a Cloud Storage bucket.
	ReasonUnsubscribed = "Unsubscribed"
	// ReasonFailedSubscribe indicates a failure while synchronizing the
	// notification configuration of a Cloud Storage bucket, or the Pub/Sub
	// subscription it depends on.
	ReasonFailedSubscribe = "FailedSubscribe"
	// ReasonFailedUnsubscribe indicates a failure while deleting the
	// notification configuration of a Cloud Storage bucket, or the Pub/Sub
	// subscription it depends on.
	ReasonFailedUnsubscribe = "FailedUnsubscribe"
)
