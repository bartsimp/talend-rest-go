// Code generated by go-swagger; DO NOT EDIT.

package artifacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new artifacts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for artifacts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteArtifact(params *DeleteArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteArtifactNoContent, error)

	DeleteArtifactOfVersion(params *DeleteArtifactOfVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteArtifactOfVersionNoContent, error)

	GetArtifact(params *GetArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetArtifactOK, error)

	GetArtifactOfVersion(params *GetArtifactOfVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetArtifactOfVersionOK, error)

	GetArtifactsAvailable(params *GetArtifactsAvailableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetArtifactsAvailableOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteArtifact deletes artifact by id

Delete Artifact by id
*/
func (a *Client) DeleteArtifact(params *DeleteArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteArtifactNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteArtifactParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteArtifact",
		Method:             "DELETE",
		PathPattern:        "/artifacts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteArtifactReader{formats: a.formats},
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
	success, ok := result.(*DeleteArtifactNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteArtifact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteArtifactOfVersion deletes artifact of a specified version

Delete Artifact of a specified version
*/
func (a *Client) DeleteArtifactOfVersion(params *DeleteArtifactOfVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteArtifactOfVersionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteArtifactOfVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteArtifactOfVersion",
		Method:             "DELETE",
		PathPattern:        "/artifacts/{id}/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteArtifactOfVersionReader{formats: a.formats},
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
	success, ok := result.(*DeleteArtifactOfVersionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteArtifactOfVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetArtifact gets artifact by id

Get Artifact by id
*/
func (a *Client) GetArtifact(params *GetArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetArtifactParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getArtifact",
		Method:             "GET",
		PathPattern:        "/artifacts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetArtifactReader{formats: a.formats},
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
	success, ok := result.(*GetArtifactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getArtifact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetArtifactOfVersion gets artifact of a specified version

Get Artifact of a specified version
*/
func (a *Client) GetArtifactOfVersion(params *GetArtifactOfVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetArtifactOfVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetArtifactOfVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getArtifactOfVersion",
		Method:             "GET",
		PathPattern:        "/artifacts/{id}/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetArtifactOfVersionReader{formats: a.formats},
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
	success, ok := result.(*GetArtifactOfVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getArtifactOfVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetArtifactsAvailable gets available artifacts

Get available Artifacts
*/
func (a *Client) GetArtifactsAvailable(params *GetArtifactsAvailableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetArtifactsAvailableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetArtifactsAvailableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getArtifactsAvailable",
		Method:             "GET",
		PathPattern:        "/artifacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetArtifactsAvailableReader{formats: a.formats},
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
	success, ok := result.(*GetArtifactsAvailableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getArtifactsAvailable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}