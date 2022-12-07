// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DebugRemoteEngineModel debug remote engine model
//
// swagger:model DebugRemoteEngineModel
type DebugRemoteEngineModel struct {

	// Host/ip
	Host string `json:"host,omitempty"`

	// Is debug mode on
	IsDebugModeOn bool `json:"isDebugModeOn,omitempty"`
}

// Validate validates this debug remote engine model
func (m *DebugRemoteEngineModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this debug remote engine model based on context it is used
func (m *DebugRemoteEngineModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DebugRemoteEngineModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebugRemoteEngineModel) UnmarshalBinary(b []byte) error {
	var res DebugRemoteEngineModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
