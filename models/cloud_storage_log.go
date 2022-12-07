// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudStorageLog cloud storage log
//
// swagger:model CloudStorageLog
type CloudStorageLog struct {

	// bundle name
	BundleName string `json:"bundleName,omitempty"`

	// bundle version
	BundleVersion string `json:"bundleVersion,omitempty"`

	// log message
	LogMessage string `json:"logMessage,omitempty"`

	// log timestamp
	LogTimestamp int64 `json:"logTimestamp,omitempty"`

	// log type
	LogType string `json:"logType,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`
}

// Validate validates this cloud storage log
func (m *CloudStorageLog) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud storage log based on context it is used
func (m *CloudStorageLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudStorageLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStorageLog) UnmarshalBinary(b []byte) error {
	var res CloudStorageLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
