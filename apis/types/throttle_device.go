// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ThrottleDevice throttle device
// swagger:model ThrottleDevice
type ThrottleDevice struct {

	// Device path
	Path string `json:"Path,omitempty"`

	// Rate
	// Minimum: 0
	Rate uint64 `json:"Rate,omitempty"`
}

// Validate validates this throttle device
func (m *ThrottleDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThrottleDevice) validateRate(formats strfmt.Registry) error {

	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if err := validate.MinimumInt("Rate", "body", int64(m.Rate), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThrottleDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThrottleDevice) UnmarshalBinary(b []byte) error {
	var res ThrottleDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
