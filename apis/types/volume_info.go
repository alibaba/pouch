// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumeInfo Volume represents the configuration of a volume for the container.
// swagger:model VolumeInfo
type VolumeInfo struct {

	// Date/Time the volume was created.
	CreatedAt string `json:"CreatedAt,omitempty"`

	// Driver is the Driver name used to create the volume.
	Driver string `json:"Driver,omitempty"`

	// Labels is metadata specific to the volume.
	Labels map[string]string `json:"Labels,omitempty"`

	// Mountpoint is the location on disk of the volume.
	Mountpoint string `json:"Mountpoint,omitempty"`

	// Name is the name of the volume.
	Name string `json:"Name,omitempty"`

	// Scope describes the level at which the volume exists
	// (e.g. `global` for cluster-wide or `local` for machine level)
	//
	Scope string `json:"Scope,omitempty"`

	// Status provides low-level status information about the volume.
	Status map[string]interface{} `json:"Status,omitempty"`
}

// Validate validates this volume info
func (m *VolumeInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// additional properties value enum
var volumeInfoStatusValueEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`[{}]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeInfoStatusValueEnum = append(volumeInfoStatusValueEnum, v)
	}
}

func (m *VolumeInfo) validateStatusValueEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, volumeInfoStatusValueEnum); err != nil {
		return err
	}
	return nil
}

func (m *VolumeInfo) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	for k := range m.Status {

		if err := m.validateStatusValueEnum("Status"+"."+k, "body", m.Status[k]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeInfo) UnmarshalBinary(b []byte) error {
	var res VolumeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
