// Code generated by go-swagger; DO NOT EDIT.

package plans_executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new plans executions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plans executions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ExecutePlan(params *ExecutePlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExecutePlanCreated, error)

	GetAllStepExecutions(params *GetAllStepExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllStepExecutionsOK, error)

	GetPlanExecutionStatus(params *GetPlanExecutionStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlanExecutionStatusOK, error)

	GetStepExecution(params *GetStepExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStepExecutionOK, error)

	GetStepHandlerExecution(params *GetStepHandlerExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStepHandlerExecutionOK, error)

	ListAvailablePlansExecutions(params *ListAvailablePlansExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAvailablePlansExecutionsOK, error)

	ListPlanExecutions(params *ListPlanExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPlanExecutionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ExecutePlan executes plan

Allows to run a Plan (starting from V2.4: allows also to rerun the plan from a specific step)
*/
func (a *Client) ExecutePlan(params *ExecutePlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExecutePlanCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecutePlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "executePlan",
		Method:             "POST",
		PathPattern:        "/executions/plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecutePlanReader{formats: a.formats},
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
	success, ok := result.(*ExecutePlanCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for executePlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllStepExecutions gets all steps for a plan execution order by designed execution only steps without error handlers

Get Steps executions status
*/
func (a *Client) GetAllStepExecutions(params *GetAllStepExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllStepExecutionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllStepExecutionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllStepExecutions",
		Method:             "GET",
		PathPattern:        "/executions/plans/{planExecutionId}/steps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllStepExecutionsReader{formats: a.formats},
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
	success, ok := result.(*GetAllStepExecutionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllStepExecutions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPlanExecutionStatus gets plan execution status

Get Plan execution status
*/
func (a *Client) GetPlanExecutionStatus(params *GetPlanExecutionStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlanExecutionStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlanExecutionStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPlanExecutionStatus",
		Method:             "GET",
		PathPattern:        "/executions/plans/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPlanExecutionStatusReader{formats: a.formats},
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
	success, ok := result.(*GetPlanExecutionStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPlanExecutionStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetStepExecution gets the step for plan execution without error handlers

Get Step execution status
*/
func (a *Client) GetStepExecution(params *GetStepExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStepExecutionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStepExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStepExecution",
		Method:             "GET",
		PathPattern:        "/executions/plans/{planExecutionId}/steps/{stepExecutionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStepExecutionReader{formats: a.formats},
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
	success, ok := result.(*GetStepExecutionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStepExecution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetStepHandlerExecution gets execution status of the handler for the specified step

Get Step handler execution status
*/
func (a *Client) GetStepHandlerExecution(params *GetStepHandlerExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStepHandlerExecutionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStepHandlerExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStepHandlerExecution",
		Method:             "GET",
		PathPattern:        "/executions/plans/{planExecutionId}/steps/{stepExecutionId}/error-handler",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStepHandlerExecutionReader{formats: a.formats},
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
	success, ok := result.(*GetStepHandlerExecutionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStepHandlerExecution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListAvailablePlansExecutions gets available plans executions

Get available Plans executions
*/
func (a *Client) ListAvailablePlansExecutions(params *ListAvailablePlansExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAvailablePlansExecutionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAvailablePlansExecutionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAvailablePlansExecutions",
		Method:             "GET",
		PathPattern:        "/executables/plans/executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAvailablePlansExecutionsReader{formats: a.formats},
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
	success, ok := result.(*ListAvailablePlansExecutionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAvailablePlansExecutions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListPlanExecutions gets plan executions

Get Plan executions
*/
func (a *Client) ListPlanExecutions(params *ListPlanExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPlanExecutionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPlanExecutionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPlanExecutions",
		Method:             "GET",
		PathPattern:        "/executables/plans/{planId}/executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPlanExecutionsReader{formats: a.formats},
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
	success, ok := result.(*ListPlanExecutionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPlanExecutions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
