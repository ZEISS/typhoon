// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListSystems request
	ListSystems(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ShowSystem request
	ShowSystem(ctx context.Context, systemId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTeamTeamId request
	GetTeamTeamId(ctx context.Context, teamId openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListTeam request
	ListTeam(ctx context.Context, params *ListTeamParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateTeamWithBody request with any body
	CreateTeamWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateTeam(ctx context.Context, body CreateTeamJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Version request
	Version(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListSystems(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListSystemsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ShowSystem(ctx context.Context, systemId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewShowSystemRequest(c.Server, systemId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTeamTeamId(ctx context.Context, teamId openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTeamTeamIdRequest(c.Server, teamId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListTeam(ctx context.Context, params *ListTeamParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListTeamRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateTeamWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateTeamRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateTeam(ctx context.Context, body CreateTeamJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateTeamRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Version(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewVersionRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListSystemsRequest generates requests for ListSystems
func NewListSystemsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/systems")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewShowSystemRequest generates requests for ShowSystem
func NewShowSystemRequest(server string, systemId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "systemId", runtime.ParamLocationPath, systemId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/systems/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetTeamTeamIdRequest generates requests for GetTeamTeamId
func NewGetTeamTeamIdRequest(server string, teamId openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "teamId", runtime.ParamLocationPath, teamId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/team/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListTeamRequest generates requests for ListTeam
func NewListTeamRequest(server string, params *ListTeamParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/teams")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Offset != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "offset", runtime.ParamLocationQuery, *params.Offset); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Limit != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateTeamRequest calls the generic CreateTeam builder with application/json body
func NewCreateTeamRequest(server string, body CreateTeamJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateTeamRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateTeamRequestWithBody generates requests for CreateTeam with any type of body
func NewCreateTeamRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/teams")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewVersionRequest generates requests for Version
func NewVersionRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/version")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ListSystemsWithResponse request
	ListSystemsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListSystemsResponse, error)

	// ShowSystemWithResponse request
	ShowSystemWithResponse(ctx context.Context, systemId string, reqEditors ...RequestEditorFn) (*ShowSystemResponse, error)

	// GetTeamTeamIdWithResponse request
	GetTeamTeamIdWithResponse(ctx context.Context, teamId openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetTeamTeamIdResponse, error)

	// ListTeamWithResponse request
	ListTeamWithResponse(ctx context.Context, params *ListTeamParams, reqEditors ...RequestEditorFn) (*ListTeamResponse, error)

	// CreateTeamWithBodyWithResponse request with any body
	CreateTeamWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateTeamResponse, error)

	CreateTeamWithResponse(ctx context.Context, body CreateTeamJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateTeamResponse, error)

	// VersionWithResponse request
	VersionWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*VersionResponse, error)
}

type ListSystemsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Systems
}

// Status returns HTTPResponse.Status
func (r ListSystemsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListSystemsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ShowSystemResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ShowSystemResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ShowSystemResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTeamTeamIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Team
}

// Status returns HTTPResponse.Status
func (r GetTeamTeamIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTeamTeamIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListTeamResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Limit   *float32 `json:"limit,omitempty"`
		Offset  *float32 `json:"offset,omitempty"`
		Results *[]Team  `json:"results,omitempty"`
		Total   *float32 `json:"total,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListTeamResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListTeamResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateTeamResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *Team
}

// Status returns HTTPResponse.Status
func (r CreateTeamResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateTeamResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type VersionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Version
}

// Status returns HTTPResponse.Status
func (r VersionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r VersionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListSystemsWithResponse request returning *ListSystemsResponse
func (c *ClientWithResponses) ListSystemsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListSystemsResponse, error) {
	rsp, err := c.ListSystems(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListSystemsResponse(rsp)
}

// ShowSystemWithResponse request returning *ShowSystemResponse
func (c *ClientWithResponses) ShowSystemWithResponse(ctx context.Context, systemId string, reqEditors ...RequestEditorFn) (*ShowSystemResponse, error) {
	rsp, err := c.ShowSystem(ctx, systemId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseShowSystemResponse(rsp)
}

// GetTeamTeamIdWithResponse request returning *GetTeamTeamIdResponse
func (c *ClientWithResponses) GetTeamTeamIdWithResponse(ctx context.Context, teamId openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetTeamTeamIdResponse, error) {
	rsp, err := c.GetTeamTeamId(ctx, teamId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTeamTeamIdResponse(rsp)
}

// ListTeamWithResponse request returning *ListTeamResponse
func (c *ClientWithResponses) ListTeamWithResponse(ctx context.Context, params *ListTeamParams, reqEditors ...RequestEditorFn) (*ListTeamResponse, error) {
	rsp, err := c.ListTeam(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListTeamResponse(rsp)
}

// CreateTeamWithBodyWithResponse request with arbitrary body returning *CreateTeamResponse
func (c *ClientWithResponses) CreateTeamWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateTeamResponse, error) {
	rsp, err := c.CreateTeamWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateTeamResponse(rsp)
}

func (c *ClientWithResponses) CreateTeamWithResponse(ctx context.Context, body CreateTeamJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateTeamResponse, error) {
	rsp, err := c.CreateTeam(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateTeamResponse(rsp)
}

// VersionWithResponse request returning *VersionResponse
func (c *ClientWithResponses) VersionWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*VersionResponse, error) {
	rsp, err := c.Version(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseVersionResponse(rsp)
}

// ParseListSystemsResponse parses an HTTP response from a ListSystemsWithResponse call
func ParseListSystemsResponse(rsp *http.Response) (*ListSystemsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListSystemsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Systems
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseShowSystemResponse parses an HTTP response from a ShowSystemWithResponse call
func ParseShowSystemResponse(rsp *http.Response) (*ShowSystemResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ShowSystemResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetTeamTeamIdResponse parses an HTTP response from a GetTeamTeamIdWithResponse call
func ParseGetTeamTeamIdResponse(rsp *http.Response) (*GetTeamTeamIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTeamTeamIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Team
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseListTeamResponse parses an HTTP response from a ListTeamWithResponse call
func ParseListTeamResponse(rsp *http.Response) (*ListTeamResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListTeamResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Limit   *float32 `json:"limit,omitempty"`
			Offset  *float32 `json:"offset,omitempty"`
			Results *[]Team  `json:"results,omitempty"`
			Total   *float32 `json:"total,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateTeamResponse parses an HTTP response from a CreateTeamWithResponse call
func ParseCreateTeamResponse(rsp *http.Response) (*CreateTeamResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateTeamResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest Team
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseVersionResponse parses an HTTP response from a VersionWithResponse call
func ParseVersionResponse(rsp *http.Response) (*VersionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &VersionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Version
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
