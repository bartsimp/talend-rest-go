// Code generated by go-swagger; DO NOT EDIT.

package workspaces_permissions

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

// NewGetPermissionsParams creates a new GetPermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPermissionsParams() *GetPermissionsParams {
	return &GetPermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPermissionsParamsWithTimeout creates a new GetPermissionsParams object
// with the ability to set a timeout on a request.
func NewGetPermissionsParamsWithTimeout(timeout time.Duration) *GetPermissionsParams {
	return &GetPermissionsParams{
		timeout: timeout,
	}
}

// NewGetPermissionsParamsWithContext creates a new GetPermissionsParams object
// with the ability to set a context for a request.
func NewGetPermissionsParamsWithContext(ctx context.Context) *GetPermissionsParams {
	return &GetPermissionsParams{
		Context: ctx,
	}
}

// NewGetPermissionsParamsWithHTTPClient creates a new GetPermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPermissionsParamsWithHTTPClient(client *http.Client) *GetPermissionsParams {
	return &GetPermissionsParams{
		HTTPClient: client,
	}
}

/*
GetPermissionsParams contains all the parameters to send to the API endpoint

	for the get permissions operation.

	Typically these are written to a http.Request.
*/
type GetPermissionsParams struct {

	/* EnvironmentID.

	     The environment id (Optional)
	Example Value: 6089228181ef4423736e47a8
	*/
	EnvironmentID *string

	/* UserID.

	     The user id (Optional)
	Example Value: b9e10a3f-9d68-44bb-862f-b2aa56dc7191
	*/
	UserID *string

	/* WorkspaceID.

	     The workspace id (Optional)
	Example Value: 6089228181ef4423736e47a9
	*/
	WorkspaceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPermissionsParams) WithDefaults() *GetPermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get permissions params
func (o *GetPermissionsParams) WithTimeout(timeout time.Duration) *GetPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get permissions params
func (o *GetPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get permissions params
func (o *GetPermissionsParams) WithContext(ctx context.Context) *GetPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get permissions params
func (o *GetPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get permissions params
func (o *GetPermissionsParams) WithHTTPClient(client *http.Client) *GetPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get permissions params
func (o *GetPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the get permissions params
func (o *GetPermissionsParams) WithEnvironmentID(environmentID *string) *GetPermissionsParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the get permissions params
func (o *GetPermissionsParams) SetEnvironmentID(environmentID *string) {
	o.EnvironmentID = environmentID
}

// WithUserID adds the userID to the get permissions params
func (o *GetPermissionsParams) WithUserID(userID *string) *GetPermissionsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get permissions params
func (o *GetPermissionsParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WithWorkspaceID adds the workspaceID to the get permissions params
func (o *GetPermissionsParams) WithWorkspaceID(workspaceID *string) *GetPermissionsParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get permissions params
func (o *GetPermissionsParams) SetWorkspaceID(workspaceID *string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvironmentID != nil {

		// query param environmentId
		var qrEnvironmentID string

		if o.EnvironmentID != nil {
			qrEnvironmentID = *o.EnvironmentID
		}
		qEnvironmentID := qrEnvironmentID
		if qEnvironmentID != "" {

			if err := r.SetQueryParam("environmentId", qEnvironmentID); err != nil {
				return err
			}
		}
	}

	if o.UserID != nil {

		// query param userId
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("userId", qUserID); err != nil {
				return err
			}
		}
	}

	if o.WorkspaceID != nil {

		// query param workspaceId
		var qrWorkspaceID string

		if o.WorkspaceID != nil {
			qrWorkspaceID = *o.WorkspaceID
		}
		qWorkspaceID := qrWorkspaceID
		if qWorkspaceID != "" {

			if err := r.SetQueryParam("workspaceId", qWorkspaceID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
