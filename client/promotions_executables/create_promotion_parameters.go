// Code generated by go-swagger; DO NOT EDIT.

package promotions_executables

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

// NewCreatePromotionParams creates a new CreatePromotionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePromotionParams() *CreatePromotionParams {
	return &CreatePromotionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePromotionParamsWithTimeout creates a new CreatePromotionParams object
// with the ability to set a timeout on a request.
func NewCreatePromotionParamsWithTimeout(timeout time.Duration) *CreatePromotionParams {
	return &CreatePromotionParams{
		timeout: timeout,
	}
}

// NewCreatePromotionParamsWithContext creates a new CreatePromotionParams object
// with the ability to set a context for a request.
func NewCreatePromotionParamsWithContext(ctx context.Context) *CreatePromotionParams {
	return &CreatePromotionParams{
		Context: ctx,
	}
}

// NewCreatePromotionParamsWithHTTPClient creates a new CreatePromotionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePromotionParamsWithHTTPClient(client *http.Client) *CreatePromotionParams {
	return &CreatePromotionParams{
		HTTPClient: client,
	}
}

/*
CreatePromotionParams contains all the parameters to send to the API endpoint

	for the create promotion operation.

	Typically these are written to a http.Request.
*/
type CreatePromotionParams struct {

	/* Body.

	   Promotion
	*/
	Body *models.CreatePromotionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create promotion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePromotionParams) WithDefaults() *CreatePromotionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create promotion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePromotionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create promotion params
func (o *CreatePromotionParams) WithTimeout(timeout time.Duration) *CreatePromotionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create promotion params
func (o *CreatePromotionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create promotion params
func (o *CreatePromotionParams) WithContext(ctx context.Context) *CreatePromotionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create promotion params
func (o *CreatePromotionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create promotion params
func (o *CreatePromotionParams) WithHTTPClient(client *http.Client) *CreatePromotionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create promotion params
func (o *CreatePromotionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create promotion params
func (o *CreatePromotionParams) WithBody(body *models.CreatePromotionRequest) *CreatePromotionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create promotion params
func (o *CreatePromotionParams) SetBody(body *models.CreatePromotionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePromotionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
