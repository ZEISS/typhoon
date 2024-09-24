package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	zapt "go.uber.org/zap/zaptest"
)

const (
	tAPIVersion      = "53.0"
	tClientID        = "abcde"
	tResponseChannel = "/test/channel"
	tMetaChannel     = "/meta"
	tReconnect       = "handshake"
	// tResponseFinishChannel will signal that the
	// test case has completed and the server should
	// not expect new events.
	tResponseFinishChannel = "/test/finish"
)

var tSubscription = []Subscription{
	{
		Channel:  "/channel1",
		ReplayID: -1,
	},
}

type response struct {
	handshake []HandshakeResponse
	subscribe []SubscriptionResponse
	connect   []ConnectResponse
}

func TestBayeux(t *testing.T) {
	logger := zapt.NewLogger(t).Sugar()

	testCases := map[string]struct {
		headers map[string]string

		// responses need follow the order they are
		// expected to be requested.
		responses []response

		expectedConnectResponses int
		expectedConnectErrors    int
	}{
		"no connect": {
			responses: []response{
				{handshake: handshakeResponse()},
				{subscribe: subscribeResponse()},
			},

			expectedConnectResponses: 0,
			expectedConnectErrors:    0,
		},

		"connect, receive 1 message": {
			responses: []response{
				{handshake: handshakeResponse()},
				{subscribe: subscribeResponse()},
				{connect: connectResponse()},
			},

			expectedConnectResponses: 1,
			expectedConnectErrors:    0,
		},

		"connect, receive 2 messages": {
			responses: []response{
				{handshake: handshakeResponse()},
				{subscribe: subscribeResponse()},
				{connect: connectResponse()},
				{connect: connectResponse()},
			},

			expectedConnectResponses: 2,
			expectedConnectErrors:    0,
		},

		"connect, with meta successful and 2 messages": {
			responses: []response{
				{handshake: handshakeResponse()},
				{subscribe: subscribeResponse()},
				{connect: connectResponse()},
				{connect: connectResponse(connectWithChannel(tMetaChannel))},
				{connect: connectResponse()},
			},

			expectedConnectResponses: 2,
			expectedConnectErrors:    0,
		},

		"connect, with meta advicing handshake": {
			responses: []response{
				{handshake: handshakeResponse()},
				{subscribe: subscribeResponse()},
				{connect: connectResponse(
					connectWithChannel(tMetaChannel),
					connectWithSuccessful(false),
					connectWithAdviceReconnect(tReconnect),
				)},
				{handshake: handshakeResponse()},
				{subscribe: subscribeResponse()},
				{connect: connectResponse()},
			},

			expectedConnectResponses: 1,
			expectedConnectErrors:    0,
		},

		"connect, with meta advicing handshake and 2 messages": {
			responses: []response{
				{handshake: handshakeResponse()},
				{subscribe: subscribeResponse()},
				{connect: connectResponse()},
				{connect: connectResponse(
					connectWithChannel(tMetaChannel),
					connectWithSuccessful(false),
					connectWithAdviceReconnect(tReconnect),
				)},
				{handshake: handshakeResponse()},
				{subscribe: subscribeResponse()},
				{connect: connectResponse()},
			},

			expectedConnectResponses: 2,
			expectedConnectErrors:    0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			sf := httptest.NewServer(mockBayeuxServerHandler(tc.responses))
			defer sf.Close()

			dispatcher := &eventDispatcher{
				eof: make(chan struct{}),
			}

			b := NewBayeux(tAPIVersion, sf.URL, tSubscription, dispatcher, sf.Client(), logger)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			go func() {
				clientErr := b.Start(ctx)
				assert.NoError(t, clientErr, "The bayeux client failed")
			}()

			select {
			case <-time.After(4000 * time.Millisecond):
				t.Fatal("Test timed out.")

			case <-dispatcher.eof:
			}

			require.Len(t, dispatcher.dispatchedEvents, tc.expectedConnectResponses, "Unexpected number of responses from connect")
			require.Len(t, dispatcher.dispatchedErrors, tc.expectedConnectErrors, "Unexpected number of errors from dispatcher")
		})
	}
}

func handshakeResponse() []HandshakeResponse {
	return []HandshakeResponse{{
		CommonResponse: CommonResponse{
			Successful: true,
			ClientID:   tClientID,
		},
	}}
}

