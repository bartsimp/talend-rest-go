// Code generated by go-swagger; DO NOT EDIT.

package executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new executions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for executions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Execute(params *ExecuteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExecuteCreated, error)

	GetTaskExecutionStatus(params *GetTaskExecutionStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTaskExecutionStatusOK, error)

	StopExecution(params *StopExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopExecutionOK, *StopExecutionNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Execute executes task

Execute Task
*/
func (a *Client) Execute(params *ExecuteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExecuteCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "execute",
		Method:             "POST",
		PathPattern:        "/executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecuteReader{formats: a.formats},
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
	success, ok := result.(*ExecuteCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for execute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTaskExecutionStatus gets task execution status

Get Task execution status
*/
func (a *Client) GetTaskExecutionStatus(params *GetTaskExecutionStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTaskExecutionStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskExecutionStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTaskExecutionStatus",
		Method:             "GET",
		PathPattern:        "/executions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTaskExecutionStatusReader{formats: a.formats},
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
	success, ok := result.(*GetTaskExecutionStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTaskExecutionStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopExecution terminates task execution

Terminate Task execution
*/
func (a *Client) StopExecution(params *StopExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopExecutionOK, *StopExecutionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopExecution",
		Method:             "DELETE",
		PathPattern:        "/executions/{id}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StopExecutionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *StopExecutionOK:
		return value, nil, nil
	case *StopExecutionNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for executions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}