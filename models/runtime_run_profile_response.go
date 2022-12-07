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

// RuntimeRunProfileResponse runtime run profile response
//
// swagger:model RuntimeRunProfileResponse
type RuntimeRunProfileResponse struct {

	// Created on
	// Read Only: true
	// Format: date-time
	CreateDate strfmt.DateTime `json:"createDate,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Jvm arguments
	JvmArguments []string `json:"jvmArguments"`

	// name
	Name string `json:"name,omitempty"`

	// runtime Id
	RuntimeID string `json:"runtimeId,omitempty"`

	// Run profile type
	// Enum: [JOB_SERVER MICROSERVICE TALEND_RUNTIME]
	Type string `json:"type,omitempty"`

	// Updated on
	// Read Only: true
	// Format: date-time
	UpdateDate strfmt.DateTime `json:"updateDate,omitempty"`

	// Run profile version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this runtime run profile response
func (m *RuntimeRunProfileResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuntimeRunProfileResponse) validateCreateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createDate", "body", "date-time", m.CreateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var runtimeRunProfileResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["JOB_SERVER","MICROSERVICE","TALEND_RUNTIME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runtimeRunProfileResponseTypeTypePropEnum = append(runtimeRunProfileResponseTypeTypePropEnum, v)
	}
}

const (

	// RuntimeRunProfileResponseTypeJOBSERVER captures enum value "JOB_SERVER"
	RuntimeRunProfileResponseTypeJOBSERVER string = "JOB_SERVER"

	// RuntimeRunProfileResponseTypeMICROSERVICE captures enum value "MICROSERVICE"
	RuntimeRunProfileResponseTypeMICROSERVICE string = "MICROSERVICE"

	// RuntimeRunProfileResponseTypeTALENDRUNTIME captures enum value "TALEND_RUNTIME"
	RuntimeRunProfileResponseTypeTALENDRUNTIME string = "TALEND_RUNTIME"
)

// prop value enum
func (m *RuntimeRunProfileResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, runtimeRunProfileResponseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RuntimeRunProfileResponse) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeRunProfileResponse) validateUpdateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updateDate", "body", "date-time", m.UpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this runtime run profile response based on the context it is used
func (m *RuntimeRunProfileResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreateDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuntimeRunProfileResponse) contextValidateCreateDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createDate", "body", strfmt.DateTime(m.CreateDate)); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeRunProfileResponse) contextValidateUpdateDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updateDate", "body", strfmt.DateTime(m.UpdateDate)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuntimeRunProfileResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeRunProfileResponse) UnmarshalBinary(b []byte) error {
	var res RuntimeRunProfileResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
