// Code generated by go-swagger; DO NOT EDIT.

package plans_executables

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

// NewUnlinkPlanScheduleParams creates a new UnlinkPlanScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnlinkPlanScheduleParams() *UnlinkPlanScheduleParams {
	return &UnlinkPlanScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnlinkPlanScheduleParamsWithTimeout creates a new UnlinkPlanScheduleParams object
// with the ability to set a timeout on a request.
func NewUnlinkPlanScheduleParamsWithTimeout(timeout time.Duration) *UnlinkPlanScheduleParams {
	return &UnlinkPlanScheduleParams{
		timeout: timeout,
	}
}

// NewUnlinkPlanScheduleParamsWithContext creates a new UnlinkPlanScheduleParams object
// with the ability to set a context for a request.
func NewUnlinkPlanScheduleParamsWithContext(ctx context.Context) *UnlinkPlanScheduleParams {
	return &UnlinkPlanScheduleParams{
		Context: ctx,
	}
}

// NewUnlinkPlanScheduleParamsWithHTTPClient creates a new UnlinkPlanScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnlinkPlanScheduleParamsWithHTTPClient(client *http.Client) *UnlinkPlanScheduleParams {
	return &UnlinkPlanScheduleParams{
		HTTPClient: client,
	}
}

/*
UnlinkPlanScheduleParams contains all the parameters to send to the API endpoint

	for the unlink plan schedule operation.

	Typically these are written to a http.Request.
*/
type UnlinkPlanScheduleParams struct {

	/* PlanID.

	   plan id
	*/
	PlanID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unlink plan schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnlinkPlanScheduleParams) WithDefaults() *UnlinkPlanScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unlink plan schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnlinkPlanScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unlink plan schedule params
func (o *UnlinkPlanScheduleParams) WithTimeout(timeout time.Duration) *UnlinkPlanScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unlink plan schedule params
func (o *UnlinkPlanScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unlink plan schedule params
func (o *UnlinkPlanScheduleParams) WithContext(ctx context.Context) *UnlinkPlanScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unlink plan schedule params
func (o *UnlinkPlanScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unlink plan schedule params
func (o *UnlinkPlanScheduleParams) WithHTTPClient(client *http.Client) *UnlinkPlanScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unlink plan schedule params
func (o *UnlinkPlanScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanID adds the planID to the unlink plan schedule params
func (o *UnlinkPlanScheduleParams) WithPlanID(planID string) *UnlinkPlanScheduleParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the unlink plan schedule params
func (o *UnlinkPlanScheduleParams) SetPlanID(planID string) {
	o.PlanID = planID
}

// WriteToRequest writes these params to a swagger request
func (o *UnlinkPlanScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param planId
	if err := r.SetPathParam("planId", o.PlanID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}