// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateScheduleMultipleTriggerRequest update schedule multiple trigger request
//
// swagger:model UpdateScheduleMultipleTriggerRequest
type UpdateScheduleMultipleTriggerRequest struct {

	// Some description related to the usage of the scheduler
	// Example: Schedule with five triggers
	Description string `json:"description,omitempty"`
}

// Validate validates this update schedule multiple trigger request
func (m *UpdateScheduleMultipleTriggerRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update schedule multiple trigger request based on context it is used
func (m *UpdateScheduleMultipleTriggerRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateScheduleMultipleTriggerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateScheduleMultipleTriggerRequest) UnmarshalBinary(b []byte) error {
	var res UpdateScheduleMultipleTriggerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
