package targets

import "k8s.io/apimachinery/pkg/runtime/schema"

// GroupName is the name of the API group this package's resources belong to.
const GroupName = "targets.typhoon.zeiss.com"

var (
	// HTTPTargetResource respresents an event target for HTTP endpoint.
	HTTPTargetResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "httptargets",
	}
	// KafkaTargetResource respresents an event target for Kafka.
	KafkaTargetResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "kafkatargets",
	}
	// LogzTargetResource respresents an event target for Logz.
	LogzTargetResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "logztargets",
	}
	// CloudEventsTargetResource respresents an event target for CloudEvents gateway.
	CloudEventsTargetResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "cloudeventstargets",
	}
	// DatadogTargetResource respresents an event target for Datadog.
	DatadogTargetResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "datadogtargets",
	}
	// LogzMetricsTargetResource respresents an event target for Logz Metrics.
	LogzMetricsTargetResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "logzmetricstargets",
	}
	// SplunkTargetResource respresents an event target for Splunk.
	SplunkTargetResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "splunktargets",
	}
	// SalesforceTargetResource respresents an event target for Salesforce.
	SalesforceTargetResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "salesforcetargets",
	}
	// JiraTargetResource respresents an event target for Jira.
	JiraTargetResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "jiratargets",
	}
)
