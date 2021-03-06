// Package flex provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package flex

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/faetools/client"
)

// operation paths

const (
	opPathDeleteChannelFormat    = "./v1/Channels/%s"
	opPathFetchChannelFormat     = "./v1/Channels/%s"
	opPathDeleteFlexFlowFormat   = "./v1/FlexFlows/%s"
	opPathFetchFlexFlowFormat    = "./v1/FlexFlows/%s"
	opPathUpdateFlexFlowFormat   = "./v1/FlexFlows/%s"
	opPathDeleteWebChannelFormat = "./v1/WebChannels/%s"
	opPathFetchWebChannelFormat  = "./v1/WebChannels/%s"
	opPathUpdateWebChannelFormat = "./v1/WebChannels/%s"
)

var (
	opPathListChannel        = client.MustParseURL("./v1/Channels")
	opPathCreateChannel      = client.MustParseURL("./v1/Channels")
	opPathFetchConfiguration = client.MustParseURL("./v1/Configuration")
	opPathListFlexFlow       = client.MustParseURL("./v1/FlexFlows")
	opPathCreateFlexFlow     = client.MustParseURL("./v1/FlexFlows")
	opPathListWebChannel     = client.MustParseURL("./v1/WebChannels")
	opPathCreateWebChannel   = client.MustParseURL("./v1/WebChannels")
)

// ClientInterface interface specification for the client.
type ClientInterface interface {
	// ListChannel request
	ListChannel(ctx context.Context, params *ListChannelParams, reqEditors ...client.RequestEditorFn) (*ListChannelResponse, error)

	// CreateChannel request with any body
	CreateChannelWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateChannelResponse, error)

	// DeleteChannel request
	DeleteChannel(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteChannelResponse, error)

	// FetchChannel request
	FetchChannel(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchChannelResponse, error)

	// FetchConfiguration request
	FetchConfiguration(ctx context.Context, params *FetchConfigurationParams, reqEditors ...client.RequestEditorFn) (*FetchConfigurationResponse, error)

	// ListFlexFlow request
	ListFlexFlow(ctx context.Context, params *ListFlexFlowParams, reqEditors ...client.RequestEditorFn) (*ListFlexFlowResponse, error)

	// CreateFlexFlow request with any body
	CreateFlexFlowWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateFlexFlowResponse, error)

	// DeleteFlexFlow request
	DeleteFlexFlow(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteFlexFlowResponse, error)

	// FetchFlexFlow request
	FetchFlexFlow(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchFlexFlowResponse, error)

	// UpdateFlexFlow request with any body
	UpdateFlexFlowWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateFlexFlowResponse, error)

	// ListWebChannel request
	ListWebChannel(ctx context.Context, params *ListWebChannelParams, reqEditors ...client.RequestEditorFn) (*ListWebChannelResponse, error)

	// CreateWebChannel request with any body
	CreateWebChannelWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateWebChannelResponse, error)

	// DeleteWebChannel request
	DeleteWebChannel(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteWebChannelResponse, error)

	// FetchWebChannel request
	FetchWebChannel(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchWebChannelResponse, error)

	// UpdateWebChannel request with any body
	UpdateWebChannelWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateWebChannelResponse, error)
}

// Client definition

// compile time assert that it fulfils the interface
var _ ClientInterface = (*Client)(nil)

// Client conforms to the OpenAPI3 specification for this service.
type Client client.Client

