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

	"github.com/bartsimp/talend-rest-go/models"
)

// NewExecutePromotionParams creates a new ExecutePromotionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecutePromotionParams() *ExecutePromotionParams {
	return &ExecutePromotionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecutePromotionParamsWithTimeout creates a new ExecutePromotionParams object
// with the ability to set a timeout on a request.
func NewExecutePromotionParamsWithTimeout(timeout time.Duration) *ExecutePromotionParams {
	return &ExecutePromotionParams{
		timeout: timeout,
	}
}

// NewExecutePromotionParamsWithContext creates a new ExecutePromotionParams object
// with the ability to set a context for a request.
func NewExecutePromotionParamsWithContext(ctx context.Context) *ExecutePromotionParams {
	return &ExecutePromotionParams{
		Context: ctx,
	}
}

// NewExecutePromotionParamsWithHTTPClient creates a new ExecutePromotionParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecutePromotionParamsWithHTTPClient(client *http.Client) *ExecutePromotionParams {
	return &ExecutePromotionParams{
		HTTPClient: client,
	}
}

/*
ExecutePromotionParams contains all the parameters to send to the API endpoint

	for the execute promotion operation.

	Typically these are written to a http.Request.
*/
type ExecutePromotionParams struct {

	/* Body.

	   Executable task
	*/
	Body *models.PromotionExecutableTask

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute promotion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutePromotionParams) WithDefaults() *ExecutePromotionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute promotion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutePromotionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execute promotion params
func (o *ExecutePromotionParams) WithTimeout(timeout time.Duration) *ExecutePromotionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute promotion params
func (o *ExecutePromotionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute promotion params
func (o *ExecutePromotionParams) WithContext(ctx context.Context) *ExecutePromotionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute promotion params
func (o *ExecutePromotionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute promotion params
func (o *ExecutePromotionParams) WithHTTPClient(client *http.Client) *ExecutePromotionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute promotion params
func (o *ExecutePromotionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the execute promotion params
func (o *ExecutePromotionParams) WithBody(body *models.PromotionExecutableTask) *ExecutePromotionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the execute promotion params
func (o *ExecutePromotionParams) SetBody(body *models.PromotionExecutableTask) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ExecutePromotionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
