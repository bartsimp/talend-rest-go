// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectUserAuthResponse project user auth response
//
// swagger:model ProjectUserAuthResponse
type ProjectUserAuthResponse struct {

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// user ids
	UserIds []string `json:"userIds"`
}

// Validate validates this project user auth response
func (m *ProjectUserAuthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project user auth response based on context it is used
func (m *ProjectUserAuthResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectUserAuthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectUserAuthResponse) UnmarshalBinary(b []byte) error {
	var res ProjectUserAuthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}