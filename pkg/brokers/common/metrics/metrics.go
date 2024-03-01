package metrics

import (
	"context"
	"sync"

	"go.opencensus.io/resource"
	"go.opencensus.io/tag"
	"knative.dev/pkg/metrics/metricskey"
)

const (
	resourceTypeTyphoonBroker = "typhoon_broker"

	labelBrokerName = "broker_name"
	// labelUniqueName        = "unique_name"
	labelReceivedEventType = "received_type"
)

var (
	once sync.Once

	ReceivedEventTypeKey = tag.MustNewKey(labelReceivedEventType)
)

// func InitializeReportingContext(ctx context.Context, brokerName, instanceID string) context.Context {
func InitializeReportingContext(ctx context.Context, brokerName string) context.Context {
	once.Do(func() {
		ctx = metricskey.WithResource(ctx, resource.Resource{
			Type: resourceTypeTyphoonBroker,
			Labels: map[string]string{
				labelBrokerName: brokerName,
				// labelUniqueName: instanceID,
			},
		})
	})
	return ctx
}
