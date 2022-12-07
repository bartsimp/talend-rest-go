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

// PlanExecutionStatusV21 Plan Execution info
//
// swagger:model PlanExecutionStatusV21
type PlanExecutionStatusV21 struct {

	// Number of done executables
	// Example: 1
	DoneExecutableCount int32 `json:"doneExecutableCount,omitempty"`

	// Execution activity info
	DoneExecutableDetails []*JobExecutionStatus `json:"doneExecutableDetails"`

	// Job execution ID
	// Example: 7b2b122e-d6b8-42de-b0ba-fa2f0d36306e
	// Required: true
	ExecutionID *string `json:"executionId"`

	// Execution (detailed) status
	// Example: FINISHED
	// Required: true
	// Enum: [UNDEFINED STARTED FINISHED EXECUTION_FAILED EXECUTION_SUCCESS]
	ExecutionStatus *string `json:"executionStatus"`

	// End time of job execution
	// Format: date-time
	FinishTimestamp strfmt.DateTime `json:"finishTimestamp,omitempty"`

	// Plan ID
	// Example: b91cf8b2-5dd1-4b18-915b-4c447cee5267
	// Required: true
	PlanID *string `json:"planId"`

	// Number of planned executables
	// Example: 2
	PlannedExecutableCount int32 `json:"plannedExecutableCount,omitempty"`

	// Start time of job execution
	// Required: true
	// Format: date-time
	StartTimestamp *strfmt.DateTime `json:"startTimestamp"`

	// Execution status
	// Example: executing
	// Required: true
	// Enum: [executing execution_successful execution_failed]
	Status *string `json:"status"`

	// User who triggered or scheduled the execution
	// Example: fupton
	// Required: true
	UserID *string `json:"userId"`

	// Type of user who triggered or scheduled the execution
	// Example: HUMAN
	// Required: true
	// Enum: [HUMAN SERVICE]
	UserType *string `json:"userType"`
}

// Validate validates this plan execution status v21
func (m *PlanExecutionStatusV21) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDoneExecutableDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanExecutionStatusV21) validateDoneExecutableDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.DoneExecutableDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.DoneExecutableDetails); i++ {
		if swag.IsZero(m.DoneExecutableDetails[i]) { // not required
			continue
		}

		if m.DoneExecutableDetails[i] != nil {
			if err := m.DoneExecutableDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("doneExecutableDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("doneExecutableDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlanExecutionStatusV21) validateExecutionID(formats strfmt.Registry) error {

	if err := validate.Required("executionId", "body", m.ExecutionID); err != nil {
		return err
	}

	return nil
}

var planExecutionStatusV21TypeExecutionStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNDEFINED","STARTED","FINISHED","EXECUTION_FAILED","EXECUTION_SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planExecutionStatusV21TypeExecutionStatusPropEnum = append(planExecutionStatusV21TypeExecutionStatusPropEnum, v)
	}
}

const (

	// PlanExecutionStatusV21ExecutionStatusUNDEFINED captures enum value "UNDEFINED"
	PlanExecutionStatusV21ExecutionStatusUNDEFINED string = "UNDEFINED"

	// PlanExecutionStatusV21ExecutionStatusSTARTED captures enum value "STARTED"
	PlanExecutionStatusV21ExecutionStatusSTARTED string = "STARTED"

	// PlanExecutionStatusV21ExecutionStatusFINISHED captures enum value "FINISHED"
	PlanExecutionStatusV21ExecutionStatusFINISHED string = "FINISHED"

	// PlanExecutionStatusV21ExecutionStatusEXECUTIONFAILED captures enum value "EXECUTION_FAILED"
	PlanExecutionStatusV21ExecutionStatusEXECUTIONFAILED string = "EXECUTION_FAILED"

	// PlanExecutionStatusV21ExecutionStatusEXECUTIONSUCCESS captures enum value "EXECUTION_SUCCESS"
	PlanExecutionStatusV21ExecutionStatusEXECUTIONSUCCESS string = "EXECUTION_SUCCESS"
)

// prop value enum
func (m *PlanExecutionStatusV21) validateExecutionStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, planExecutionStatusV21TypeExecutionStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlanExecutionStatusV21) validateExecutionStatus(formats strfmt.Registry) error {

	if err := validate.Required("executionStatus", "body", m.ExecutionStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateExecutionStatusEnum("executionStatus", "body", *m.ExecutionStatus); err != nil {
		return err
	}

	return nil
}

func (m *PlanExecutionStatusV21) validateFinishTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.FinishTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("finishTimestamp", "body", "date-time", m.FinishTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PlanExecutionStatusV21) validatePlanID(formats strfmt.Registry) error {

	if err := validate.Required("planId", "body", m.PlanID); err != nil {
		return err
	}

	return nil
}

func (m *PlanExecutionStatusV21) validateStartTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("startTimestamp", "body", m.StartTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("startTimestamp", "body", "date-time", m.StartTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

var planExecutionStatusV21TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["executing","execution_successful","execution_failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planExecutionStatusV21TypeStatusPropEnum = append(planExecutionStatusV21TypeStatusPropEnum, v)
	}
}

const (

	// PlanExecutionStatusV21StatusExecuting captures enum value "executing"
	PlanExecutionStatusV21StatusExecuting string = "executing"

	// PlanExecutionStatusV21StatusExecutionSuccessful captures enum value "execution_successful"
	PlanExecutionStatusV21StatusExecutionSuccessful string = "execution_successful"

	// PlanExecutionStatusV21StatusExecutionFailed captures enum value "execution_failed"
	PlanExecutionStatusV21StatusExecutionFailed string = "execution_failed"
)

// prop value enum
func (m *PlanExecutionStatusV21) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, planExecutionStatusV21TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlanExecutionStatusV21) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PlanExecutionStatusV21) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

var planExecutionStatusV21TypeUserTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HUMAN","SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planExecutionStatusV21TypeUserTypePropEnum = append(planExecutionStatusV21TypeUserTypePropEnum, v)
	}
}

const (

	// PlanExecutionStatusV21UserTypeHUMAN captures enum value "HUMAN"
	PlanExecutionStatusV21UserTypeHUMAN string = "HUMAN"

	// PlanExecutionStatusV21UserTypeSERVICE captures enum value "SERVICE"
	PlanExecutionStatusV21UserTypeSERVICE string = "SERVICE"
)

// prop value enum
func (m *PlanExecutionStatusV21) validateUserTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, planExecutionStatusV21TypeUserTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlanExecutionStatusV21) validateUserType(formats strfmt.Registry) error {

	if err := validate.Required("userType", "body", m.UserType); err != nil {
		return err
	}

	// value enum
	if err := m.validateUserTypeEnum("userType", "body", *m.UserType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this plan execution status v21 based on the context it is used
func (m *PlanExecutionStatusV21) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDoneExecutableDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanExecutionStatusV21) contextValidateDoneExecutableDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DoneExecutableDetails); i++ {

		if m.DoneExecutableDetails[i] != nil {
			if err := m.DoneExecutableDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("doneExecutableDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("doneExecutableDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanExecutionStatusV21) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanExecutionStatusV21) UnmarshalBinary(b []byte) error {
	var res PlanExecutionStatusV21
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}