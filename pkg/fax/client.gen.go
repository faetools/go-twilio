// Package fax provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package fax

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
	opPathListFaxMediaFormat   = "./v1/Faxes/%s/Media"
	opPathDeleteFaxMediaFormat = "./v1/Faxes/%s/Media/%s"
	opPathFetchFaxMediaFormat  = "./v1/Faxes/%s/Media/%s"
	opPathDeleteFaxFormat      = "./v1/Faxes/%s"
	opPathFetchFaxFormat       = "./v1/Faxes/%s"
)

var opPathListFax = client.MustParseURL("./v1/Faxes")

// ClientInterface interface specification for the client.
type ClientInterface interface {
	// ListFax request
	ListFax(ctx context.Context, params *ListFaxParams, reqEditors ...client.RequestEditorFn) (*ListFaxResponse, error)

	// ListFaxMedia request
	ListFaxMedia(ctx context.Context, faxSid string, params *ListFaxMediaParams, reqEditors ...client.RequestEditorFn) (*ListFaxMediaResponse, error)

	// DeleteFaxMedia request
	DeleteFaxMedia(ctx context.Context, faxSid string, sid string, reqEditors ...client.RequestEditorFn) (*DeleteFaxMediaResponse, error)

	// FetchFaxMedia request
	FetchFaxMedia(ctx context.Context, faxSid string, sid string, reqEditors ...client.RequestEditorFn) (*FetchFaxMediaResponse, error)

	// DeleteFax request
	DeleteFax(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteFaxResponse, error)

	// FetchFax request
	FetchFax(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchFaxResponse, error)
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

// ListFax: GET /v1/Faxes

type ListFaxResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Faxes *[]FaxV1Fax `json:"faxes,omitempty"`
		Meta  *struct {
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
func (r ListFaxResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFaxResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newListFaxRequest generates requests for ListFax
func newListFaxRequest(baseURL *url.URL, params *ListFaxParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathListFax)

	q := queryURL.Query()

	if params.From != nil {
		if err := client.AddQueryParam(q, "From", *params.From); err != nil {
			return nil, err
		}
	}

	if params.To != nil {
		if err := client.AddQueryParam(q, "To", *params.To); err != nil {
			return nil, err
		}
	}

	if params.DateCreatedOnOrBefore != nil {
		if err := client.AddQueryParam(q, "DateCreatedOnOrBefore", *params.DateCreatedOnOrBefore); err != nil {
			return nil, err
		}
	}

	if params.DateCreatedAfter != nil {
		if err := client.AddQueryParam(q, "DateCreatedAfter", *params.DateCreatedAfter); err != nil {
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

// ListFax returns a parsed response.
// GET /v1/Faxes
func (c *Client) ListFax(ctx context.Context, params *ListFaxParams, reqEditors ...client.RequestEditorFn) (*ListFaxResponse, error) {
	req, err := newListFaxRequest(c.BaseURL, params)
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

	response := &ListFaxResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Faxes *[]FaxV1Fax `json:"faxes,omitempty"`
			Meta  *struct {
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

// ListFaxMedia: GET /v1/Faxes/{FaxSid}/Media

type ListFaxMediaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Media *[]FaxV1FaxFaxMedia `json:"media,omitempty"`
		Meta  *struct {
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
func (r ListFaxMediaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFaxMediaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newListFaxMediaRequest generates requests for ListFaxMedia
func newListFaxMediaRequest(baseURL *url.URL, faxSid string, params *ListFaxMediaParams) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("FaxSid", faxSid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathListFaxMediaFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

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

// ListFaxMedia returns a parsed response.
// GET /v1/Faxes/{FaxSid}/Media
func (c *Client) ListFaxMedia(ctx context.Context, faxSid string, params *ListFaxMediaParams, reqEditors ...client.RequestEditorFn) (*ListFaxMediaResponse, error) {
	req, err := newListFaxMediaRequest(c.BaseURL, faxSid, params)
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

	response := &ListFaxMediaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Media *[]FaxV1FaxFaxMedia `json:"media,omitempty"`
			Meta  *struct {
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

// DeleteFaxMedia: DELETE /v1/Faxes/{FaxSid}/Media/{Sid}

type DeleteFaxMediaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteFaxMediaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteFaxMediaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newDeleteFaxMediaRequest generates requests for DeleteFaxMedia
func newDeleteFaxMediaRequest(baseURL *url.URL, faxSid string, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("FaxSid", faxSid)
	if err != nil {
		return nil, err
	}

	pathParam1, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathDeleteFaxMediaFormat, pathParam0, pathParam1)

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

// DeleteFaxMedia returns a parsed response.
// DELETE /v1/Faxes/{FaxSid}/Media/{Sid}
func (c *Client) DeleteFaxMedia(ctx context.Context, faxSid string, sid string, reqEditors ...client.RequestEditorFn) (*DeleteFaxMediaResponse, error) {
	req, err := newDeleteFaxMediaRequest(c.BaseURL, faxSid, sid)
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

	response := &DeleteFaxMediaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// FetchFaxMedia: GET /v1/Faxes/{FaxSid}/Media/{Sid}

type FetchFaxMediaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FaxV1FaxFaxMedia
}

// Status returns HTTPResponse.Status
func (r FetchFaxMediaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchFaxMediaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchFaxMediaRequest generates requests for FetchFaxMedia
func newFetchFaxMediaRequest(baseURL *url.URL, faxSid string, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("FaxSid", faxSid)
	if err != nil {
		return nil, err
	}

	pathParam1, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchFaxMediaFormat, pathParam0, pathParam1)

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

// FetchFaxMedia returns a parsed response.
// GET /v1/Faxes/{FaxSid}/Media/{Sid}
func (c *Client) FetchFaxMedia(ctx context.Context, faxSid string, sid string, reqEditors ...client.RequestEditorFn) (*FetchFaxMediaResponse, error) {
	req, err := newFetchFaxMediaRequest(c.BaseURL, faxSid, sid)
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

	response := &FetchFaxMediaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FaxV1FaxFaxMedia
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// DeleteFax: DELETE /v1/Faxes/{Sid}

type DeleteFaxResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteFaxResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteFaxResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newDeleteFaxRequest generates requests for DeleteFax
func newDeleteFaxRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathDeleteFaxFormat, pathParam0)

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

// DeleteFax returns a parsed response.
// DELETE /v1/Faxes/{Sid}
func (c *Client) DeleteFax(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*DeleteFaxResponse, error) {
	req, err := newDeleteFaxRequest(c.BaseURL, sid)
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

	response := &DeleteFaxResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// FetchFax: GET /v1/Faxes/{Sid}

type FetchFaxResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FaxV1Fax
}

// Status returns HTTPResponse.Status
func (r FetchFaxResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchFaxResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchFaxRequest generates requests for FetchFax
func newFetchFaxRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchFaxFormat, pathParam0)

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

// FetchFax returns a parsed response.
// GET /v1/Faxes/{Sid}
func (c *Client) FetchFax(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchFaxResponse, error) {
	req, err := newFetchFaxRequest(c.BaseURL, sid)
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

	response := &FetchFaxResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FaxV1Fax
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}