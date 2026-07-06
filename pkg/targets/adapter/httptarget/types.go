package httptarget

import "encoding/json"

// RequestData contains the parametrizable fields that users can provide
// for each request.
type RequestData struct {
	Headers    map[string]string `json:"headers,omitempty"`
	Query      string            `json:"query_string,omitempty"`
	PathSuffix string            `json:"path_suffix,omitempty"`
	Body       json.RawMessage   `json:"body,omitempty"`
}
