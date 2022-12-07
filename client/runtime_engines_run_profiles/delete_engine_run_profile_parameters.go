// Code generated by go-swagger; DO NOT EDIT.

package runtime_engines_run_profiles

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

// NewDeleteEngineRunProfileParams creates a new DeleteEngineRunProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEngineRunProfileParams() *DeleteEngineRunProfileParams {
	return &DeleteEngineRunProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEngineRunProfileParamsWithTimeout creates a new DeleteEngineRunProfileParams object
// with the ability to set a timeout on a request.
func NewDeleteEngineRunProfileParamsWithTimeout(timeout time.Duration) *DeleteEngineRunProfileParams {
	return &DeleteEngineRunProfileParams{
		timeout: timeout,
	}
}

// NewDeleteEngineRunProfileParamsWithContext creates a new DeleteEngineRunProfileParams object
// with the ability to set a context for a request.
func NewDeleteEngineRunProfileParamsWithContext(ctx context.Context) *DeleteEngineRunProfileParams {
	return &DeleteEngineRunProfileParams{
		Context: ctx,
	}
}

// NewDeleteEngineRunProfileParamsWithHTTPClient creates a new DeleteEngineRunProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEngineRunProfileParamsWithHTTPClient(client *http.Client) *DeleteEngineRunProfileParams {
	return &DeleteEngineRunProfileParams{
		HTTPClient: client,
	}
}

/*
DeleteEngineRunProfileParams contains all the parameters to send to the API endpoint

	for the delete engine run profile operation.

	Typically these are written to a http.Request.
*/
type DeleteEngineRunProfileParams struct {

	/* EngineID.

	   remote engine id
	*/
	EngineID string

	/* NewRunProfileID.

	   new run profile id
	*/
	NewRunProfileID *string

	/* RunProfileID.

	   run profile id to delete
	*/
	RunProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete engine run profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEngineRunProfileParams) WithDefaults() *DeleteEngineRunProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete engine run profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEngineRunProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) WithTimeout(timeout time.Duration) *DeleteEngineRunProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) WithContext(ctx context.Context) *DeleteEngineRunProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) WithHTTPClient(client *http.Client) *DeleteEngineRunProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEngineID adds the engineID to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) WithEngineID(engineID string) *DeleteEngineRunProfileParams {
	o.SetEngineID(engineID)
	return o
}

// SetEngineID adds the engineId to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) SetEngineID(engineID string) {
	o.EngineID = engineID
}

// WithNewRunProfileID adds the newRunProfileID to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) WithNewRunProfileID(newRunProfileID *string) *DeleteEngineRunProfileParams {
	o.SetNewRunProfileID(newRunProfileID)
	return o
}

// SetNewRunProfileID adds the newRunProfileId to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) SetNewRunProfileID(newRunProfileID *string) {
	o.NewRunProfileID = newRunProfileID
}

// WithRunProfileID adds the runProfileID to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) WithRunProfileID(runProfileID string) *DeleteEngineRunProfileParams {
	o.SetRunProfileID(runProfileID)
	return o
}

// SetRunProfileID adds the runProfileId to the delete engine run profile params
func (o *DeleteEngineRunProfileParams) SetRunProfileID(runProfileID string) {
	o.RunProfileID = runProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEngineRunProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param engineId
	if err := r.SetPathParam("engineId", o.EngineID); err != nil {
		return err
	}

	if o.NewRunProfileID != nil {

		// query param newRunProfileId
		var qrNewRunProfileID string

		if o.NewRunProfileID != nil {
			qrNewRunProfileID = *o.NewRunProfileID
		}
		qNewRunProfileID := qrNewRunProfileID
		if qNewRunProfileID != "" {

			if err := r.SetQueryParam("newRunProfileId", qNewRunProfileID); err != nil {
				return err
			}
		}
	}

	// path param runProfileId
	if err := r.SetPathParam("runProfileId", o.RunProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
