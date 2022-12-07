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

// SourcesRepositoryData sources repository data
//
// swagger:model SourcesRepositoryData
type SourcesRepositoryData struct {

	// repository branch
	Branch string `json:"branch,omitempty"`

	// repository commit info
	Commit *RepositoryCommitData `json:"commit,omitempty"`

	// project name
	// Required: true
	Project *string `json:"project"`
}

// Validate validates this sources repository data
func (m *SourcesRepositoryData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourcesRepositoryData) validateCommit(formats strfmt.Registry) error {
	if swag.IsZero(m.Commit) { // not required
		return nil
	}

	if m.Commit != nil {
		if err := m.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

func (m *SourcesRepositoryData) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sources repository data based on the context it is used
func (m *SourcesRepositoryData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourcesRepositoryData) contextValidateCommit(ctx context.Context, formats strfmt.Registry) error {

	if m.Commit != nil {
		if err := m.Commit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourcesRepositoryData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourcesRepositoryData) UnmarshalBinary(b []byte) error {
	var res SourcesRepositoryData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}