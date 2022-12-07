// Code generated by go-swagger; DO NOT EDIT.

package promotions_executables_authorization

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

// NewDeletePromotionUserParams creates a new DeletePromotionUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePromotionUserParams() *DeletePromotionUserParams {
	return &DeletePromotionUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePromotionUserParamsWithTimeout creates a new DeletePromotionUserParams object
// with the ability to set a timeout on a request.
func NewDeletePromotionUserParamsWithTimeout(timeout time.Duration) *DeletePromotionUserParams {
	return &DeletePromotionUserParams{
		timeout: timeout,
	}
}

// NewDeletePromotionUserParamsWithContext creates a new DeletePromotionUserParams object
// with the ability to set a context for a request.
func NewDeletePromotionUserParamsWithContext(ctx context.Context) *DeletePromotionUserParams {
	return &DeletePromotionUserParams{
		Context: ctx,
	}
}

// NewDeletePromotionUserParamsWithHTTPClient creates a new DeletePromotionUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePromotionUserParamsWithHTTPClient(client *http.Client) *DeletePromotionUserParams {
	return &DeletePromotionUserParams{
		HTTPClient: client,
	}
}

/*
DeletePromotionUserParams contains all the parameters to send to the API endpoint

	for the delete promotion user operation.

	Typically these are written to a http.Request.
*/
type DeletePromotionUserParams struct {

	/* PromotionID.

	   Promotion ID
	*/
	PromotionID string

	/* UserID.

	   User ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete promotion user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePromotionUserParams) WithDefaults() *DeletePromotionUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete promotion user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePromotionUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete promotion user params
func (o *DeletePromotionUserParams) WithTimeout(timeout time.Duration) *DeletePromotionUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete promotion user params
func (o *DeletePromotionUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete promotion user params
func (o *DeletePromotionUserParams) WithContext(ctx context.Context) *DeletePromotionUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete promotion user params
func (o *DeletePromotionUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete promotion user params
func (o *DeletePromotionUserParams) WithHTTPClient(client *http.Client) *DeletePromotionUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete promotion user params
func (o *DeletePromotionUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPromotionID adds the promotionID to the delete promotion user params
func (o *DeletePromotionUserParams) WithPromotionID(promotionID string) *DeletePromotionUserParams {
	o.SetPromotionID(promotionID)
	return o
}

// SetPromotionID adds the promotionId to the delete promotion user params
func (o *DeletePromotionUserParams) SetPromotionID(promotionID string) {
	o.PromotionID = promotionID
}

// WithUserID adds the userID to the delete promotion user params
func (o *DeletePromotionUserParams) WithUserID(userID string) *DeletePromotionUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete promotion user params
func (o *DeletePromotionUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePromotionUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param promotionId
	if err := r.SetPathParam("promotionId", o.PromotionID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}