// Code generated by go-swagger; DO NOT EDIT.

package runtime_engines

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

// NewDeleteRemoteEngineParams creates a new DeleteRemoteEngineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRemoteEngineParams() *DeleteRemoteEngineParams {
	return &DeleteRemoteEngineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRemoteEngineParamsWithTimeout creates a new DeleteRemoteEngineParams object
// with the ability to set a timeout on a request.
func NewDeleteRemoteEngineParamsWithTimeout(timeout time.Duration) *DeleteRemoteEngineParams {
	return &DeleteRemoteEngineParams{
		timeout: timeout,
	}
}

// NewDeleteRemoteEngineParamsWithContext creates a new DeleteRemoteEngineParams object
// with the ability to set a context for a request.
func NewDeleteRemoteEngineParamsWithContext(ctx context.Context) *DeleteRemoteEngineParams {
	return &DeleteRemoteEngineParams{
		Context: ctx,
	}
}

// NewDeleteRemoteEngineParamsWithHTTPClient creates a new DeleteRemoteEngineParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRemoteEngineParamsWithHTTPClient(client *http.Client) *DeleteRemoteEngineParams {
	return &DeleteRemoteEngineParams{
		HTTPClient: client,
	}
}

/*
DeleteRemoteEngineParams contains all the parameters to send to the API endpoint

	for the delete remote engine operation.

	Typically these are written to a http.Request.
*/
type DeleteRemoteEngineParams struct {

	/* ID.

	   remote engine id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete remote engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRemoteEngineParams) WithDefaults() *DeleteRemoteEngineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete remote engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRemoteEngineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete remote engine params
func (o *DeleteRemoteEngineParams) WithTimeout(timeout time.Duration) *DeleteRemoteEngineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete remote engine params
func (o *DeleteRemoteEngineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete remote engine params
func (o *DeleteRemoteEngineParams) WithContext(ctx context.Context) *DeleteRemoteEngineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete remote engine params
func (o *DeleteRemoteEngineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete remote engine params
func (o *DeleteRemoteEngineParams) WithHTTPClient(client *http.Client) *DeleteRemoteEngineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete remote engine params
func (o *DeleteRemoteEngineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete remote engine params
func (o *DeleteRemoteEngineParams) WithID(id string) *DeleteRemoteEngineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete remote engine params
func (o *DeleteRemoteEngineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRemoteEngineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}