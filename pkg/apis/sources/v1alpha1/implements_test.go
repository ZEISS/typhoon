package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/require"
	duck "knative.dev/pkg/apis/duck"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

func TestTypesImplements(t *testing.T) {
	tests := []struct {
		instance interface{}
		iface    duck.Implementable
	}{
		{&AzureServiceBusSource{}, &duckv1.Conditions{}},
		{&AzureServiceBusSource{}, &duckv1.Source{}},
		{&CloudEventsSource{}, &duckv1.Conditions{}},
		{&CloudEventsSource{}, &duckv1.Source{}},
		{&HTTPPollerSource{}, &duckv1.Conditions{}},
		{&HTTPPollerSource{}, &duckv1.Source{}},
		{&KafkaSource{}, &duckv1.Conditions{}},
		{&KafkaSource{}, &duckv1.Source{}},
		{&OCIMetricsSource{}, &duckv1.Conditions{}},
		{&OCIMetricsSource{}, &duckv1.Source{}},
		{&PingSource{}, &duckv1.Conditions{}},
		{&PingSource{}, &duckv1.Source{}},
		{&SalesforceSource{}, &duckv1.Conditions{}},
		{&SalesforceSource{}, &duckv1.Source{}},
		{&WebhookSource{}, &duckv1.Conditions{}},
		{&WebhookSource{}, &duckv1.Source{}},
	}

	for _, test := range tests {
		require.NoError(t, duck.VerifyType(test.instance, test.iface))
	}
}
