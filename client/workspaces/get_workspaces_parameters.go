// Code generated by go-swagger; DO NOT EDIT.

package workspaces

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

// NewGetWorkspacesParams creates a new GetWorkspacesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkspacesParams() *GetWorkspacesParams {
	return &GetWorkspacesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkspacesParamsWithTimeout creates a new GetWorkspacesParams object
// with the ability to set a timeout on a request.
func NewGetWorkspacesParamsWithTimeout(timeout time.Duration) *GetWorkspacesParams {
	return &GetWorkspacesParams{
		timeout: timeout,
	}
}

// NewGetWorkspacesParamsWithContext creates a new GetWorkspacesParams object
// with the ability to set a context for a request.
func NewGetWorkspacesParamsWithContext(ctx context.Context) *GetWorkspacesParams {
	return &GetWorkspacesParams{
		Context: ctx,
	}
}

// NewGetWorkspacesParamsWithHTTPClient creates a new GetWorkspacesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkspacesParamsWithHTTPClient(client *http.Client) *GetWorkspacesParams {
	return &GetWorkspacesParams{
		HTTPClient: client,
	}
}

/*
GetWorkspacesParams contains all the parameters to send to the API endpoint

	for the get workspaces operation.

	Typically these are written to a http.Request.
*/
type GetWorkspacesParams struct {

	/* Query.

	   search query (FIQL format), e.g. name==TestWorkspace,environment.name==TestEnvironment
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workspaces params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkspacesParams) WithDefaults() *GetWorkspacesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workspaces params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkspacesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workspaces params
func (o *GetWorkspacesParams) WithTimeout(timeout time.Duration) *GetWorkspacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workspaces params
func (o *GetWorkspacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workspaces params
func (o *GetWorkspacesParams) WithContext(ctx context.Context) *GetWorkspacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workspaces params
func (o *GetWorkspacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workspaces params
func (o *GetWorkspacesParams) WithHTTPClient(client *http.Client) *GetWorkspacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workspaces params
func (o *GetWorkspacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the get workspaces params
func (o *GetWorkspacesParams) WithQuery(query *string) *GetWorkspacesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get workspaces params
func (o *GetWorkspacesParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkspacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}