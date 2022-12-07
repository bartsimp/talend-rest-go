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

	"github.com/bartsimp/talend-rest-go/models"
)

// NewUpdateCustomWorkspaceParams creates a new UpdateCustomWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCustomWorkspaceParams() *UpdateCustomWorkspaceParams {
	return &UpdateCustomWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCustomWorkspaceParamsWithTimeout creates a new UpdateCustomWorkspaceParams object
// with the ability to set a timeout on a request.
func NewUpdateCustomWorkspaceParamsWithTimeout(timeout time.Duration) *UpdateCustomWorkspaceParams {
	return &UpdateCustomWorkspaceParams{
		timeout: timeout,
	}
}

// NewUpdateCustomWorkspaceParamsWithContext creates a new UpdateCustomWorkspaceParams object
// with the ability to set a context for a request.
func NewUpdateCustomWorkspaceParamsWithContext(ctx context.Context) *UpdateCustomWorkspaceParams {
	return &UpdateCustomWorkspaceParams{
		Context: ctx,
	}
}

// NewUpdateCustomWorkspaceParamsWithHTTPClient creates a new UpdateCustomWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCustomWorkspaceParamsWithHTTPClient(client *http.Client) *UpdateCustomWorkspaceParams {
	return &UpdateCustomWorkspaceParams{
		HTTPClient: client,
	}
}

/*
UpdateCustomWorkspaceParams contains all the parameters to send to the API endpoint

	for the update custom workspace operation.

	Typically these are written to a http.Request.
*/
type UpdateCustomWorkspaceParams struct {

	/* Body.

	   Workspace to update
	*/
	Body *models.UpdateWorkspaceRequest

	// WorkspaceID.
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update custom workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCustomWorkspaceParams) WithDefaults() *UpdateCustomWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update custom workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCustomWorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) WithTimeout(timeout time.Duration) *UpdateCustomWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) WithContext(ctx context.Context) *UpdateCustomWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) WithHTTPClient(client *http.Client) *UpdateCustomWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) WithBody(body *models.UpdateWorkspaceRequest) *UpdateCustomWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) SetBody(body *models.UpdateWorkspaceRequest) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) WithWorkspaceID(workspaceID string) *UpdateCustomWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the update custom workspace params
func (o *UpdateCustomWorkspaceParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCustomWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
