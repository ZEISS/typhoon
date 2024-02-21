package azureservicebussource

const (
	// ReasonSubscribed indicates that a Subscription was created for an Azure Service Bus Topic.
	ReasonSubscribed = "Subscribed"
	// ReasonUnsubscribed indicates that a Subscription was removed from an Azure Service Bus Topic.
	ReasonUnsubscribed = "Unsubscribed"
	// ReasonFailedSubscribe indicates a failure while synchronizing a Subscription for an Azure Service Bus Topic.
	ReasonFailedSubscribe = "FailedSubscribe"
	// ReasonFailedUnsubscribe indicates a failure while removing a Subscription from an Azure Service Bus Topic.
	ReasonFailedUnsubscribe = "FailedUnsubscribe"
)
