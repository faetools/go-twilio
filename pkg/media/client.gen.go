// Package media provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package media

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
	opPathFetchMediaProcessorFormat               = "./v1/MediaProcessors/%s"
	opPathUpdateMediaProcessorFormat              = "./v1/MediaProcessors/%s"
	opPathDeleteMediaRecordingFormat              = "./v1/MediaRecordings/%s"
	opPathFetchMediaRecordingFormat               = "./v1/MediaRecordings/%s"
	opPathFetchPlayerStreamerFormat               = "./v1/PlayerStreamers/%s"
	opPathUpdatePlayerStreamerFormat              = "./v1/PlayerStreamers/%s"
	opPathFetchPlayerStreamerPlaybackGrantFormat  = "./v1/PlayerStreamers/%s/PlaybackGrant"
	opPathCreatePlayerStreamerPlaybackGrantFormat = "./v1/PlayerStreamers/%s/PlaybackGrant"
)

var (
	opPathListMediaProcessor   = client.MustParseURL("./v1/MediaProcessors")
	opPathCreateMediaProcessor = client.MustParseURL("./v1/MediaProcessors")
	opPathListMediaRecording   = client.MustParseURL("./v1/MediaRecordings")
	opPathListPlayerStreamer   = client.MustParseURL("./v1/PlayerStreamers")
	opPathCreatePlayerStreamer = client.MustParseURL("./v1/PlayerStreamers")
)

// ClientInterface interface specification for the client.
type ClientInterface interface {
	// ListMediaProcessor request
	ListMediaProcessor(ctx context.Context, params *ListMediaProcessorParams, reqEditors ...client.RequestEditorFn) (*ListMediaProcessorResponse, error)

	// CreateMediaProcessor request with any body
	CreateMediaProcessorWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateMediaProcessorResponse, error)

	// FetchMediaProcessor request
	FetchMediaProcessor(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchMediaProcessorResponse, error)

	// UpdateMediaProcessor request with any body
	UpdateMediaProcessorWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateMediaProcessorResponse, error)

	// ListMediaRecording request
	ListMediaRecording(ctx context.Context, params *ListMediaRecordingParams, reqEditors ...client.RequestEditorFn) (*ListMediaRecordingResponse, error)

	// DeleteMediaRecording request
	DeleteMediaRecording(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteMediaRecordingResponse, error)

	// FetchMediaRecording request
	FetchMediaRecording(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchMediaRecordingResponse, error)

	// ListPlayerStreamer request
	ListPlayerStreamer(ctx context.Context, params *ListPlayerStreamerParams, reqEditors ...client.RequestEditorFn) (*ListPlayerStreamerResponse, error)

	// CreatePlayerStreamer request with any body
	CreatePlayerStreamerWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreatePlayerStreamerResponse, error)

	// FetchPlayerStreamer request
	FetchPlayerStreamer(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchPlayerStreamerResponse, error)

	// UpdatePlayerStreamer request with any body
	UpdatePlayerStreamerWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdatePlayerStreamerResponse, error)

	// FetchPlayerStreamerPlaybackGrant request
	FetchPlayerStreamerPlaybackGrant(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchPlayerStreamerPlaybackGrantResponse, error)

	// CreatePlayerStreamerPlaybackGrant request with any body
	CreatePlayerStreamerPlaybackGrantWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreatePlayerStreamerPlaybackGrantResponse, error)
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

// ListMediaProcessor: GET /v1/MediaProcessors

type ListMediaProcessorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		MediaProcessors *[]MediaV1MediaProcessor `json:"media_processors,omitempty"`
		Meta            *struct {
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
func (r ListMediaProcessorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListMediaProcessorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newListMediaProcessorRequest generates requests for ListMediaProcessor
func newListMediaProcessorRequest(baseURL *url.URL, params *ListMediaProcessorParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathListMediaProcessor)

	q := queryURL.Query()

	if params.Order != nil {
		if err := client.AddQueryParam(q, "Order", *params.Order); err != nil {
			return nil, err
		}
	}

	if params.Status != nil {
		if err := client.AddQueryParam(q, "Status", *params.Status); err != nil {
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

// ListMediaProcessor returns a parsed response.
// GET /v1/MediaProcessors
func (c *Client) ListMediaProcessor(ctx context.Context, params *ListMediaProcessorParams, reqEditors ...client.RequestEditorFn) (*ListMediaProcessorResponse, error) {
	req, err := newListMediaProcessorRequest(c.BaseURL, params)
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

	response := &ListMediaProcessorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			MediaProcessors *[]MediaV1MediaProcessor `json:"media_processors,omitempty"`
			Meta            *struct {
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

// CreateMediaProcessor: POST /v1/MediaProcessors

type CreateMediaProcessorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *MediaV1MediaProcessor
}

// Status returns HTTPResponse.Status
func (r CreateMediaProcessorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateMediaProcessorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newCreateMediaProcessorRequestWithBody generates requests for CreateMediaProcessor with any type of body
func newCreateMediaProcessorRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathCreateMediaProcessor)

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// CreateMediaProcessorWithBody returns a parsed response.
// POST /v1/MediaProcessors
func (c *Client) CreateMediaProcessorWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateMediaProcessorResponse, error) {
	req, err := newCreateMediaProcessorRequestWithBody(c.BaseURL, contentType, body)
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

	response := &CreateMediaProcessorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest MediaV1MediaProcessor
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}

// FetchMediaProcessor: GET /v1/MediaProcessors/{Sid}

type FetchMediaProcessorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MediaV1MediaProcessor
}

// Status returns HTTPResponse.Status
func (r FetchMediaProcessorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchMediaProcessorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchMediaProcessorRequest generates requests for FetchMediaProcessor
func newFetchMediaProcessorRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchMediaProcessorFormat, pathParam0)

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

// FetchMediaProcessor returns a parsed response.
// GET /v1/MediaProcessors/{Sid}
func (c *Client) FetchMediaProcessor(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchMediaProcessorResponse, error) {
	req, err := newFetchMediaProcessorRequest(c.BaseURL, sid)
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

	response := &FetchMediaProcessorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MediaV1MediaProcessor
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// UpdateMediaProcessor: POST /v1/MediaProcessors/{Sid}

type UpdateMediaProcessorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MediaV1MediaProcessor
}

// Status returns HTTPResponse.Status
func (r UpdateMediaProcessorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateMediaProcessorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newUpdateMediaProcessorRequestWithBody generates requests for UpdateMediaProcessor with any type of body
func newUpdateMediaProcessorRequestWithBody(baseURL *url.URL, sid string, contentType string, body io.Reader) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathUpdateMediaProcessorFormat, pathParam0)

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

// UpdateMediaProcessorWithBody returns a parsed response.
// POST /v1/MediaProcessors/{Sid}
func (c *Client) UpdateMediaProcessorWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateMediaProcessorResponse, error) {
	req, err := newUpdateMediaProcessorRequestWithBody(c.BaseURL, sid, contentType, body)
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

	response := &UpdateMediaProcessorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MediaV1MediaProcessor
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// ListMediaRecording: GET /v1/MediaRecordings

type ListMediaRecordingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		MediaRecordings *[]MediaV1MediaRecording `json:"media_recordings,omitempty"`
		Meta            *struct {
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
func (r ListMediaRecordingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListMediaRecordingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newListMediaRecordingRequest generates requests for ListMediaRecording
func newListMediaRecordingRequest(baseURL *url.URL, params *ListMediaRecordingParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathListMediaRecording)

	q := queryURL.Query()

	if params.Order != nil {
		if err := client.AddQueryParam(q, "Order", *params.Order); err != nil {
			return nil, err
		}
	}

	if params.Status != nil {
		if err := client.AddQueryParam(q, "Status", *params.Status); err != nil {
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

// ListMediaRecording returns a parsed response.
// GET /v1/MediaRecordings
func (c *Client) ListMediaRecording(ctx context.Context, params *ListMediaRecordingParams, reqEditors ...client.RequestEditorFn) (*ListMediaRecordingResponse, error) {
	req, err := newListMediaRecordingRequest(c.BaseURL, params)
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

	response := &ListMediaRecordingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			MediaRecordings *[]MediaV1MediaRecording `json:"media_recordings,omitempty"`
			Meta            *struct {
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

// DeleteMediaRecording: DELETE /v1/MediaRecordings/{Sid}

type DeleteMediaRecordingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteMediaRecordingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteMediaRecordingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newDeleteMediaRecordingRequest generates requests for DeleteMediaRecording
func newDeleteMediaRecordingRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathDeleteMediaRecordingFormat, pathParam0)

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

// DeleteMediaRecording returns a parsed response.
// DELETE /v1/MediaRecordings/{Sid}
func (c *Client) DeleteMediaRecording(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteMediaRecordingResponse, error) {
	req, err := newDeleteMediaRecordingRequest(c.BaseURL, sid)
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

	response := &DeleteMediaRecordingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// FetchMediaRecording: GET /v1/MediaRecordings/{Sid}

type FetchMediaRecordingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MediaV1MediaRecording
}

// Status returns HTTPResponse.Status
func (r FetchMediaRecordingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchMediaRecordingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchMediaRecordingRequest generates requests for FetchMediaRecording
func newFetchMediaRecordingRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchMediaRecordingFormat, pathParam0)

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

// FetchMediaRecording returns a parsed response.
// GET /v1/MediaRecordings/{Sid}
func (c *Client) FetchMediaRecording(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchMediaRecordingResponse, error) {
	req, err := newFetchMediaRecordingRequest(c.BaseURL, sid)
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

	response := &FetchMediaRecordingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MediaV1MediaRecording
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// ListPlayerStreamer: GET /v1/PlayerStreamers

type ListPlayerStreamerResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Meta *struct {
			FirstPageUrl    *string `json:"first_page_url,omitempty"`
			Key             *string `json:"key,omitempty"`
			NextPageUrl     *string `json:"next_page_url,omitempty"`
			Page            *int    `json:"page,omitempty"`
			PageSize        *int    `json:"page_size,omitempty"`
			PreviousPageUrl *string `json:"previous_page_url,omitempty"`
			Url             *string `json:"url,omitempty"`
		} `json:"meta,omitempty"`
		PlayerStreamers *[]MediaV1PlayerStreamer `json:"player_streamers,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListPlayerStreamerResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListPlayerStreamerResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newListPlayerStreamerRequest generates requests for ListPlayerStreamer
func newListPlayerStreamerRequest(baseURL *url.URL, params *ListPlayerStreamerParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathListPlayerStreamer)

	q := queryURL.Query()

	if params.Order != nil {
		if err := client.AddQueryParam(q, "Order", *params.Order); err != nil {
			return nil, err
		}
	}

	if params.Status != nil {
		if err := client.AddQueryParam(q, "Status", *params.Status); err != nil {
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

// ListPlayerStreamer returns a parsed response.
// GET /v1/PlayerStreamers
func (c *Client) ListPlayerStreamer(ctx context.Context, params *ListPlayerStreamerParams, reqEditors ...client.RequestEditorFn) (*ListPlayerStreamerResponse, error) {
	req, err := newListPlayerStreamerRequest(c.BaseURL, params)
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

	response := &ListPlayerStreamerResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Meta *struct {
				FirstPageUrl    *string `json:"first_page_url,omitempty"`
				Key             *string `json:"key,omitempty"`
				NextPageUrl     *string `json:"next_page_url,omitempty"`
				Page            *int    `json:"page,omitempty"`
				PageSize        *int    `json:"page_size,omitempty"`
				PreviousPageUrl *string `json:"previous_page_url,omitempty"`
				Url             *string `json:"url,omitempty"`
			} `json:"meta,omitempty"`
			PlayerStreamers *[]MediaV1PlayerStreamer `json:"player_streamers,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// CreatePlayerStreamer: POST /v1/PlayerStreamers

type CreatePlayerStreamerResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *MediaV1PlayerStreamer
}

// Status returns HTTPResponse.Status
func (r CreatePlayerStreamerResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreatePlayerStreamerResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newCreatePlayerStreamerRequestWithBody generates requests for CreatePlayerStreamer with any type of body
func newCreatePlayerStreamerRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathCreatePlayerStreamer)

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// CreatePlayerStreamerWithBody returns a parsed response.
// POST /v1/PlayerStreamers
func (c *Client) CreatePlayerStreamerWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreatePlayerStreamerResponse, error) {
	req, err := newCreatePlayerStreamerRequestWithBody(c.BaseURL, contentType, body)
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

	response := &CreatePlayerStreamerResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest MediaV1PlayerStreamer
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}

// FetchPlayerStreamer: GET /v1/PlayerStreamers/{Sid}

type FetchPlayerStreamerResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MediaV1PlayerStreamer
}

// Status returns HTTPResponse.Status
func (r FetchPlayerStreamerResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchPlayerStreamerResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchPlayerStreamerRequest generates requests for FetchPlayerStreamer
func newFetchPlayerStreamerRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchPlayerStreamerFormat, pathParam0)

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

// FetchPlayerStreamer returns a parsed response.
// GET /v1/PlayerStreamers/{Sid}
func (c *Client) FetchPlayerStreamer(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchPlayerStreamerResponse, error) {
	req, err := newFetchPlayerStreamerRequest(c.BaseURL, sid)
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

	response := &FetchPlayerStreamerResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MediaV1PlayerStreamer
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// UpdatePlayerStreamer: POST /v1/PlayerStreamers/{Sid}

type UpdatePlayerStreamerResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MediaV1PlayerStreamer
}

// Status returns HTTPResponse.Status
func (r UpdatePlayerStreamerResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdatePlayerStreamerResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newUpdatePlayerStreamerRequestWithBody generates requests for UpdatePlayerStreamer with any type of body
func newUpdatePlayerStreamerRequestWithBody(baseURL *url.URL, sid string, contentType string, body io.Reader) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathUpdatePlayerStreamerFormat, pathParam0)

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

// UpdatePlayerStreamerWithBody returns a parsed response.
// POST /v1/PlayerStreamers/{Sid}
func (c *Client) UpdatePlayerStreamerWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdatePlayerStreamerResponse, error) {
	req, err := newUpdatePlayerStreamerRequestWithBody(c.BaseURL, sid, contentType, body)
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

	response := &UpdatePlayerStreamerResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MediaV1PlayerStreamer
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// FetchPlayerStreamerPlaybackGrant: GET /v1/PlayerStreamers/{Sid}/PlaybackGrant

type FetchPlayerStreamerPlaybackGrantResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MediaV1PlayerStreamerPlayerStreamerPlaybackGrant
}

// Status returns HTTPResponse.Status
func (r FetchPlayerStreamerPlaybackGrantResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchPlayerStreamerPlaybackGrantResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchPlayerStreamerPlaybackGrantRequest generates requests for FetchPlayerStreamerPlaybackGrant
func newFetchPlayerStreamerPlaybackGrantRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchPlayerStreamerPlaybackGrantFormat, pathParam0)

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

// FetchPlayerStreamerPlaybackGrant returns a parsed response.
// GET /v1/PlayerStreamers/{Sid}/PlaybackGrant
func (c *Client) FetchPlayerStreamerPlaybackGrant(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchPlayerStreamerPlaybackGrantResponse, error) {
	req, err := newFetchPlayerStreamerPlaybackGrantRequest(c.BaseURL, sid)
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

	response := &FetchPlayerStreamerPlaybackGrantResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MediaV1PlayerStreamerPlayerStreamerPlaybackGrant
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// CreatePlayerStreamerPlaybackGrant: POST /v1/PlayerStreamers/{Sid}/PlaybackGrant

type CreatePlayerStreamerPlaybackGrantResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *MediaV1PlayerStreamerPlayerStreamerPlaybackGrant
}

// Status returns HTTPResponse.Status
func (r CreatePlayerStreamerPlaybackGrantResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreatePlayerStreamerPlaybackGrantResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newCreatePlayerStreamerPlaybackGrantRequestWithBody generates requests for CreatePlayerStreamerPlaybackGrant with any type of body
func newCreatePlayerStreamerPlaybackGrantRequestWithBody(baseURL *url.URL, sid string, contentType string, body io.Reader) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathCreatePlayerStreamerPlaybackGrantFormat, pathParam0)

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

// CreatePlayerStreamerPlaybackGrantWithBody returns a parsed response.
// POST /v1/PlayerStreamers/{Sid}/PlaybackGrant
func (c *Client) CreatePlayerStreamerPlaybackGrantWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreatePlayerStreamerPlaybackGrantResponse, error) {
	req, err := newCreatePlayerStreamerPlaybackGrantRequestWithBody(c.BaseURL, sid, contentType, body)
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

	response := &CreatePlayerStreamerPlaybackGrantResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest MediaV1PlayerStreamerPlayerStreamerPlaybackGrant
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}
