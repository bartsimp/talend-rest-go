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

// NewRemoveRemoteEngineFromClusterParams creates a new RemoveRemoteEngineFromClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveRemoteEngineFromClusterParams() *RemoveRemoteEngineFromClusterParams {
	return &RemoveRemoteEngineFromClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveRemoteEngineFromClusterParamsWithTimeout creates a new RemoveRemoteEngineFromClusterParams object
// with the ability to set a timeout on a request.
func NewRemoveRemoteEngineFromClusterParamsWithTimeout(timeout time.Duration) *RemoveRemoteEngineFromClusterParams {
	return &RemoveRemoteEngineFromClusterParams{
		timeout: timeout,
	}
}

// NewRemoveRemoteEngineFromClusterParamsWithContext creates a new RemoveRemoteEngineFromClusterParams object
// with the ability to set a context for a request.
func NewRemoveRemoteEngineFromClusterParamsWithContext(ctx context.Context) *RemoveRemoteEngineFromClusterParams {
	return &RemoveRemoteEngineFromClusterParams{
		Context: ctx,
	}
}

// NewRemoveRemoteEngineFromClusterParamsWithHTTPClient creates a new RemoveRemoteEngineFromClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveRemoteEngineFromClusterParamsWithHTTPClient(client *http.Client) *RemoveRemoteEngineFromClusterParams {
	return &RemoveRemoteEngineFromClusterParams{
		HTTPClient: client,
	}
}

/*
RemoveRemoteEngineFromClusterParams contains all the parameters to send to the API endpoint

	for the remove remote engine from cluster operation.

	Typically these are written to a http.Request.
*/
type RemoveRemoteEngineFromClusterParams struct {

	/* ClusterID.

	   remote engine cluster id
	*/
	ClusterID string

	/* EngineID.

	   remote engine id
	*/
	EngineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove remote engine from cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveRemoteEngineFromClusterParams) WithDefaults() *RemoveRemoteEngineFromClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove remote engine from cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveRemoteEngineFromClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) WithTimeout(timeout time.Duration) *RemoveRemoteEngineFromClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) WithContext(ctx context.Context) *RemoveRemoteEngineFromClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) WithHTTPClient(client *http.Client) *RemoveRemoteEngineFromClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) WithClusterID(clusterID string) *RemoveRemoteEngineFromClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithEngineID adds the engineID to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) WithEngineID(engineID string) *RemoveRemoteEngineFromClusterParams {
	o.SetEngineID(engineID)
	return o
}

// SetEngineID adds the engineId to the remove remote engine from cluster params
func (o *RemoveRemoteEngineFromClusterParams) SetEngineID(engineID string) {
	o.EngineID = engineID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveRemoteEngineFromClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	// path param engineId
	if err := r.SetPathParam("engineId", o.EngineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}