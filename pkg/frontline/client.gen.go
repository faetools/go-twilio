// Package frontline provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package frontline

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
	opPathFetchUserFormat  = "./v1/Users/%s"
	opPathUpdateUserFormat = "./v1/Users/%s"
)

// ClientInterface interface specification for the client.
type ClientInterface interface {
	// FetchUser request
	FetchUser(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchUserResponse, error)

	// UpdateUser request with any body
	UpdateUserWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateUserResponse, error)
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

// FetchUser: GET /v1/Users/{Sid}

type FetchUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FrontlineV1User
}

// Status returns HTTPResponse.Status
func (r FetchUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newFetchUserRequest generates requests for FetchUser
func newFetchUserRequest(baseURL *url.URL, sid string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathFetchUserFormat, pathParam0)

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

// FetchUser returns a parsed response.
// GET /v1/Users/{Sid}
func (c *Client) FetchUser(ctx context.Context, sid string, reqEditors ...client.RequestEditorFn) (*FetchUserResponse, error) {
	req, err := newFetchUserRequest(c.BaseURL, sid)
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

	response := &FetchUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FrontlineV1User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// UpdateUser: POST /v1/Users/{Sid}

type UpdateUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FrontlineV1User
}

// Status returns HTTPResponse.Status
func (r UpdateUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newUpdateUserRequestWithBody generates requests for UpdateUser with any type of body
func newUpdateUserRequestWithBody(baseURL *url.URL, sid string, contentType string, body io.Reader) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("Sid", sid)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathUpdateUserFormat, pathParam0)

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

// UpdateUserWithBody returns a parsed response.
// POST /v1/Users/{Sid}
func (c *Client) UpdateUserWithBody(ctx context.Context, sid string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateUserResponse, error) {
	req, err := newUpdateUserRequestWithBody(c.BaseURL, sid, contentType, body)
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

	response := &UpdateUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FrontlineV1User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