// NewClient creates a new Client with reasonable defaults.
func NewClient(opts ...client.Option) (*Client, error) {
	c, err := client.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	if c.BaseURL == nil {
		if err := client.WithBaseURL(DefaultServer)(c); err != nil {
			return nil, err
		}
	}

	return (*Client)(c), nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []client.RequestEditorFn) error {
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

// ListChannel: GET /v1/Channels

type ListChannelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		FlexChatChannels *[]FlexV1Channel `json:"flex_chat_channels,omitempty"`
		Meta             *struct {
			FirstPageUrl    *string `json:"first_page_url,omitempty"`
			Key             *string `json:"key,omitempty"`
			NextPageUrl     *string `json:"next_page_url,omitempty"`
			Page            *int    `json:"page,omitempty"`
			PageSize        *int    `json:"page_size,omitempty"`
			PreviousPageUrl *string `json:"previous_page_url,omitempty"`
			Url             *string `json:"url,omitempty"`
		} `json:"meta,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListChannelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListChannelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newListChannelRequest generates requests for ListChannel
func newListChannelRequest(baseURL *url.URL, params *ListChannelParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathListChannel)

	q := queryURL.Query()

	if params.PageSize != nil {
		if err := client.AddQueryParam(q, "PageSize", *params.PageSize); err != nil {
			return nil, err
		}
	}

	queryURL.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ListChannel returns a parsed response.
// GET /v1/Channels
func (c *Client) ListChannel(ctx context.Context, params *ListChannelParams, reqEditors ...client.RequestEditorFn) (*ListChannelResponse, error) {
	req, err := newListChannelRequest(c.BaseURL, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &ListChannelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			FlexChatChannels *[]FlexV1Channel `json:"flex_chat_channels,omitempty"`
			Meta             *struct {
				FirstPageUrl    *string `json:"first_page_url,omitempty"`
				Key             *string `json:"key,omitempty"`
				NextPageUrl     *string `json:"next_page_url,omitempty"`
				Page            *int    `json:"page,omitempty"`
				PageSize        *int    `json:"page_size,omitempty"`
				PreviousPageUrl *string `json:"previous_page_url,omitempty"`
				Url             *string `json:"url,omitempty"`
			} `json:"meta,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// CreateChannel: POST /v1/Channels

type CreateChannelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *FlexV1Channel
}

