// Code generated by go-swagger; DO NOT EDIT.

package executions_logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new executions logs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for executions logs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CheckDownloadTokenStatus(params *CheckDownloadTokenStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CheckDownloadTokenStatusOK, error)

	GenerateFullTaskExecutionLogs(params *GenerateFullTaskExecutionLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GenerateFullTaskExecutionLogsOK, error)

	GetTaskExecutionLog(params *GetTaskExecutionLogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTaskExecutionLogOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CheckDownloadTokenStatus checks download token status

Check download token status, possible status values: READY, IN_PROGRESS. Returns a pre-signed URL valid for one hour once the status is READY
*/
func (a *Client) CheckDownloadTokenStatus(params *CheckDownloadTokenStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CheckDownloadTokenStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckDownloadTokenStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "checkDownloadTokenStatus",
		Method:             "POST",
		PathPattern:        "/executions/{id}/logs/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CheckDownloadTokenStatusReader{formats: a.formats},
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
	success, ok := result.(*CheckDownloadTokenStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for checkDownloadTokenStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GenerateFullTaskExecutionLogs generates full logs for task execution

Generate full logs for Task execution. Please note the customer log download API is asynchronous, after receiving a token, the frequency on which you should check logs file status depends on the size of logs, the status endpoint should be used for this purpose
*/
func (a *Client) GenerateFullTaskExecutionLogs(params *GenerateFullTaskExecutionLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GenerateFullTaskExecutionLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateFullTaskExecutionLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "generateFullTaskExecutionLogs",
		Method:             "POST",
		PathPattern:        "/executions/{id}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenerateFullTaskExecutionLogsReader{formats: a.formats},
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
	success, ok := result.(*GenerateFullTaskExecutionLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for generateFullTaskExecutionLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTaskExecutionLog gets task execution logs

Get Task execution logs
*/
func (a *Client) GetTaskExecutionLog(params *GetTaskExecutionLogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTaskExecutionLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskExecutionLogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTaskExecutionLog",
		Method:             "GET",
		PathPattern:        "/executions/{id}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTaskExecutionLogReader{formats: a.formats},
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
	success, ok := result.(*GetTaskExecutionLogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTaskExecutionLog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}