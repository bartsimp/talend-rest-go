// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlanExecutabledetails plan executabledetails
//
// swagger:model PlanExecutabledetails
type PlanExecutabledetails struct {

	// Plan execution hierarchy
	Chart *ExecutionStep `json:"chart,omitempty"`

	// Executable description
	// Example: plan detail description
	Description string `json:"description,omitempty"`

	// Executable identifier
	// Example: 57f64991e4b0b689a64feed0
	// Required: true
	Executable *string `json:"executable"`

	// Executable name
	// Example: simple executable
	// Required: true
	Name *string `json:"name"`

	// Pause details of Plan (available from V2.6)
	PlanPauseDetails *PlanPauseDetails `json:"planPauseDetails,omitempty"`

	// Plan execution status
	// Enum: [executing execution_successful execution_failed]
	Status string `json:"status,omitempty"`

	// Executable workspace
	Workspace *Workspaceinfo `json:"workspace,omitempty"`
}

// Validate validates this plan executabledetails
func (m *PlanExecutabledetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanPauseDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanExecutabledetails) validateChart(formats strfmt.Registry) error {
	if swag.IsZero(m.Chart) { // not required
		return nil
	}

	if m.Chart != nil {
		if err := m.Chart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chart")
			}
			return err
		}
	}

	return nil
}

func (m *PlanExecutabledetails) validateExecutable(formats strfmt.Registry) error {

	if err := validate.Required("executable", "body", m.Executable); err != nil {
		return err
	}

	return nil
}

func (m *PlanExecutabledetails) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PlanExecutabledetails) validatePlanPauseDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.PlanPauseDetails) { // not required
		return nil
	}

	if m.PlanPauseDetails != nil {
		if err := m.PlanPauseDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("planPauseDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("planPauseDetails")
			}
			return err
		}
	}

	return nil
}

var planExecutabledetailsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["executing","execution_successful","execution_failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planExecutabledetailsTypeStatusPropEnum = append(planExecutabledetailsTypeStatusPropEnum, v)
	}
}

const (

	// PlanExecutabledetailsStatusExecuting captures enum value "executing"
	PlanExecutabledetailsStatusExecuting string = "executing"

	// PlanExecutabledetailsStatusExecutionSuccessful captures enum value "execution_successful"
	PlanExecutabledetailsStatusExecutionSuccessful string = "execution_successful"

	// PlanExecutabledetailsStatusExecutionFailed captures enum value "execution_failed"
	PlanExecutabledetailsStatusExecutionFailed string = "execution_failed"
)

// prop value enum
func (m *PlanExecutabledetails) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, planExecutabledetailsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlanExecutabledetails) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PlanExecutabledetails) validateWorkspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plan executabledetails based on the context it is used
func (m *PlanExecutabledetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlanPauseDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanExecutabledetails) contextValidateChart(ctx context.Context, formats strfmt.Registry) error {

	if m.Chart != nil {
		if err := m.Chart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chart")
			}
			return err
		}
	}

	return nil
}

func (m *PlanExecutabledetails) contextValidatePlanPauseDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.PlanPauseDetails != nil {
		if err := m.PlanPauseDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("planPauseDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("planPauseDetails")
			}
			return err
		}
	}

	return nil
}

func (m *PlanExecutabledetails) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Workspace != nil {
		if err := m.Workspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanExecutabledetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanExecutabledetails) UnmarshalBinary(b []byte) error {
	var res PlanExecutabledetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}