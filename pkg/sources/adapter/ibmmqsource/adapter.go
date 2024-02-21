//go:build !noclibs



package ibmmqsource

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/sources"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	"github.com/zeiss/typhoon/pkg/sources/adapter/ibmmqsource/mq"
)

var _ pkgadapter.Adapter = (*ibmmqsourceAdapter)(nil)

type ibmmqsourceAdapter struct {
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger
	mt       *pkgadapter.MetricTag

	mqEnvs *SourceEnvAccessor
}

//github.com/zeiss/typhoonentation
fugithub.com/zeiss/typhoon envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: sources.IBMMQSourceResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	env := envAcc.(*SourceEnvAccessor)

	return &ibmmqsourceAdapter{
		ceClient: ceClient,
		logger:   logger,
		mt:       mt,
		mqEnvs:   env,
	}
}

// Returns if stopCh is closed or Send() returns an error.
func (a *ibmmqsourceAdapter) Start(ctx context.Context) error {
	a.logger.Info("Starting IBMMQSource Adapter")

	conn, err := mq.NewConnection(a.mqEnvs.ConnectionConfig, a.mqEnvs.Auth)
	if err != nil {
		return fmt.Errorf("failed to create IBM MQ connection: %w", err)
	}
	defer conn.Disc()

	queue, err := mq.OpenQueue(a.mqEnvs.ConnectionConfig.QueueName, a.mqEnvs.DeadLetterQueue, conn)
	if err != nil {
		return fmt.Errorf("failed to open IBM MQ queue: %w", err)
	}
	defer queue.Close()

	ctx = pkgadapter.ContextWithMetricTag(ctx, a.mt)

	err = queue.RegisterCallback(a.eventHandler(ctx), a.mqEnvs.Delivery, a.logger)
	if err != nil {
		return fmt.Errorf("failed to register callback: %w", err)
	}
	defer queue.DeleteMessageHandle()
	defer queue.DeregisterCallback()

	if err := queue.StartListen(conn); err != nil {
		return fmt.Errorf("failed to start callback listener: %w", err)
	}
	defer queue.StopCallback(conn)

	<-ctx.Done()
	return nil
}

func (a *ibmmqsourceAdapter) eventHandler(ctx context.Context) mq.Handler {
	return func(data []byte, correlID string) error {
		event := cloudevents.NewEvent(cloudevents.VersionV1)
		event.SetType(v1alpha1.IBMMQSourceEventType)
		event.SetSource(fmt.Sprintf("%s/%s", a.mqEnvs.ConnectionConfig.ConnectionName, strings.ToLower(a.mqEnvs.ConnectionConfig.QueueName)))
		if correlID != "" {
			event.SetExtension(mq.CECorrelIDAttr, correlID)
		}
		contentType := cloudevents.TextPlain
		if json.Valid(data) {
			contentType = cloudevents.ApplicationJSON
		}
		if err := event.SetData(contentType, data); err != nil {
			a.logger.Errorw("Can't set Cloudevent data", zap.Error(err))
			return err
		}
		if res := a.ceClient.Send(ctx, event); cloudevents.IsUndelivered(res) {
			a.logger.Errorw("Cloudevent is not delivered", zap.Error(res))
			return res
		}
		return nil
	}
}
