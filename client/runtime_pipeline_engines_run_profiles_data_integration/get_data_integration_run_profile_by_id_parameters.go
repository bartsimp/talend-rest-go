// Code generated by go-swagger; DO NOT EDIT.

package runtime_pipeline_engines_run_profiles_data_integration

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

// NewGetDataIntegrationRunProfileByIDParams creates a new GetDataIntegrationRunProfileByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDataIntegrationRunProfileByIDParams() *GetDataIntegrationRunProfileByIDParams {
	return &GetDataIntegrationRunProfileByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataIntegrationRunProfileByIDParamsWithTimeout creates a new GetDataIntegrationRunProfileByIDParams object
// with the ability to set a timeout on a request.
func NewGetDataIntegrationRunProfileByIDParamsWithTimeout(timeout time.Duration) *GetDataIntegrationRunProfileByIDParams {
	return &GetDataIntegrationRunProfileByIDParams{
		timeout: timeout,
	}
}

// NewGetDataIntegrationRunProfileByIDParamsWithContext creates a new GetDataIntegrationRunProfileByIDParams object
// with the ability to set a context for a request.
func NewGetDataIntegrationRunProfileByIDParamsWithContext(ctx context.Context) *GetDataIntegrationRunProfileByIDParams {
	return &GetDataIntegrationRunProfileByIDParams{
		Context: ctx,
	}
}

// NewGetDataIntegrationRunProfileByIDParamsWithHTTPClient creates a new GetDataIntegrationRunProfileByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDataIntegrationRunProfileByIDParamsWithHTTPClient(client *http.Client) *GetDataIntegrationRunProfileByIDParams {
	return &GetDataIntegrationRunProfileByIDParams{
		HTTPClient: client,
	}
}

/*
GetDataIntegrationRunProfileByIDParams contains all the parameters to send to the API endpoint

	for the get data integration run profile by Id operation.

	Typically these are written to a http.Request.
*/
type GetDataIntegrationRunProfileByIDParams struct {

	/* EngineID.

	   engine id
	*/
	EngineID string

	/* RunProfileID.

	   run profile id
	*/
	RunProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get data integration run profile by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataIntegrationRunProfileByIDParams) WithDefaults() *GetDataIntegrationRunProfileByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get data integration run profile by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataIntegrationRunProfileByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) WithTimeout(timeout time.Duration) *GetDataIntegrationRunProfileByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) WithContext(ctx context.Context) *GetDataIntegrationRunProfileByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) WithHTTPClient(client *http.Client) *GetDataIntegrationRunProfileByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEngineID adds the engineID to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) WithEngineID(engineID string) *GetDataIntegrationRunProfileByIDParams {
	o.SetEngineID(engineID)
	return o
}

// SetEngineID adds the engineId to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) SetEngineID(engineID string) {
	o.EngineID = engineID
}

// WithRunProfileID adds the runProfileID to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) WithRunProfileID(runProfileID string) *GetDataIntegrationRunProfileByIDParams {
	o.SetRunProfileID(runProfileID)
	return o
}

// SetRunProfileID adds the runProfileId to the get data integration run profile by Id params
func (o *GetDataIntegrationRunProfileByIDParams) SetRunProfileID(runProfileID string) {
	o.RunProfileID = runProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataIntegrationRunProfileByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param engineId
	if err := r.SetPathParam("engineId", o.EngineID); err != nil {
		return err
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
