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

// Seat Seat information
//
// swagger:model Seat
type Seat struct {

	// Name of application
	// Example: ["Data Inventory","Data Preparation","Data Stewardship","Pipeline Designer","Dictionary service","API Designer / API Tester","Management Console","Studio","Service Accounts"]
	// Required: true
	Application *string `json:"application"`

	// Count of users currently using the application
	// Required: true
	CurrentUsers *int32 `json:"currentUsers"`

	// Maximum of users allowed to use the application
	// Required: true
	MaxUsers *int32 `json:"maxUsers"`
}

// Validate validates this seat
func (m *Seat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Seat) validateApplication(formats strfmt.Registry) error {

	if err := validate.Required("application", "body", m.Application); err != nil {
		return err
	}

	return nil
}

func (m *Seat) validateCurrentUsers(formats strfmt.Registry) error {

	if err := validate.Required("currentUsers", "body", m.CurrentUsers); err != nil {
		return err
	}

	return nil
}

func (m *Seat) validateMaxUsers(formats strfmt.Registry) error {

	if err := validate.Required("maxUsers", "body", m.MaxUsers); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this seat based on context it is used
func (m *Seat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Seat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Seat) UnmarshalBinary(b []byte) error {
	var res Seat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}