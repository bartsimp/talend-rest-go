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

// EngineStatus Information about engine status
//
// swagger:model EngineStatus
type EngineStatus struct {

	// Number of engines used by type
	// Required: true
	Consumed *int32 `json:"consumed"`

	// Engine type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this engine status
func (m *EngineStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsumed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineStatus) validateConsumed(formats strfmt.Registry) error {

	if err := validate.Required("consumed", "body", m.Consumed); err != nil {
		return err
	}

	return nil
}

func (m *EngineStatus) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this engine status based on context it is used
func (m *EngineStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineStatus) UnmarshalBinary(b []byte) error {
	var res EngineStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}