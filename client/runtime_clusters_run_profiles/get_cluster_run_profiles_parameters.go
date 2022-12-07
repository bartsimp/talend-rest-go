// Code generated by go-swagger; DO NOT EDIT.

package runtime_clusters_run_profiles

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

// NewGetClusterRunProfilesParams creates a new GetClusterRunProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterRunProfilesParams() *GetClusterRunProfilesParams {
	return &GetClusterRunProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterRunProfilesParamsWithTimeout creates a new GetClusterRunProfilesParams object
// with the ability to set a timeout on a request.
func NewGetClusterRunProfilesParamsWithTimeout(timeout time.Duration) *GetClusterRunProfilesParams {
	return &GetClusterRunProfilesParams{
		timeout: timeout,
	}
}

// NewGetClusterRunProfilesParamsWithContext creates a new GetClusterRunProfilesParams object
// with the ability to set a context for a request.
func NewGetClusterRunProfilesParamsWithContext(ctx context.Context) *GetClusterRunProfilesParams {
	return &GetClusterRunProfilesParams{
		Context: ctx,
	}
}

// NewGetClusterRunProfilesParamsWithHTTPClient creates a new GetClusterRunProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterRunProfilesParamsWithHTTPClient(client *http.Client) *GetClusterRunProfilesParams {
	return &GetClusterRunProfilesParams{
		HTTPClient: client,
	}
}

/*
GetClusterRunProfilesParams contains all the parameters to send to the API endpoint

	for the get cluster run profiles operation.

	Typically these are written to a http.Request.
*/
type GetClusterRunProfilesParams struct {

	/* ClusterID.

	   cluster id
	*/
	ClusterID string

	/* Type.

	   Type of run profiles to be returned. By default returns all
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster run profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterRunProfilesParams) WithDefaults() *GetClusterRunProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster run profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterRunProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) WithTimeout(timeout time.Duration) *GetClusterRunProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) WithContext(ctx context.Context) *GetClusterRunProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) WithHTTPClient(client *http.Client) *GetClusterRunProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) WithClusterID(clusterID string) *GetClusterRunProfilesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithType adds the typeVar to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) WithType(typeVar *string) *GetClusterRunProfilesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get cluster run profiles params
func (o *GetClusterRunProfilesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterRunProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
