// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ErrorResponse Error response object
//
// swagger:model ErrorResponse
type ErrorResponse struct {

	// Internal error code
	// Example: TIPASS-SCHED-123
	Code string `json:"code,omitempty"`

	// Developer message (not translated). Info about error for developer.
	// Example: more descriptive info about something went wrong
	Details string `json:"details,omitempty"`

	// Error message (multilanguage). Info about error for user.
	// Example: something went wrong
	// Required: true
	Message *string `json:"message"`

	// The unique identification of the request involved with this error
	// Example: 16fefb53-035a-4249-af9d-f80a3b47b132
	RequestID string `json:"requestId,omitempty"`

	// Status code
	// Example: 400, 401, 403, 404, 409, 500, 502, 503, etc.
	// Required: true
	Status *int32 `json:"status"`

	// URL provided detailed info about error
	// Example: https://error.talend.com/details/tic?code=SCHED-123
	URL string `json:"url,omitempty"`
}

// Validate validates this error response
func (m *ErrorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponse) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *ErrorResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error response based on context it is used
func (m *ErrorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponse) UnmarshalBinary(b []byte) error {
	var res ErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}