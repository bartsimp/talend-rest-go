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

// DaySchedule day schedule
//
// swagger:model DaySchedule
type DaySchedule struct {

	// Day of month to run task/plan, required only if type of schedule at days is equal to DAY_OF_MONTH
	// Example: 3
	Day int32 `json:"day,omitempty"`

	// List of days of week to run task/plan, required only if type of schedule at days is equal to DAY_OF_WEEK
	// Example: ["MONDAY","WEDNESDAY","FRIDAY"]
	Days []string `json:"days"`

	// Type of schedule at days
	// Example: DAY_OF_WEEK
	// Required: true
	// Enum: [DAY_OF_WEEK DAY_OF_MONTH]
	Type *string `json:"type"`
}

// Validate validates this day schedule
func (m *DaySchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dayScheduleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DAY_OF_WEEK","DAY_OF_MONTH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dayScheduleTypeTypePropEnum = append(dayScheduleTypeTypePropEnum, v)
	}
}

const (

	// DayScheduleTypeDAYOFWEEK captures enum value "DAY_OF_WEEK"
	DayScheduleTypeDAYOFWEEK string = "DAY_OF_WEEK"

	// DayScheduleTypeDAYOFMONTH captures enum value "DAY_OF_MONTH"
	DayScheduleTypeDAYOFMONTH string = "DAY_OF_MONTH"
)

// prop value enum
func (m *DaySchedule) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dayScheduleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DaySchedule) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this day schedule based on context it is used
func (m *DaySchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DaySchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DaySchedule) UnmarshalBinary(b []byte) error {
	var res DaySchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}