// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SystemVersion system version
// swagger:model SystemVersion
type SystemVersion struct {

	// Api Version held by daemon
	APIVersion string `json:"ApiVersion,omitempty"`

	// Arch type of underlying hardware
	Arch string `json:"Arch,omitempty"`

	// The time when this binary of daemon is built
	BuildTime string `json:"BuildTime,omitempty"`

	// Commit ID held by the latest commit operation
	GitCommit string `json:"GitCommit,omitempty"`

	// version of Go runtime
	GoVersion string `json:"GoVersion,omitempty"`

	// Operating system kernel version
	KernelVersion string `json:"KernelVersion,omitempty"`

	// Operating system type of underlying system
	Os string `json:"Os,omitempty"`

	// version of Pouch Daemon
	Version string `json:"Version,omitempty"`
}

// Validate validates this system version
func (m *SystemVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemVersion) UnmarshalBinary(b []byte) error {
	var res SystemVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
