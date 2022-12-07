// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectGroupAuthResponse project group auth response
//
// swagger:model ProjectGroupAuthResponse
type ProjectGroupAuthResponse struct {

	// group ids
	GroupIds []string `json:"groupIds"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`
}

// Validate validates this project group auth response
func (m *ProjectGroupAuthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project group auth response based on context it is used
func (m *ProjectGroupAuthResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectGroupAuthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectGroupAuthResponse) UnmarshalBinary(b []byte) error {
	var res ProjectGroupAuthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}