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

// Workspaceinfo workspaceinfo
//
// swagger:model Workspaceinfo
type Workspaceinfo struct {

	// Workspace description
	// Example: workspace detail description
	Description string `json:"description,omitempty"`

	// Workspace environment
	Environment *EnvironmentInfo `json:"environment,omitempty"`

	// Workspace identifier
	// Example: 57f64991e4b0b689a64feed2
	// Required: true
	ID *string `json:"id"`

	// Workspace name
	// Example: Personal
	// Required: true
	Name *string `json:"name"`

	// Workspace owner
	// Example: admin
	Owner string `json:"owner,omitempty"`

	// Workspace type
	// Example: custom
	// Required: true
	// Enum: [shared personal custom]
	Type *string `json:"type"`
}

// Validate validates this workspaceinfo
func (m *Workspaceinfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *Workspaceinfo) validateEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *Workspaceinfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Workspaceinfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var workspaceinfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["shared","personal","custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workspaceinfoTypeTypePropEnum = append(workspaceinfoTypeTypePropEnum, v)
	}
}

const (

	// WorkspaceinfoTypeShared captures enum value "shared"
	WorkspaceinfoTypeShared string = "shared"

	// WorkspaceinfoTypePersonal captures enum value "personal"
	WorkspaceinfoTypePersonal string = "personal"

	// WorkspaceinfoTypeCustom captures enum value "custom"
	WorkspaceinfoTypeCustom string = "custom"
)

// prop value enum
func (m *Workspaceinfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, workspaceinfoTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Workspaceinfo) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this workspaceinfo based on the context it is used
func (m *Workspaceinfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnvironment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workspaceinfo) contextValidateEnvironment(ctx context.Context, formats strfmt.Registry) error {

	if m.Environment != nil {
		if err := m.Environment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Workspaceinfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workspaceinfo) UnmarshalBinary(b []byte) error {
	var res Workspaceinfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}