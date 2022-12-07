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

// CreatePromotionRequest create promotion request
//
// swagger:model CreatePromotionRequest
type CreatePromotionRequest struct {

	// Promotion description
	// Example: This is description
	Description string `json:"description,omitempty"`

	// Source environment ID
	// Example: 5b3ee8676e8f510dbb3fa806
	// Required: true
	SourceEnvironmentID *string `json:"sourceEnvironmentId"`

	// Target environment ID
	// Example: 5b3ee8676e8f510dbb3fa807
	// Required: true
	TargetEnvironmentID *string `json:"targetEnvironmentId"`
}

// Validate validates this create promotion request
func (m *CreatePromotionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceEnvironmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEnvironmentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePromotionRequest) validateSourceEnvironmentID(formats strfmt.Registry) error {

	if err := validate.Required("sourceEnvironmentId", "body", m.SourceEnvironmentID); err != nil {
		return err
	}

	return nil
}

func (m *CreatePromotionRequest) validateTargetEnvironmentID(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironmentId", "body", m.TargetEnvironmentID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create promotion request based on context it is used
func (m *CreatePromotionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreatePromotionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePromotionRequest) UnmarshalBinary(b []byte) error {
	var res CreatePromotionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
