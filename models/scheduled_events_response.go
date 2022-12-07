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

// ScheduledEventsResponse scheduled events response
//
// swagger:model ScheduledEventsResponse
type ScheduledEventsResponse struct {

	// The Trigger's date for the event
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// The entity's id for this event
	ID string `json:"id,omitempty"`

	// The source trigger's name (only for schedule for multiple triggers)
	SourceTriggerName string `json:"sourceTriggerName,omitempty"`

	// The entity's type for the event
	// Enum: [TASK PLAN]
	Type string `json:"type,omitempty"`
}

// Validate validates this scheduled events response
func (m *ScheduledEventsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
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

func (m *ScheduledEventsResponse) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

var scheduledEventsResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TASK","PLAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduledEventsResponseTypeTypePropEnum = append(scheduledEventsResponseTypeTypePropEnum, v)
	}
}

const (

	// ScheduledEventsResponseTypeTASK captures enum value "TASK"
	ScheduledEventsResponseTypeTASK string = "TASK"

	// ScheduledEventsResponseTypePLAN captures enum value "PLAN"
	ScheduledEventsResponseTypePLAN string = "PLAN"
)

// prop value enum
func (m *ScheduledEventsResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scheduledEventsResponseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScheduledEventsResponse) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this scheduled events response based on context it is used
func (m *ScheduledEventsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScheduledEventsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledEventsResponse) UnmarshalBinary(b []byte) error {
	var res ScheduledEventsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
