// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new environments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for environments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEnvironment(params *CreateEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEnvironmentCreated, error)

	DeleteEnvironment(params *DeleteEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEnvironmentNoContent, error)

	GetEnvironments(params *GetEnvironmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentsOK, error)

	UpdateEnvironment(params *UpdateEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateEnvironmentNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateEnvironment creates an environment

Create an environment.
*/
func (a *Client) CreateEnvironment(params *CreateEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEnvironmentCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createEnvironment",
		Method:             "POST",
		PathPattern:        "/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*CreateEnvironmentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteEnvironment deletes the environment for a given identifier
*/
func (a *Client) DeleteEnvironment(params *DeleteEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEnvironmentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteEnvironment",
		Method:             "DELETE",
		PathPattern:        "/environments/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*DeleteEnvironmentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEnvironments gets available environments

Get available environments
*/
func (a *Client) GetEnvironments(params *GetEnvironmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEnvironments",
		Method:             "GET",
		PathPattern:        "/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEnvironmentsReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetEnvironmentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateEnvironment updates environment
*/
func (a *Client) UpdateEnvironment(params *UpdateEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateEnvironmentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateEnvironment",
		Method:             "PUT",
		PathPattern:        "/environments/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdateEnvironmentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