// Status returns HTTPResponse.Status
func (r CreateChannelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateChannelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newCreateChannelRequestWithBody generates requests for CreateChannel with any type of body
func newCreateChannelRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathCreateChannel)

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// CreateChannelWithBody returns a parsed response.
// POST /v1/Channels
func (c *Client) CreateChannelWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateChannelResponse, error) {
	req, err := newCreateChannelRequestWithBody(c.BaseURL, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &CreateChannelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest FlexV1Channel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}

// DeleteChannel: DELETE /v1/Channels/{Sid}

type DeleteChannelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteChannelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteChannelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newDeleteChannelRequest generates requests for DeleteChannel
func newDeleteChannelRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathDeleteChannelFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// DeleteChannel returns a parsed response.
// DELETE /v1/Channels/{Sid}
func (c *Client) DeleteChannel(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteChannelResponse, error) {
	req, err := newDeleteChannelRequest(c.BaseURL, sid)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &DeleteChannelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// FetchChannel: GET /v1/Channels/{Sid}

type FetchChannelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FlexV1Channel
}

// Status returns HTTPResponse.Status
func (r FetchChannelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchChannelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchChannelRequest generates requests for FetchChannel
func newFetchChannelRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchChannelFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// FetchChannel returns a parsed response.
// GET /v1/Channels/{Sid}
func (c *Client) FetchChannel(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchChannelResponse, error) {
	req, err := newFetchChannelRequest(c.BaseURL, sid)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &FetchChannelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FlexV1Channel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// FetchConfiguration: GET /v1/Configuration

type FetchConfigurationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FlexV1Configuration
}

// Status returns HTTPResponse.Status
func (r FetchConfigurationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchConfigurationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchConfigurationRequest generates requests for FetchConfiguration
func newFetchConfigurationRequest(baseURL *url.URL, params *FetchConfigurationParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathFetchConfiguration)

	q := queryURL.Query()

	if params.UiVersion != nil {
		if err := client.AddQueryParam(q, "UiVersion", *params.UiVersion); err != nil {
			return nil, err
		}
	}

	queryURL.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// FetchConfiguration returns a parsed response.
// GET /v1/Configuration
func (c *Client) FetchConfiguration(ctx context.Context, params *FetchConfigurationParams, reqEditors ...client.RequestEditorFn) (*FetchConfigurationResponse, error) {
	req, err := newFetchConfigurationRequest(c.BaseURL, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &FetchConfigurationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FlexV1Configuration
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// ListFlexFlow: GET /v1/FlexFlows

type ListFlexFlowResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		FlexFlows *[]FlexV1FlexFlow `json:"flex_flows,omitempty"`
		Meta      *struct {
			FirstPageUrl    *string `json:"first_page_url,omitempty"`
			Key             *string `json:"key,omitempty"`
			NextPageUrl     *string `json:"next_page_url,omitempty"`
			Page            *int    `json:"page,omitempty"`
			PageSize        *int    `json:"page_size,omitempty"`
			PreviousPageUrl *string `json:"previous_page_url,omitempty"`
			Url             *string `json:"url,omitempty"`
		} `json:"meta,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListFlexFlowResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFlexFlowResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newListFlexFlowRequest generates requests for ListFlexFlow
func newListFlexFlowRequest(baseURL *url.URL, params *ListFlexFlowParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathListFlexFlow)

	q := queryURL.Query()

	if params.FriendlyName != nil {
		if err := client.AddQueryParam(q, "FriendlyName", *params.FriendlyName); err != nil {
			return nil, err
		}
	}

	if params.PageSize != nil {
		if err := client.AddQueryParam(q, "PageSize", *params.PageSize); err != nil {
			return nil, err
		}
	}

	queryURL.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ListFlexFlow returns a parsed response.
// GET /v1/FlexFlows
func (c *Client) ListFlexFlow(ctx context.Context, params *ListFlexFlowParams, reqEditors ...client.RequestEditorFn) (*ListFlexFlowResponse, error) {
	req, err := newListFlexFlowRequest(c.BaseURL, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &ListFlexFlowResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			FlexFlows *[]FlexV1FlexFlow `json:"flex_flows,omitempty"`
			Meta      *struct {
				FirstPageUrl    *string `json:"first_page_url,omitempty"`
				Key             *string `json:"key,omitempty"`
				NextPageUrl     *string `json:"next_page_url,omitempty"`
				Page            *int    `json:"page,omitempty"`
				PageSize        *int    `json:"page_size,omitempty"`
				PreviousPageUrl *string `json:"previous_page_url,omitempty"`
				Url             *string `json:"url,omitempty"`
			} `json:"meta,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// CreateFlexFlow: POST /v1/FlexFlows

type CreateFlexFlowResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *FlexV1FlexFlow
}

// Status returns HTTPResponse.Status
func (r CreateFlexFlowResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateFlexFlowResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newCreateFlexFlowRequestWithBody generates requests for CreateFlexFlow with any type of body
func newCreateFlexFlowRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathCreateFlexFlow)

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// CreateFlexFlowWithBody returns a parsed response.
// POST /v1/FlexFlows
func (c *Client) CreateFlexFlowWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateFlexFlowResponse, error) {
	req, err := newCreateFlexFlowRequestWithBody(c.BaseURL, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &CreateFlexFlowResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest FlexV1FlexFlow
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}

// DeleteFlexFlow: DELETE /v1/FlexFlows/{Sid}

type DeleteFlexFlowResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteFlexFlowResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteFlexFlowResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newDeleteFlexFlowRequest generates requests for DeleteFlexFlow
func newDeleteFlexFlowRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathDeleteFlexFlowFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// DeleteFlexFlow returns a parsed response.
// DELETE /v1/FlexFlows/{Sid}
func (c *Client) DeleteFlexFlow(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteFlexFlowResponse, error) {
	req, err := newDeleteFlexFlowRequest(c.BaseURL, sid)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &DeleteFlexFlowResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// FetchFlexFlow: GET /v1/FlexFlows/{Sid}

type FetchFlexFlowResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FlexV1FlexFlow
}

// Status returns HTTPResponse.Status
func (r FetchFlexFlowResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchFlexFlowResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchFlexFlowRequest generates requests for FetchFlexFlow
func newFetchFlexFlowRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchFlexFlowFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// FetchFlexFlow returns a parsed response.
// GET /v1/FlexFlows/{Sid}
func (c *Client) FetchFlexFlow(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchFlexFlowResponse, error) {
	req, err := newFetchFlexFlowRequest(c.BaseURL, sid)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &FetchFlexFlowResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FlexV1FlexFlow
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// UpdateFlexFlow: POST /v1/FlexFlows/{Sid}

type UpdateFlexFlowResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FlexV1FlexFlow
}

// Status returns HTTPResponse.Status
func (r UpdateFlexFlowResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateFlexFlowResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newUpdateFlexFlowRequestWithBody generates requests for UpdateFlexFlow with any type of body
func newUpdateFlexFlowRequestWithBody(baseURL *url.URL, sid string, contentType string, body io.Reader) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathUpdateFlexFlowFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// UpdateFlexFlowWithBody returns a parsed response.
// POST /v1/FlexFlows/{Sid}
func (c *Client) UpdateFlexFlowWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateFlexFlowResponse, error) {
	req, err := newUpdateFlexFlowRequestWithBody(c.BaseURL, sid, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &UpdateFlexFlowResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FlexV1FlexFlow
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// ListWebChannel: GET /v1/WebChannels

type ListWebChannelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		FlexChatChannels *[]FlexV1WebChannel `json:"flex_chat_channels,omitempty"`
		Meta             *struct {
			FirstPageUrl    *string `json:"first_page_url,omitempty"`
			Key             *string `json:"key,omitempty"`
			NextPageUrl     *string `json:"next_page_url,omitempty"`
			Page            *int    `json:"page,omitempty"`
			PageSize        *int    `json:"page_size,omitempty"`
			PreviousPageUrl *string `json:"previous_page_url,omitempty"`
			Url             *string `json:"url,omitempty"`
		} `json:"meta,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListWebChannelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListWebChannelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newListWebChannelRequest generates requests for ListWebChannel
func newListWebChannelRequest(baseURL *url.URL, params *ListWebChannelParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathListWebChannel)

	q := queryURL.Query()

	if params.PageSize != nil {
		if err := client.AddQueryParam(q, "PageSize", *params.PageSize); err != nil {
			return nil, err
		}
	}

	queryURL.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ListWebChannel returns a parsed response.
// GET /v1/WebChannels
func (c *Client) ListWebChannel(ctx context.Context, params *ListWebChannelParams, reqEditors ...client.RequestEditorFn) (*ListWebChannelResponse, error) {
	req, err := newListWebChannelRequest(c.BaseURL, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &ListWebChannelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			FlexChatChannels *[]FlexV1WebChannel `json:"flex_chat_channels,omitempty"`
			Meta             *struct {
				FirstPageUrl    *string `json:"first_page_url,omitempty"`
				Key             *string `json:"key,omitempty"`
				NextPageUrl     *string `json:"next_page_url,omitempty"`
				Page            *int    `json:"page,omitempty"`
				PageSize        *int    `json:"page_size,omitempty"`
				PreviousPageUrl *string `json:"previous_page_url,omitempty"`
				Url             *string `json:"url,omitempty"`
			} `json:"meta,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// CreateWebChannel: POST /v1/WebChannels

type CreateWebChannelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *FlexV1WebChannel
}

// Status returns HTTPResponse.Status
func (r CreateWebChannelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateWebChannelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newCreateWebChannelRequestWithBody generates requests for CreateWebChannel with any type of body
func newCreateWebChannelRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathCreateWebChannel)

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// CreateWebChannelWithBody returns a parsed response.
// POST /v1/WebChannels
func (c *Client) CreateWebChannelWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateWebChannelResponse, error) {
	req, err := newCreateWebChannelRequestWithBody(c.BaseURL, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &CreateWebChannelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest FlexV1WebChannel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}

// DeleteWebChannel: DELETE /v1/WebChannels/{Sid}

type DeleteWebChannelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteWebChannelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteWebChannelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newDeleteWebChannelRequest generates requests for DeleteWebChannel
func newDeleteWebChannelRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathDeleteWebChannelFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// DeleteWebChannel returns a parsed response.
// DELETE /v1/WebChannels/{Sid}
func (c *Client) DeleteWebChannel(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteWebChannelResponse, error) {
	req, err := newDeleteWebChannelRequest(c.BaseURL, sid)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &DeleteWebChannelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// FetchWebChannel: GET /v1/WebChannels/{Sid}

type FetchWebChannelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FlexV1WebChannel
}

// Status returns HTTPResponse.Status
func (r FetchWebChannelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchWebChannelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchWebChannelRequest generates requests for FetchWebChannel
func newFetchWebChannelRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchWebChannelFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// FetchWebChannel returns a parsed response.
// GET /v1/WebChannels/{Sid}
func (c *Client) FetchWebChannel(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchWebChannelResponse, error) {
	req, err := newFetchWebChannelRequest(c.BaseURL, sid)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &FetchWebChannelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FlexV1WebChannel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// UpdateWebChannel: POST /v1/WebChannels/{Sid}

type UpdateWebChannelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FlexV1WebChannel
}

// Status returns HTTPResponse.Status
func (r UpdateWebChannelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateWebChannelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newUpdateWebChannelRequestWithBody generates requests for UpdateWebChannel with any type of body
func newUpdateWebChannelRequestWithBody(baseURL *url.URL, sid string, contentType string, body io.Reader) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathUpdateWebChannelFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// UpdateWebChannelWithBody returns a parsed response.
// POST /v1/WebChannels/{Sid}
func (c *Client) UpdateWebChannelWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateWebChannelResponse, error) {
	req, err := newUpdateWebChannelRequestWithBody(c.BaseURL, sid, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &UpdateWebChannelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FlexV1WebChannel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
