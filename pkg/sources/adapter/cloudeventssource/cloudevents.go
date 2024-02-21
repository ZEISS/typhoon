package cloudeventssource

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/adapter/fs"
)

type cloudEventsHandler struct {
	basicAuths KeyMountedValues

	cfw      fs.CachedFileWatcher
	ceServer cloudevents.Client
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger
	mt       *pkgadapter.MetricTag
}

// Start implements adapter.Adapter.
func (h *cloudEventsHandler) Start(ctx context.Context) error {
	h.cfw.Start(ctx)
	return h.ceServer.StartReceiver(ctx, h.handle)
}

func (h *cloudEventsHandler) handle(ctx context.Context, e event.Event) protocol.Result {
	err := e.Validate()
	if err != nil {
		h.logger.Errorw("Incoming CloudEvent is not valid", zap.Error(err))
		return protocol.ResultNACK
	}

	result := h.ceClient.Send(ctx, e)
	if !cloudevents.IsACK(result) {
		h.logger.Errorw("Could not send CloudEvent", zap.Error(result))
	}

	return result
}

// code based on VMware's VEBA's webhook:
// https://github.com/vmware-samples/vcenter-event-broker-appliance/blob/e91e4bd8a17dad6ce4fe370c42a15694c03dac88/vmware-event-router/internal/provider/webhook/webhook.go#L167-L189
func (h *cloudEventsHandler) handleAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if ok {
			// reduce brute-force guessing attacks with constant-time comparisons
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))

			for _, kv := range h.basicAuths {
				p, err := h.cfw.GetContent(kv.MountedValueFile)
				if err != nil {
					h.logger.Errorw(
						fmt.Sprintf("Could not retrieve password for user %q", kv.Key),
						zap.Error(err))
					continue
				}

				expectedUsernameHash := sha256.Sum256([]byte(kv.Key))
				expectedPasswordHash := sha256.Sum256(p)

				usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
				passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

				if usernameMatch && passwordMatch {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
