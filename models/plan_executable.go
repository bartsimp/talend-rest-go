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

// PlanExecutable plan executable
//
// swagger:model PlanExecutable
type PlanExecutable struct {

	// Executable identifier
	// Example: b91cf8b2-5dd1-4b18-915b-4c447cee5267
	// Required: true
	Executable *string `json:"executable"`

	// The plan execution identifier in case of re-execution (available from V2.4)
	// Example: 0798b8d1-0e12-472f-be02-a0f04e792daa
	ExecutionPlanID string `json:"executionPlanId,omitempty"`

	// Optionally run only failing tasks (available from V2.4)
	// Example: true
	RerunOnlyFailedTasks bool `json:"rerunOnlyFailedTasks,omitempty"`

	// The execution step identifier (executionStepId) from which restart plan execution (available from V2.4)
	// Example: 09043c9f-02d0-41f6-b3cb-0ea53ffde377
	StepID string `json:"stepId,omitempty"`
}

// Validate validates this plan executable
func (m *PlanExecutable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanExecutable) validateExecutable(formats strfmt.Registry) error {

	if err := validate.Required("executable", "body", m.Executable); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plan executable based on context it is used
func (m *PlanExecutable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanExecutable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanExecutable) UnmarshalBinary(b []byte) error {
	var res PlanExecutable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
