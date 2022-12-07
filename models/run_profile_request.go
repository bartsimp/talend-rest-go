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

// RunProfileRequest run profile request
//
// swagger:model RunProfileRequest
type RunProfileRequest struct {

	// Description
	// Example: Run profile with default options
	Description string `json:"description,omitempty"`

	// Jvm arguments
	// Example: ["-Xms128m","-Xmx256m"]
	// Required: true
	JvmArguments []string `json:"jvmArguments"`

	// Run profile name
	// Example: Run profile
	// Required: true
	Name *string `json:"name"`

	// Runtime id
	// Example: 5e713618de3c75a450447262
	// Required: true
	RuntimeID *string `json:"runtimeId"`

	// Run profile type
	// Example: JOB_SERVER
	// Required: true
	// Enum: [JOB_SERVER MICROSERVICE TALEND_RUNTIME]
	Type *string `json:"type"`
}

// Validate validates this run profile request
func (m *RunProfileRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJvmArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntimeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunProfileRequest) validateJvmArguments(formats strfmt.Registry) error {

	if err := validate.Required("jvmArguments", "body", m.JvmArguments); err != nil {
		return err
	}

	return nil
}

func (m *RunProfileRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RunProfileRequest) validateRuntimeID(formats strfmt.Registry) error {

	if err := validate.Required("runtimeId", "body", m.RuntimeID); err != nil {
		return err
	}

	return nil
}

var runProfileRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["JOB_SERVER","MICROSERVICE","TALEND_RUNTIME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runProfileRequestTypeTypePropEnum = append(runProfileRequestTypeTypePropEnum, v)
	}
}

const (

	// RunProfileRequestTypeJOBSERVER captures enum value "JOB_SERVER"
	RunProfileRequestTypeJOBSERVER string = "JOB_SERVER"

	// RunProfileRequestTypeMICROSERVICE captures enum value "MICROSERVICE"
	RunProfileRequestTypeMICROSERVICE string = "MICROSERVICE"

	// RunProfileRequestTypeTALENDRUNTIME captures enum value "TALEND_RUNTIME"
	RunProfileRequestTypeTALENDRUNTIME string = "TALEND_RUNTIME"
)

// prop value enum
func (m *RunProfileRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, runProfileRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RunProfileRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this run profile request based on context it is used
func (m *RunProfileRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunProfileRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunProfileRequest) UnmarshalBinary(b []byte) error {
	var res RunProfileRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
