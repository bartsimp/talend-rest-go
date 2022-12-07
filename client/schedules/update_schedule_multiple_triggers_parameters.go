// Code generated by go-swagger; DO NOT EDIT.

package schedules

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

	"github.com/bartsimp/talend-rest-go/models"
)

// NewUpdateScheduleMultipleTriggersParams creates a new UpdateScheduleMultipleTriggersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateScheduleMultipleTriggersParams() *UpdateScheduleMultipleTriggersParams {
	return &UpdateScheduleMultipleTriggersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScheduleMultipleTriggersParamsWithTimeout creates a new UpdateScheduleMultipleTriggersParams object
// with the ability to set a timeout on a request.
func NewUpdateScheduleMultipleTriggersParamsWithTimeout(timeout time.Duration) *UpdateScheduleMultipleTriggersParams {
	return &UpdateScheduleMultipleTriggersParams{
		timeout: timeout,
	}
}

// NewUpdateScheduleMultipleTriggersParamsWithContext creates a new UpdateScheduleMultipleTriggersParams object
// with the ability to set a context for a request.
func NewUpdateScheduleMultipleTriggersParamsWithContext(ctx context.Context) *UpdateScheduleMultipleTriggersParams {
	return &UpdateScheduleMultipleTriggersParams{
		Context: ctx,
	}
}

// NewUpdateScheduleMultipleTriggersParamsWithHTTPClient creates a new UpdateScheduleMultipleTriggersParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateScheduleMultipleTriggersParamsWithHTTPClient(client *http.Client) *UpdateScheduleMultipleTriggersParams {
	return &UpdateScheduleMultipleTriggersParams{
		HTTPClient: client,
	}
}

/*
UpdateScheduleMultipleTriggersParams contains all the parameters to send to the API endpoint

	for the update schedule multiple triggers operation.

	Typically these are written to a http.Request.
*/
type UpdateScheduleMultipleTriggersParams struct {

	/* Body.

	   Description update
	*/
	Body *models.UpdateScheduleMultipleTriggerRequest

	/* ScheduleID.

	   schedules id
	*/
	ScheduleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update schedule multiple triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScheduleMultipleTriggersParams) WithDefaults() *UpdateScheduleMultipleTriggersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update schedule multiple triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScheduleMultipleTriggersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) WithTimeout(timeout time.Duration) *UpdateScheduleMultipleTriggersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) WithContext(ctx context.Context) *UpdateScheduleMultipleTriggersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) WithHTTPClient(client *http.Client) *UpdateScheduleMultipleTriggersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) WithBody(body *models.UpdateScheduleMultipleTriggerRequest) *UpdateScheduleMultipleTriggersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) SetBody(body *models.UpdateScheduleMultipleTriggerRequest) {
	o.Body = body
}

// WithScheduleID adds the scheduleID to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) WithScheduleID(scheduleID string) *UpdateScheduleMultipleTriggersParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the update schedule multiple triggers params
func (o *UpdateScheduleMultipleTriggersParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScheduleMultipleTriggersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param scheduleId
	if err := r.SetPathParam("scheduleId", o.ScheduleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}