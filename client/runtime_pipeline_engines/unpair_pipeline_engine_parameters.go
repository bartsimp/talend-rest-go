// Code generated by go-swagger; DO NOT EDIT.

package runtime_pipeline_engines

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

// NewUnpairPipelineEngineParams creates a new UnpairPipelineEngineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnpairPipelineEngineParams() *UnpairPipelineEngineParams {
	return &UnpairPipelineEngineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnpairPipelineEngineParamsWithTimeout creates a new UnpairPipelineEngineParams object
// with the ability to set a timeout on a request.
func NewUnpairPipelineEngineParamsWithTimeout(timeout time.Duration) *UnpairPipelineEngineParams {
	return &UnpairPipelineEngineParams{
		timeout: timeout,
	}
}

// NewUnpairPipelineEngineParamsWithContext creates a new UnpairPipelineEngineParams object
// with the ability to set a context for a request.
func NewUnpairPipelineEngineParamsWithContext(ctx context.Context) *UnpairPipelineEngineParams {
	return &UnpairPipelineEngineParams{
		Context: ctx,
	}
}

// NewUnpairPipelineEngineParamsWithHTTPClient creates a new UnpairPipelineEngineParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnpairPipelineEngineParamsWithHTTPClient(client *http.Client) *UnpairPipelineEngineParams {
	return &UnpairPipelineEngineParams{
		HTTPClient: client,
	}
}

/*
UnpairPipelineEngineParams contains all the parameters to send to the API endpoint

	for the unpair pipeline engine operation.

	Typically these are written to a http.Request.
*/
type UnpairPipelineEngineParams struct {

	/* EngineID.

	   pipeline engine id
	*/
	EngineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unpair pipeline engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnpairPipelineEngineParams) WithDefaults() *UnpairPipelineEngineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unpair pipeline engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnpairPipelineEngineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unpair pipeline engine params
func (o *UnpairPipelineEngineParams) WithTimeout(timeout time.Duration) *UnpairPipelineEngineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unpair pipeline engine params
func (o *UnpairPipelineEngineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unpair pipeline engine params
func (o *UnpairPipelineEngineParams) WithContext(ctx context.Context) *UnpairPipelineEngineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unpair pipeline engine params
func (o *UnpairPipelineEngineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unpair pipeline engine params
func (o *UnpairPipelineEngineParams) WithHTTPClient(client *http.Client) *UnpairPipelineEngineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unpair pipeline engine params
func (o *UnpairPipelineEngineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEngineID adds the engineID to the unpair pipeline engine params
func (o *UnpairPipelineEngineParams) WithEngineID(engineID string) *UnpairPipelineEngineParams {
	o.SetEngineID(engineID)
	return o
}

// SetEngineID adds the engineId to the unpair pipeline engine params
func (o *UnpairPipelineEngineParams) SetEngineID(engineID string) {
	o.EngineID = engineID
}

// WriteToRequest writes these params to a swagger request
func (o *UnpairPipelineEngineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param engineId
	if err := r.SetPathParam("engineId", o.EngineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}