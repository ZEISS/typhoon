package sources

import "k8s.io/apimachinery/pkg/runtime/schema"

// GroupName is the name of the API group this package's resources belong to.
const GroupName = "sources.typhoon.zeiss.com"

var (
	// CloudEventsSourceResource respresents an event source for CloudEvents.
	CloudEventsSourceResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "cloudeventssources",
	}

	// KafkaSourceResource respresents an event source for Kafka.
	KafkaSourceResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "kafkasources",
	}

	// HTTPPollerSourceResource represents an event source for polling HTTP endpoints.
	HTTPPollerSourceResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "httppollersources",
	}

	// SalesforceSourceResource represents an event source for Salesforce.
	SalesforceSourceResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "salesforcesources",
	}

	// WebhookSourceResource represents an event source for HTTP webhooks.
	WebhookSourceResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "webhooksources",
	}
)
