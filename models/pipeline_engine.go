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

// PipelineEngine pipeline engine
//
// swagger:model PipelineEngine
type PipelineEngine struct {

	// Availability status of engine|cluster
	// Example: NOT_AVAILABLE
	// Enum: [AVAILABLE NOT_AVAILABLE RETIRED]
	Availability string `json:"availability,omitempty"`

	// cloud runner
	CloudRunner bool `json:"cloudRunner,omitempty"`

	// Created on
	// Read Only: true
	// Format: date-time
	CreateDate strfmt.DateTime `json:"createDate,omitempty"`

	// Resource description
	// Example: Some description
	Description string `json:"description,omitempty"`

	// Resource id
	// Example: 5cdd2ce2c737973f9d581b98
	// Required: true
	ID *string `json:"id"`

	// Indicates whether target runtime (engine/cluster) is managed or not
	Managed bool `json:"managed,omitempty"`

	// Resource name
	// Example: Remote Engine
	// Required: true
	Name *string `json:"name"`

	// pre authorized key
	PreAuthorizedKey string `json:"preAuthorizedKey,omitempty"`

	// Resource runtime id
	// Example: 9e8acf72-00d5-451b-9175-97bd44cd6b13
	// Required: true
	RuntimeID *string `json:"runtimeId"`

	// status
	// Enum: [PAIRED NOT_PAIRED]
	Status string `json:"status,omitempty"`

	// Updated on
	// Read Only: true
	// Format: date-time
	UpdateDate strfmt.DateTime `json:"updateDate,omitempty"`

	// Resource workspace
	Workspace *Workspaceinfo `json:"workspace,omitempty"`
}

// Validate validates this pipeline engine
func (m *PipelineEngine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntimeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateDate(formats); err != nil {
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

var pipelineEngineTypeAvailabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AVAILABLE","NOT_AVAILABLE","RETIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pipelineEngineTypeAvailabilityPropEnum = append(pipelineEngineTypeAvailabilityPropEnum, v)
	}
}

const (

	// PipelineEngineAvailabilityAVAILABLE captures enum value "AVAILABLE"
	PipelineEngineAvailabilityAVAILABLE string = "AVAILABLE"

	// PipelineEngineAvailabilityNOTAVAILABLE captures enum value "NOT_AVAILABLE"
	PipelineEngineAvailabilityNOTAVAILABLE string = "NOT_AVAILABLE"

	// PipelineEngineAvailabilityRETIRED captures enum value "RETIRED"
	PipelineEngineAvailabilityRETIRED string = "RETIRED"
)

// prop value enum
func (m *PipelineEngine) validateAvailabilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pipelineEngineTypeAvailabilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PipelineEngine) validateAvailability(formats strfmt.Registry) error {
	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvailabilityEnum("availability", "body", m.Availability); err != nil {
		return err
	}

	return nil
}

func (m *PipelineEngine) validateCreateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createDate", "body", "date-time", m.CreateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PipelineEngine) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PipelineEngine) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PipelineEngine) validateRuntimeID(formats strfmt.Registry) error {

	if err := validate.Required("runtimeId", "body", m.RuntimeID); err != nil {
		return err
	}

	return nil
}

var pipelineEngineTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PAIRED","NOT_PAIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pipelineEngineTypeStatusPropEnum = append(pipelineEngineTypeStatusPropEnum, v)
	}
}

const (

	// PipelineEngineStatusPAIRED captures enum value "PAIRED"
	PipelineEngineStatusPAIRED string = "PAIRED"

	// PipelineEngineStatusNOTPAIRED captures enum value "NOT_PAIRED"
	PipelineEngineStatusNOTPAIRED string = "NOT_PAIRED"
)

// prop value enum
func (m *PipelineEngine) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pipelineEngineTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PipelineEngine) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PipelineEngine) validateUpdateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updateDate", "body", "date-time", m.UpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PipelineEngine) validateWorkspace(formats strfmt.Registry) error {
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

// ContextValidate validate this pipeline engine based on the context it is used
func (m *PipelineEngine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreateDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateDate(ctx, formats); err != nil {
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

func (m *PipelineEngine) contextValidateCreateDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createDate", "body", strfmt.DateTime(m.CreateDate)); err != nil {
		return err
	}

	return nil
}

func (m *PipelineEngine) contextValidateUpdateDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updateDate", "body", strfmt.DateTime(m.UpdateDate)); err != nil {
		return err
	}

	return nil
}

func (m *PipelineEngine) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PipelineEngine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineEngine) UnmarshalBinary(b []byte) error {
	var res PipelineEngine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}