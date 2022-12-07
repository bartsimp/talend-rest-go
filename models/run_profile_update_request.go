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

// RunProfileUpdateRequest run profile update request
//
// swagger:model RunProfileUpdateRequest
type RunProfileUpdateRequest struct {

	// Description
	// Example: Run profile with default options
	Description string `json:"description,omitempty"`

	// Jvm arguments
	// Example: ["-Xms128m","-Xmx256m"]
	// Required: true
	JvmArguments []string `json:"jvmArguments"`

	// Run profile name
	// Example: Run profile
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this run profile update request
func (m *RunProfileUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJvmArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunProfileUpdateRequest) validateJvmArguments(formats strfmt.Registry) error {

	if err := validate.Required("jvmArguments", "body", m.JvmArguments); err != nil {
		return err
	}

	return nil
}

func (m *RunProfileUpdateRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this run profile update request based on context it is used
func (m *RunProfileUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunProfileUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunProfileUpdateRequest) UnmarshalBinary(b []byte) error {
	var res RunProfileUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
