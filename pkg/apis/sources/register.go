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

	// AzureServiceBusQueueSourceResource respresents an event source for
	// Azure Service Bus Queues.
	AzureServiceBusQueueSourceResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "azureservicebusqueuesources",
	}

	// AzureServiceBusSourceResource respresents an event source for
	// Azure Service Bus.
	AzureServiceBusSourceResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "azureservicebussources",
	}

	// AzureServiceBusTopicSourceResource respresents an event source for
	// Azure Service Bus Topics.
	AzureServiceBusTopicSourceResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "azureservicebustopicsources",
	}
)
