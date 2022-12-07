// Code generated by go-swagger; DO NOT EDIT.

package plans_executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetStepHandlerExecutionParams creates a new GetStepHandlerExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStepHandlerExecutionParams() *GetStepHandlerExecutionParams {
	return &GetStepHandlerExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStepHandlerExecutionParamsWithTimeout creates a new GetStepHandlerExecutionParams object
// with the ability to set a timeout on a request.
func NewGetStepHandlerExecutionParamsWithTimeout(timeout time.Duration) *GetStepHandlerExecutionParams {
	return &GetStepHandlerExecutionParams{
		timeout: timeout,
	}
}

// NewGetStepHandlerExecutionParamsWithContext creates a new GetStepHandlerExecutionParams object
// with the ability to set a context for a request.
func NewGetStepHandlerExecutionParamsWithContext(ctx context.Context) *GetStepHandlerExecutionParams {
	return &GetStepHandlerExecutionParams{
		Context: ctx,
	}
}

// NewGetStepHandlerExecutionParamsWithHTTPClient creates a new GetStepHandlerExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStepHandlerExecutionParamsWithHTTPClient(client *http.Client) *GetStepHandlerExecutionParams {
	return &GetStepHandlerExecutionParams{
		HTTPClient: client,
	}
}

/*
GetStepHandlerExecutionParams contains all the parameters to send to the API endpoint

	for the get step handler execution operation.

	Typically these are written to a http.Request.
*/
type GetStepHandlerExecutionParams struct {

	/* PlanExecutionID.

	   Plan execution ID
	*/
	PlanExecutionID string

	/* StepExecutionID.

	   Step execution ID
	*/
	StepExecutionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get step handler execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStepHandlerExecutionParams) WithDefaults() *GetStepHandlerExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get step handler execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStepHandlerExecutionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get step handler execution params
func (o *GetStepHandlerExecutionParams) WithTimeout(timeout time.Duration) *GetStepHandlerExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get step handler execution params
func (o *GetStepHandlerExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get step handler execution params
func (o *GetStepHandlerExecutionParams) WithContext(ctx context.Context) *GetStepHandlerExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get step handler execution params
func (o *GetStepHandlerExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get step handler execution params
func (o *GetStepHandlerExecutionParams) WithHTTPClient(client *http.Client) *GetStepHandlerExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get step handler execution params
func (o *GetStepHandlerExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanExecutionID adds the planExecutionID to the get step handler execution params
func (o *GetStepHandlerExecutionParams) WithPlanExecutionID(planExecutionID string) *GetStepHandlerExecutionParams {
	o.SetPlanExecutionID(planExecutionID)
	return o
}

// SetPlanExecutionID adds the planExecutionId to the get step handler execution params
func (o *GetStepHandlerExecutionParams) SetPlanExecutionID(planExecutionID string) {
	o.PlanExecutionID = planExecutionID
}

// WithStepExecutionID adds the stepExecutionID to the get step handler execution params
func (o *GetStepHandlerExecutionParams) WithStepExecutionID(stepExecutionID string) *GetStepHandlerExecutionParams {
	o.SetStepExecutionID(stepExecutionID)
	return o
}

// SetStepExecutionID adds the stepExecutionId to the get step handler execution params
func (o *GetStepHandlerExecutionParams) SetStepExecutionID(stepExecutionID string) {
	o.StepExecutionID = stepExecutionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStepHandlerExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param planExecutionId
	if err := r.SetPathParam("planExecutionId", o.PlanExecutionID); err != nil {
		return err
	}

	// path param stepExecutionId
	if err := r.SetPathParam("stepExecutionId", o.StepExecutionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
