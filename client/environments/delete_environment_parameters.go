// Code generated by go-swagger; DO NOT EDIT.

package environments

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

// NewDeleteEnvironmentParams creates a new DeleteEnvironmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEnvironmentParams() *DeleteEnvironmentParams {
	return &DeleteEnvironmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEnvironmentParamsWithTimeout creates a new DeleteEnvironmentParams object
// with the ability to set a timeout on a request.
func NewDeleteEnvironmentParamsWithTimeout(timeout time.Duration) *DeleteEnvironmentParams {
	return &DeleteEnvironmentParams{
		timeout: timeout,
	}
}

// NewDeleteEnvironmentParamsWithContext creates a new DeleteEnvironmentParams object
// with the ability to set a context for a request.
func NewDeleteEnvironmentParamsWithContext(ctx context.Context) *DeleteEnvironmentParams {
	return &DeleteEnvironmentParams{
		Context: ctx,
	}
}

// NewDeleteEnvironmentParamsWithHTTPClient creates a new DeleteEnvironmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEnvironmentParamsWithHTTPClient(client *http.Client) *DeleteEnvironmentParams {
	return &DeleteEnvironmentParams{
		HTTPClient: client,
	}
}

/*
DeleteEnvironmentParams contains all the parameters to send to the API endpoint

	for the delete environment operation.

	Typically these are written to a http.Request.
*/
type DeleteEnvironmentParams struct {

	// EnvironmentID.
	EnvironmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEnvironmentParams) WithDefaults() *DeleteEnvironmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEnvironmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete environment params
func (o *DeleteEnvironmentParams) WithTimeout(timeout time.Duration) *DeleteEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete environment params
func (o *DeleteEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete environment params
func (o *DeleteEnvironmentParams) WithContext(ctx context.Context) *DeleteEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete environment params
func (o *DeleteEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete environment params
func (o *DeleteEnvironmentParams) WithHTTPClient(client *http.Client) *DeleteEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete environment params
func (o *DeleteEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the delete environment params
func (o *DeleteEnvironmentParams) WithEnvironmentID(environmentID string) *DeleteEnvironmentParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the delete environment params
func (o *DeleteEnvironmentParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentId
	if err := r.SetPathParam("environmentId", o.EnvironmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
