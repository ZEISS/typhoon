package sources

import "k8s.io/apimachinery/pkg/runtime/schema"

// GroupName ...
const GroupName = "sources.typhoon.zeiss.com"

// Webhook ...
var WebhookResource = schema.GroupResource{
	Group:    GroupName,
	Resource: "webhook",
}

// HTTPResource ...
var HTTPResource = schema.GroupResource{
	Group:    GroupName,
	Resource: "http",
}
