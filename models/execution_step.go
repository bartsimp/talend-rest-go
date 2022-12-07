// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExecutionStep execution step
//
// swagger:model ExecutionStep
type ExecutionStep struct {

	// abstract step Id
	AbstractStepID string `json:"abstractStepId,omitempty"`

	// flows
	Flows []*ExecutionFlow `json:"flows"`

	// next step
	NextStep *ExecutionStep `json:"nextStep,omitempty"`

	// status
	// Enum: [SUCCESS PARTIAL_SUCCESS FAIL]
	Status string `json:"status,omitempty"`

	// step Id
	StepID string `json:"stepId,omitempty"`

	// step name
	StepName string `json:"stepName,omitempty"`

	// step note
	StepNote string `json:"stepNote,omitempty"`

	// step on exception
	StepOnException *ExecutionStep `json:"stepOnException,omitempty"`
}

// Validate validates this execution step
func (m *ExecutionStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlows(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextStep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStepOnException(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecutionStep) validateFlows(formats strfmt.Registry) error {
	if swag.IsZero(m.Flows) { // not required
		return nil
	}

	for i := 0; i < len(m.Flows); i++ {
		if swag.IsZero(m.Flows[i]) { // not required
			continue
		}

		if m.Flows[i] != nil {
			if err := m.Flows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExecutionStep) validateNextStep(formats strfmt.Registry) error {
	if swag.IsZero(m.NextStep) { // not required
		return nil
	}

	if m.NextStep != nil {
		if err := m.NextStep.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nextStep")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nextStep")
			}
			return err
		}
	}

	return nil
}

var executionStepTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","PARTIAL_SUCCESS","FAIL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		executionStepTypeStatusPropEnum = append(executionStepTypeStatusPropEnum, v)
	}
}

const (

	// ExecutionStepStatusSUCCESS captures enum value "SUCCESS"
	ExecutionStepStatusSUCCESS string = "SUCCESS"

	// ExecutionStepStatusPARTIALSUCCESS captures enum value "PARTIAL_SUCCESS"
	ExecutionStepStatusPARTIALSUCCESS string = "PARTIAL_SUCCESS"

	// ExecutionStepStatusFAIL captures enum value "FAIL"
	ExecutionStepStatusFAIL string = "FAIL"
)

// prop value enum
func (m *ExecutionStep) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, executionStepTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExecutionStep) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ExecutionStep) validateStepOnException(formats strfmt.Registry) error {
	if swag.IsZero(m.StepOnException) { // not required
		return nil
	}

	if m.StepOnException != nil {
		if err := m.StepOnException.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stepOnException")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stepOnException")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this execution step based on the context it is used
func (m *ExecutionStep) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFlows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextStep(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStepOnException(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecutionStep) contextValidateFlows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Flows); i++ {

		if m.Flows[i] != nil {
			if err := m.Flows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExecutionStep) contextValidateNextStep(ctx context.Context, formats strfmt.Registry) error {

	if m.NextStep != nil {
		if err := m.NextStep.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nextStep")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nextStep")
			}
			return err
		}
	}

	return nil
}

func (m *ExecutionStep) contextValidateStepOnException(ctx context.Context, formats strfmt.Registry) error {

	if m.StepOnException != nil {
		if err := m.StepOnException.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stepOnException")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stepOnException")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExecutionStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecutionStep) UnmarshalBinary(b []byte) error {
	var res ExecutionStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}