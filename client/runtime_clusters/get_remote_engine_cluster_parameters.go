// Code generated by go-swagger; DO NOT EDIT.

package runtime_clusters

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

// NewGetRemoteEngineClusterParams creates a new GetRemoteEngineClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRemoteEngineClusterParams() *GetRemoteEngineClusterParams {
	return &GetRemoteEngineClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRemoteEngineClusterParamsWithTimeout creates a new GetRemoteEngineClusterParams object
// with the ability to set a timeout on a request.
func NewGetRemoteEngineClusterParamsWithTimeout(timeout time.Duration) *GetRemoteEngineClusterParams {
	return &GetRemoteEngineClusterParams{
		timeout: timeout,
	}
}

// NewGetRemoteEngineClusterParamsWithContext creates a new GetRemoteEngineClusterParams object
// with the ability to set a context for a request.
func NewGetRemoteEngineClusterParamsWithContext(ctx context.Context) *GetRemoteEngineClusterParams {
	return &GetRemoteEngineClusterParams{
		Context: ctx,
	}
}

// NewGetRemoteEngineClusterParamsWithHTTPClient creates a new GetRemoteEngineClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRemoteEngineClusterParamsWithHTTPClient(client *http.Client) *GetRemoteEngineClusterParams {
	return &GetRemoteEngineClusterParams{
		HTTPClient: client,
	}
}

/*
GetRemoteEngineClusterParams contains all the parameters to send to the API endpoint

	for the get remote engine cluster operation.

	Typically these are written to a http.Request.
*/
type GetRemoteEngineClusterParams struct {

	/* ClusterID.

	   remote engine cluster id
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get remote engine cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRemoteEngineClusterParams) WithDefaults() *GetRemoteEngineClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get remote engine cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRemoteEngineClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get remote engine cluster params
func (o *GetRemoteEngineClusterParams) WithTimeout(timeout time.Duration) *GetRemoteEngineClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get remote engine cluster params
func (o *GetRemoteEngineClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get remote engine cluster params
func (o *GetRemoteEngineClusterParams) WithContext(ctx context.Context) *GetRemoteEngineClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get remote engine cluster params
func (o *GetRemoteEngineClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get remote engine cluster params
func (o *GetRemoteEngineClusterParams) WithHTTPClient(client *http.Client) *GetRemoteEngineClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get remote engine cluster params
func (o *GetRemoteEngineClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get remote engine cluster params
func (o *GetRemoteEngineClusterParams) WithClusterID(clusterID string) *GetRemoteEngineClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get remote engine cluster params
func (o *GetRemoteEngineClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRemoteEngineClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}