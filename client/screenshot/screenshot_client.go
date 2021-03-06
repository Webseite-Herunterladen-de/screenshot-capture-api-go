// Code generated by go-swagger; DO NOT EDIT.

package screenshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new screenshot API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for screenshot API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CaptureScreenshotAuthenticated(params *CaptureScreenshotAuthenticatedParams, writer io.Writer, opts ...ClientOption) (*CaptureScreenshotAuthenticatedOK, error)

	CaptureScreenshotUnauthenticated(params *CaptureScreenshotUnauthenticatedParams, writer io.Writer, opts ...ClientOption) (*CaptureScreenshotUnauthenticatedOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CaptureScreenshotAuthenticated Screenshot-Capture-API.com Screenshot Capture is a very simple but powerful screenshot API that anyone can easily use to create pixel-perfect website screenshots. It always uses a recent version of Chrome to ensure that all modern web features are fully supported and rendering is exactly as your customers would expect.
*/
func (a *Client) CaptureScreenshotAuthenticated(params *CaptureScreenshotAuthenticatedParams, writer io.Writer, opts ...ClientOption) (*CaptureScreenshotAuthenticatedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCaptureScreenshotAuthenticatedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "captureScreenshotAuthenticated",
		Method:             "GET",
		PathPattern:        "/capture/{token}/{hash}",
		ProducesMediaTypes: []string{"application/json", "application/pdf", "image/jpeg", "image/png", "image/webp"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CaptureScreenshotAuthenticatedReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CaptureScreenshotAuthenticatedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CaptureScreenshotAuthenticatedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CaptureScreenshotUnauthenticated Screenshot-Capture-API.com Screenshot Capture is a very simple but powerful screenshot API that anyone can easily use to create pixel-perfect website screenshots. It always uses a recent version of Chrome to ensure that all modern web features are fully supported and rendering is exactly as your customers would expect.
*/
func (a *Client) CaptureScreenshotUnauthenticated(params *CaptureScreenshotUnauthenticatedParams, writer io.Writer, opts ...ClientOption) (*CaptureScreenshotUnauthenticatedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCaptureScreenshotUnauthenticatedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "captureScreenshotUnauthenticated",
		Method:             "GET",
		PathPattern:        "/capture/{token}",
		ProducesMediaTypes: []string{"application/json", "application/pdf", "image/jpeg", "image/png", "image/webp"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CaptureScreenshotUnauthenticatedReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CaptureScreenshotUnauthenticatedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CaptureScreenshotUnauthenticatedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
