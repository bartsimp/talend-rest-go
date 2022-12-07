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

// StepExecution step execution
//
// swagger:model StepExecution
type StepExecution struct {

	// execution Id
	ExecutionID string `json:"executionId,omitempty"`

	// execution status
	ExecutionStatus string `json:"executionStatus,omitempty"`

	// finish timestamp
	// Format: date-time
	FinishTimestamp strfmt.DateTime `json:"finishTimestamp,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// start timestamp
	// Format: date-time
	StartTimestamp strfmt.DateTime `json:"startTimestamp,omitempty"`
}

// Validate validates this step execution
func (m *StepExecution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StepExecution) validateFinishTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.FinishTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("finishTimestamp", "body", "date-time", m.FinishTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StepExecution) validateStartTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("startTimestamp", "body", "date-time", m.StartTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this step execution based on context it is used
func (m *StepExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StepExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StepExecution) UnmarshalBinary(b []byte) error {
	var res StepExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}