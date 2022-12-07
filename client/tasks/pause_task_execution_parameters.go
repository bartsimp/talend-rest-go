// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewPauseTaskExecutionParams creates a new PauseTaskExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPauseTaskExecutionParams() *PauseTaskExecutionParams {
	return &PauseTaskExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPauseTaskExecutionParamsWithTimeout creates a new PauseTaskExecutionParams object
// with the ability to set a timeout on a request.
func NewPauseTaskExecutionParamsWithTimeout(timeout time.Duration) *PauseTaskExecutionParams {
	return &PauseTaskExecutionParams{
		timeout: timeout,
	}
}

// NewPauseTaskExecutionParamsWithContext creates a new PauseTaskExecutionParams object
// with the ability to set a context for a request.
func NewPauseTaskExecutionParamsWithContext(ctx context.Context) *PauseTaskExecutionParams {
	return &PauseTaskExecutionParams{
		Context: ctx,
	}
}

// NewPauseTaskExecutionParamsWithHTTPClient creates a new PauseTaskExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPauseTaskExecutionParamsWithHTTPClient(client *http.Client) *PauseTaskExecutionParams {
	return &PauseTaskExecutionParams{
		HTTPClient: client,
	}
}

/*
PauseTaskExecutionParams contains all the parameters to send to the API endpoint

	for the pause task execution operation.

	Typically these are written to a http.Request.
*/
type PauseTaskExecutionParams struct {

	/* Body.

	   Action to perform
	*/
	Body *models.TaskPauseDetails

	/* TaskID.

	   task id
	*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pause task execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseTaskExecutionParams) WithDefaults() *PauseTaskExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pause task execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseTaskExecutionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pause task execution params
func (o *PauseTaskExecutionParams) WithTimeout(timeout time.Duration) *PauseTaskExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pause task execution params
func (o *PauseTaskExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pause task execution params
func (o *PauseTaskExecutionParams) WithContext(ctx context.Context) *PauseTaskExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pause task execution params
func (o *PauseTaskExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pause task execution params
func (o *PauseTaskExecutionParams) WithHTTPClient(client *http.Client) *PauseTaskExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pause task execution params
func (o *PauseTaskExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pause task execution params
func (o *PauseTaskExecutionParams) WithBody(body *models.TaskPauseDetails) *PauseTaskExecutionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pause task execution params
func (o *PauseTaskExecutionParams) SetBody(body *models.TaskPauseDetails) {
	o.Body = body
}

// WithTaskID adds the taskID to the pause task execution params
func (o *PauseTaskExecutionParams) WithTaskID(taskID string) *PauseTaskExecutionParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the pause task execution params
func (o *PauseTaskExecutionParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *PauseTaskExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param taskId
	if err := r.SetPathParam("taskId", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
