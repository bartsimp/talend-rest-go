// Code generated by go-swagger; DO NOT EDIT.

package runtime_clusters_run_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new runtime clusters run profiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for runtime clusters run profiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateClusterRunProfile(params *CreateClusterRunProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	DeleteClusterRunProfile(params *DeleteClusterRunProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClusterRunProfileNoContent, error)

	GetClusterRunProfile(params *GetClusterRunProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterRunProfileOK, error)

	GetClusterRunProfiles(params *GetClusterRunProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterRunProfilesOK, error)

	UpdateClusterRunProfile(params *UpdateClusterRunProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateClusterRunProfileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateClusterRunProfile creates cluster run profile

Create cluster run profile
*/
func (a *Client) CreateClusterRunProfile(params *CreateClusterRunProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterRunProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createClusterRunProfile",
		Method:             "POST",
		PathPattern:        "/runtimes/remote-engine-clusters/{clusterId}/run-profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClusterRunProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteClusterRunProfile deletes cluster run profile

Delete cluster run profile
*/
func (a *Client) DeleteClusterRunProfile(params *DeleteClusterRunProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClusterRunProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterRunProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteClusterRunProfile",
		Method:             "DELETE",
		PathPattern:        "/runtimes/remote-engine-clusters/{clusterId}/run-profiles/{runProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteClusterRunProfileReader{formats: a.formats},
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
	success, ok := result.(*DeleteClusterRunProfileNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteClusterRunProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetClusterRunProfile gets cluster run profile

Get cluster run profile
*/
func (a *Client) GetClusterRunProfile(params *GetClusterRunProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterRunProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterRunProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClusterRunProfile",
		Method:             "GET",
		PathPattern:        "/runtimes/remote-engine-clusters/{clusterId}/run-profiles/{runProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterRunProfileReader{formats: a.formats},
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
	success, ok := result.(*GetClusterRunProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClusterRunProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetClusterRunProfiles gets cluster run profiles

Get cluster run profiles
*/
func (a *Client) GetClusterRunProfiles(params *GetClusterRunProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterRunProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterRunProfilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClusterRunProfiles",
		Method:             "GET",
		PathPattern:        "/runtimes/remote-engine-clusters/{clusterId}/run-profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterRunProfilesReader{formats: a.formats},
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
	success, ok := result.(*GetClusterRunProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClusterRunProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateClusterRunProfile updates cluster run profile

Update cluster run profile
*/
func (a *Client) UpdateClusterRunProfile(params *UpdateClusterRunProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateClusterRunProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterRunProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateClusterRunProfile",
		Method:             "PUT",
		PathPattern:        "/runtimes/remote-engine-clusters/{clusterId}/run-profiles/{runProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClusterRunProfileReader{formats: a.formats},
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
	success, ok := result.(*UpdateClusterRunProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateClusterRunProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
