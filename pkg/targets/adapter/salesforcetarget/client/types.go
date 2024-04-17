package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Action performed on the API URI
type Action string

// Available actions at the Salesforce API
const (
	ActionMetadata Action = http.MethodHead
	ActionCreate   Action = http.MethodPost
	ActionUpdate   Action = http.MethodPatch
	ActionRetrieve Action = http.MethodGet
	ActionDelete   Action = http.MethodDelete
)

// SalesforceAPIRequest contains common parameters used for
// interacting with Salesforce using the API.
type SalesforceAPIRequest struct {
	// Action top perform on the URI resource.
	Action Action `json:"action"`

	// ResourcePathPath defines the first part of the user defined path for the HTTP request.
	// It is placed just after the base path at the URL:
	// https://<salesforce-host>/services/data/<version>/ResourcePath
	ResourcePath string `json:"resource"`

	// ObjectPath determines the type to manage under the resource:
	// https://<salesforce-host>/services/data/<version>/<resource>/ObjectPath
	ObjectPath string `json:"object"`

	// RecordPath identifies the object instance to manage:
	// https://<salesforce-host>/services/data/<version>/<resource>/<object>/RecordPath
	//
	// In some cases this field can inform Record and Field by using "record/field" syntax.
	RecordPath string `json:"record"`

	// Query is the set of key and values appeded to the operation:
	// https://<salesforce-host>/services/data/<version>/<resource>/<object>/<record>?query
	Query map[string]string `json:"query"`

	// Payload is the JSON content of the request to be sent.
	Payload json.RawMessage `json:"payload"`
}

// Validate API request
func (sfr *SalesforceAPIRequest) Validate() error {
	if sfr.Action == "" {
		return fmt.Errorf("HTTP action on the resource needs to be specified")
	}

	if sfr.ResourcePath == "" {
		return fmt.Errorf("HTTP resource at the Salesforce API needs to be specified")
	}

	return nil
}
