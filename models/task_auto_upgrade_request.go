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

// TaskAutoUpgradeRequest task auto upgrade request
//
// swagger:model TaskAutoUpgradeRequest
type TaskAutoUpgradeRequest struct {

	// Artifact used in task
	// Required: true
	Artifact *ArtifactRequest `json:"artifact"`

	// Auto upgrade
	AutoUpgradeInfo *AutoUpgradeInfo `json:"autoUpgradeInfo,omitempty"`

	// Task connections
	// Example: {"aws":"5d405e381f40e07fbab6757c","googledrive":"5d405e831f40e07fbab6757d"}
	Connections map[string]string `json:"connections,omitempty"`

	// Task description
	// Example: Task detail description
	// Required: true
	Description *string `json:"description"`

	// Task name
	// Example: Hello world task
	// Required: true
	Name *string `json:"name"`

	// Task parameters
	// Example: {"custom_parameter":"custom","parameter_p1":"1111"}
	Parameters map[string]string `json:"parameters,omitempty"`

	// Task resources
	// Example: {"resource_file_f1":"5d84adcf38907f16498d5d74"}
	Resources map[string]string `json:"resources,omitempty"`

	// Task tags
	// Example: ["tag1","tag2"]
	Tags []string `json:"tags"`

	// Workspace id of task to create
	// Example: 57f64991e4b0b689a64feed2
	// Required: true
	WorkspaceID *string `json:"workspaceId"`
}

// Validate validates this task auto upgrade request
func (m *TaskAutoUpgradeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoUpgradeInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAutoUpgradeRequest) validateArtifact(formats strfmt.Registry) error {

	if err := validate.Required("artifact", "body", m.Artifact); err != nil {
		return err
	}

	if m.Artifact != nil {
		if err := m.Artifact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifact")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAutoUpgradeRequest) validateAutoUpgradeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoUpgradeInfo) { // not required
		return nil
	}

	if m.AutoUpgradeInfo != nil {
		if err := m.AutoUpgradeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoUpgradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoUpgradeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAutoUpgradeRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *TaskAutoUpgradeRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TaskAutoUpgradeRequest) validateWorkspaceID(formats strfmt.Registry) error {

	if err := validate.Required("workspaceId", "body", m.WorkspaceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this task auto upgrade request based on the context it is used
func (m *TaskAutoUpgradeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArtifact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutoUpgradeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAutoUpgradeRequest) contextValidateArtifact(ctx context.Context, formats strfmt.Registry) error {

	if m.Artifact != nil {
		if err := m.Artifact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifact")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAutoUpgradeRequest) contextValidateAutoUpgradeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoUpgradeInfo != nil {
		if err := m.AutoUpgradeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoUpgradeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoUpgradeInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAutoUpgradeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAutoUpgradeRequest) UnmarshalBinary(b []byte) error {
	var res TaskAutoUpgradeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
