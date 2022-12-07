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

// PromotionExecutionStatus promotion execution status
//
// swagger:model PromotionExecutionStatus
type PromotionExecutionStatus struct {

	// Advanced Promotion specification
	Advanced *AdvancedPromotionSpec `json:"advanced,omitempty"`

	// Remote Engine Clusters Promotion Execution Info
	Clusters []*RuntimePromotionResult `json:"clusters"`

	// Execution context message
	Context string `json:"context,omitempty"`

	// Defective Promotion flag
	Defective bool `json:"defective,omitempty"`

	// Remote Engines Promotion Execution Info
	Engines []*RuntimePromotionResult `json:"engines"`

	// Job execution ID
	// Example: 7b2b122e-d6b8-42de-b0ba-fa2f0d36306e
	// Required: true
	ExecutionID *string `json:"executionId"`

	// End time of job execution (UTC)
	// Format: date-time
	FinishTimestamp strfmt.DateTime `json:"finishTimestamp,omitempty"`

	// Keep Target Resources flag
	KeepTargetResources bool `json:"keepTargetResources,omitempty"`

	// Keep Target Run Profiles flag (not returned for API versions earlier V2.2)
	KeepTargetRunProfiles bool `json:"keepTargetRunProfiles,omitempty"`

	// Deprecated! Use Promotion ID instead
	PipelineID string `json:"pipelineId,omitempty"`

	// Promotion ID
	PromotionID string `json:"promotionId,omitempty"`

	// Start time of job execution (UTC)
	// Required: true
	// Format: date-time
	StartTimestamp *strfmt.DateTime `json:"startTimestamp"`

	// Execution status
	// Example: PROMOTED
	// Enum: [ANALIZED PROMOTION_STARTED PROMOTION_WARNING PROMOTION_FAILED PROMOTED]
	Status string `json:"status,omitempty"`

	// Execution status message
	StatusMessage string `json:"statusMessage,omitempty"`

	// Trigger time of job execution (UTC)
	// Required: true
	// Format: date-time
	TriggerTimestamp *strfmt.DateTime `json:"triggerTimestamp"`

	// User who triggered or scheduled the execution
	// Example: fupton
	// Required: true
	UserID *string `json:"userId"`

	// Type of user who triggered or scheduled the execution
	// Example: HUMAN
	// Required: true
	// Enum: [HUMAN SERVICE]
	UserType *string `json:"userType"`

	// Workspace Promotion Execution Info
	Workspaces []*WorkspacePromotionResult `json:"workspaces"`
}

// Validate validates this promotion execution status
func (m *PromotionExecutionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvanced(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEngines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggerTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PromotionExecutionStatus) validateAdvanced(formats strfmt.Registry) error {
	if swag.IsZero(m.Advanced) { // not required
		return nil
	}

	if m.Advanced != nil {
		if err := m.Advanced.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advanced")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advanced")
			}
			return err
		}
	}

	return nil
}

func (m *PromotionExecutionStatus) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PromotionExecutionStatus) validateEngines(formats strfmt.Registry) error {
	if swag.IsZero(m.Engines) { // not required
		return nil
	}

	for i := 0; i < len(m.Engines); i++ {
		if swag.IsZero(m.Engines[i]) { // not required
			continue
		}

		if m.Engines[i] != nil {
			if err := m.Engines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("engines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("engines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PromotionExecutionStatus) validateExecutionID(formats strfmt.Registry) error {

	if err := validate.Required("executionId", "body", m.ExecutionID); err != nil {
		return err
	}

	return nil
}

func (m *PromotionExecutionStatus) validateFinishTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.FinishTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("finishTimestamp", "body", "date-time", m.FinishTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PromotionExecutionStatus) validateStartTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("startTimestamp", "body", m.StartTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("startTimestamp", "body", "date-time", m.StartTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

var promotionExecutionStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ANALIZED","PROMOTION_STARTED","PROMOTION_WARNING","PROMOTION_FAILED","PROMOTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		promotionExecutionStatusTypeStatusPropEnum = append(promotionExecutionStatusTypeStatusPropEnum, v)
	}
}

const (

	// PromotionExecutionStatusStatusANALIZED captures enum value "ANALIZED"
	PromotionExecutionStatusStatusANALIZED string = "ANALIZED"

	// PromotionExecutionStatusStatusPROMOTIONSTARTED captures enum value "PROMOTION_STARTED"
	PromotionExecutionStatusStatusPROMOTIONSTARTED string = "PROMOTION_STARTED"

	// PromotionExecutionStatusStatusPROMOTIONWARNING captures enum value "PROMOTION_WARNING"
	PromotionExecutionStatusStatusPROMOTIONWARNING string = "PROMOTION_WARNING"

	// PromotionExecutionStatusStatusPROMOTIONFAILED captures enum value "PROMOTION_FAILED"
	PromotionExecutionStatusStatusPROMOTIONFAILED string = "PROMOTION_FAILED"

	// PromotionExecutionStatusStatusPROMOTED captures enum value "PROMOTED"
	PromotionExecutionStatusStatusPROMOTED string = "PROMOTED"
)

// prop value enum
func (m *PromotionExecutionStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, promotionExecutionStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PromotionExecutionStatus) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PromotionExecutionStatus) validateTriggerTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("triggerTimestamp", "body", m.TriggerTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("triggerTimestamp", "body", "date-time", m.TriggerTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PromotionExecutionStatus) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

var promotionExecutionStatusTypeUserTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HUMAN","SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		promotionExecutionStatusTypeUserTypePropEnum = append(promotionExecutionStatusTypeUserTypePropEnum, v)
	}
}

const (

	// PromotionExecutionStatusUserTypeHUMAN captures enum value "HUMAN"
	PromotionExecutionStatusUserTypeHUMAN string = "HUMAN"

	// PromotionExecutionStatusUserTypeSERVICE captures enum value "SERVICE"
	PromotionExecutionStatusUserTypeSERVICE string = "SERVICE"
)

// prop value enum
func (m *PromotionExecutionStatus) validateUserTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, promotionExecutionStatusTypeUserTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PromotionExecutionStatus) validateUserType(formats strfmt.Registry) error {

	if err := validate.Required("userType", "body", m.UserType); err != nil {
		return err
	}

	// value enum
	if err := m.validateUserTypeEnum("userType", "body", *m.UserType); err != nil {
		return err
	}

	return nil
}

func (m *PromotionExecutionStatus) validateWorkspaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Workspaces); i++ {
		if swag.IsZero(m.Workspaces[i]) { // not required
			continue
		}

		if m.Workspaces[i] != nil {
			if err := m.Workspaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this promotion execution status based on the context it is used
func (m *PromotionExecutionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvanced(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEngines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PromotionExecutionStatus) contextValidateAdvanced(ctx context.Context, formats strfmt.Registry) error {

	if m.Advanced != nil {
		if err := m.Advanced.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advanced")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advanced")
			}
			return err
		}
	}

	return nil
}

func (m *PromotionExecutionStatus) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clusters); i++ {

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PromotionExecutionStatus) contextValidateEngines(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Engines); i++ {

		if m.Engines[i] != nil {
			if err := m.Engines[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("engines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("engines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PromotionExecutionStatus) contextValidateWorkspaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workspaces); i++ {

		if m.Workspaces[i] != nil {
			if err := m.Workspaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PromotionExecutionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PromotionExecutionStatus) UnmarshalBinary(b []byte) error {
	var res PromotionExecutionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
