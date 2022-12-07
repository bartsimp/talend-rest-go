// Code generated by go-swagger; DO NOT EDIT.

package promotions_executions

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

// NewGetPromotionExecutionStatusParams creates a new GetPromotionExecutionStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPromotionExecutionStatusParams() *GetPromotionExecutionStatusParams {
	return &GetPromotionExecutionStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPromotionExecutionStatusParamsWithTimeout creates a new GetPromotionExecutionStatusParams object
// with the ability to set a timeout on a request.
func NewGetPromotionExecutionStatusParamsWithTimeout(timeout time.Duration) *GetPromotionExecutionStatusParams {
	return &GetPromotionExecutionStatusParams{
		timeout: timeout,
	}
}

// NewGetPromotionExecutionStatusParamsWithContext creates a new GetPromotionExecutionStatusParams object
// with the ability to set a context for a request.
func NewGetPromotionExecutionStatusParamsWithContext(ctx context.Context) *GetPromotionExecutionStatusParams {
	return &GetPromotionExecutionStatusParams{
		Context: ctx,
	}
}

// NewGetPromotionExecutionStatusParamsWithHTTPClient creates a new GetPromotionExecutionStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPromotionExecutionStatusParamsWithHTTPClient(client *http.Client) *GetPromotionExecutionStatusParams {
	return &GetPromotionExecutionStatusParams{
		HTTPClient: client,
	}
}

/*
GetPromotionExecutionStatusParams contains all the parameters to send to the API endpoint

	for the get promotion execution status operation.

	Typically these are written to a http.Request.
*/
type GetPromotionExecutionStatusParams struct {

	/* ExecutionID.

	   Execution ID
	*/
	ExecutionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get promotion execution status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPromotionExecutionStatusParams) WithDefaults() *GetPromotionExecutionStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get promotion execution status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPromotionExecutionStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get promotion execution status params
func (o *GetPromotionExecutionStatusParams) WithTimeout(timeout time.Duration) *GetPromotionExecutionStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get promotion execution status params
func (o *GetPromotionExecutionStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get promotion execution status params
func (o *GetPromotionExecutionStatusParams) WithContext(ctx context.Context) *GetPromotionExecutionStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get promotion execution status params
func (o *GetPromotionExecutionStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get promotion execution status params
func (o *GetPromotionExecutionStatusParams) WithHTTPClient(client *http.Client) *GetPromotionExecutionStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get promotion execution status params
func (o *GetPromotionExecutionStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get promotion execution status params
func (o *GetPromotionExecutionStatusParams) WithExecutionID(executionID string) *GetPromotionExecutionStatusParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get promotion execution status params
func (o *GetPromotionExecutionStatusParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPromotionExecutionStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param executionId
	if err := r.SetPathParam("executionId", o.ExecutionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}