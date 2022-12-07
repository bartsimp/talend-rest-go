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

// Artifact artifact
//
// swagger:model Artifact
type Artifact struct {

	// Artifact identifier
	// Example: 5cdd2ce2c737973f9d581b99
	// Required: true
	ID *string `json:"id"`

	// Artifact name
	// Example: Test Artifact
	// Required: true
	Name *string `json:"name"`

	// Artifact type
	// Example: standard
	// Required: true
	// Enum: [standard big_data_streaming big_data_batch route data-service]
	Type *string `json:"type"`

	// Available artifact versions
	// Example: ["0.1.0","0.1.1"]
	// Required: true
	Versions []string `json:"versions"`

	// Artifact workspace
	// Required: true
	Workspace *Workspaceinfo `json:"workspace"`
}

// Validate validates this artifact
func (m *Artifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
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

func (m *Artifact) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Artifact) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var artifactTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["standard","big_data_streaming","big_data_batch","route","data-service"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		artifactTypeTypePropEnum = append(artifactTypeTypePropEnum, v)
	}
}

const (

	// ArtifactTypeStandard captures enum value "standard"
	ArtifactTypeStandard string = "standard"

	// ArtifactTypeBigDataStreaming captures enum value "big_data_streaming"
	ArtifactTypeBigDataStreaming string = "big_data_streaming"

	// ArtifactTypeBigDataBatch captures enum value "big_data_batch"
	ArtifactTypeBigDataBatch string = "big_data_batch"

	// ArtifactTypeRoute captures enum value "route"
	ArtifactTypeRoute string = "route"

	// ArtifactTypeDataDashService captures enum value "data-service"
	ArtifactTypeDataDashService string = "data-service"
)

// prop value enum
func (m *Artifact) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, artifactTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Artifact) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Artifact) validateVersions(formats strfmt.Registry) error {

	if err := validate.Required("versions", "body", m.Versions); err != nil {
		return err
	}

	return nil
}

func (m *Artifact) validateWorkspace(formats strfmt.Registry) error {

	if err := validate.Required("workspace", "body", m.Workspace); err != nil {
		return err
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

// ContextValidate validate this artifact based on the context it is used
func (m *Artifact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Artifact) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Artifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Artifact) UnmarshalBinary(b []byte) error {
	var res Artifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