func subscribeResponse() []SubscriptionResponse {
	return []SubscriptionResponse{{
		CommonResponse: CommonResponse{
			Successful: true,
			ClientID:   tClientID,
		},
	}}
}

type connectResponseOption func(*ConnectResponse)

func connectResponse(opts ...connectResponseOption) []ConnectResponse {
	cr := ConnectResponse{
		CommonResponse: CommonResponse{
			Successful: true,
			ClientID:   tClientID,
			Channel:    tResponseChannel,
		},
	}

	for _, opt := range opts {
		opt(&cr)
	}

	return []ConnectResponse{cr}
}

func connectWithChannel(channel string) connectResponseOption {
	return func(cr *ConnectResponse) {
		cr.Channel = channel
	}
}

func connectWithSuccessful(success bool) connectResponseOption {
	return func(cr *ConnectResponse) {
		cr.Successful = success
	}
}

func connectWithAdviceReconnect(reconnect string) connectResponseOption {
	return func(cr *ConnectResponse) {
		cr.Advice.Reconnect = reconnect
	}
}

var _ EventDispatcher = (*eventDispatcher)(nil)

type eventDispatcher struct {
	eof              chan struct{}
	dispatchedEvents []*ConnectResponse
	dispatchedErrors []error
}

func (e *eventDispatcher) DispatchEvent(ctx context.Context, res *ConnectResponse) {
	if res.Channel == tResponseFinishChannel {
		close(e.eof)
		return
	}
	e.dispatchedEvents = append(e.dispatchedEvents, res)
}

func (e *eventDispatcher) DispatchError(err error) {
	e.dispatchedErrors = append(e.dispatchedErrors, err)
}

func mockBayeuxServerHandler(responses []response) http.HandlerFunc {
	// responseIndex tracks the ordered responses for
	// expected requests as defined by the test case.
	responseIndex := 0

	return func(w http.ResponseWriter, r *http.Request) {
		total := len(responses)
		switch {
		case total == responseIndex:
			// no more responses for this test case, return a response that includes
			// the finish channel so that we can exit the server from the test.
			if err := json.NewEncoder(w).Encode(connectResponse(connectWithChannel(tResponseFinishChannel))); err != nil {
				http.Error(w, fmt.Sprintf("error encoding final connect response: %+v", err), http.StatusInternalServerError)
			}

			// increase the responseIndex so that subsequent calls fall in the
			// wait and noop case
			responseIndex++
			return

		case total < responseIndex:
			// after finish flag is set we might receive further requests
			// which will be ignored. A small time lapse is waited to let
			// the main test routine to finish without the server loop
			// consuming resources.
			time.Sleep(100 * time.Millisecond)
			return
		}

		// pick next response
		response := responses[responseIndex]
		responseIndex++

		// not expected URL
		if r.URL.Path != "/cometd/"+tAPIVersion {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

		var rb map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&rb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		channel, ok := rb["channel"].(string)
		if !ok {
			http.Error(w, "unkwown channel format at request", http.StatusInternalServerError)
		}

		switch channel {
		case handshakeChannel:
			if response.handshake == nil {
				http.Error(w, fmt.Sprintf("unexpected handshake request: %+v", response), http.StatusInternalServerError)
				return
			}
			if err := json.NewEncoder(w).Encode(response.handshake); err != nil {
				http.Error(w, fmt.Sprintf("error encoding handshake response: %+v", err), http.StatusInternalServerError)
			}

		case subscribeChannel:
			if response.subscribe == nil {
				http.Error(w, fmt.Sprintf("unexpected subscribe request: %+v", response), http.StatusInternalServerError)
				return
			}
			if err := json.NewEncoder(w).Encode(response.subscribe); err != nil {
				http.Error(w, fmt.Sprintf("error encoding subscribe response: %+v", err), http.StatusInternalServerError)
			}

		case connectChannel:
			if response.connect == nil {
				http.Error(w, fmt.Sprintf("unexpected connect request: %+v", response), http.StatusInternalServerError)
				return
			}
			if err := json.NewEncoder(w).Encode(response.connect); err != nil {
				http.Error(w, fmt.Sprintf("error encoding connect response: %+v", err), http.StatusInternalServerError)
			}

		default:
			http.Error(w, fmt.Sprintf("unkwown channel at request: %s", channel), http.StatusInternalServerError)
		}
	}
}
