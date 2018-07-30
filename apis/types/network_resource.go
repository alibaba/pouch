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

// NetworkResource NetworkResource is the body of the "get network" http response message
// swagger:model NetworkResource
type NetworkResource struct {

	// Containers contains endpoints belonging to the network
	Containers interface{} `json:"Containers,omitempty"`

	// Driver is the Driver name used to create the network (e.g. `bridge`, `overlay`)
	Driver string `json:"Driver,omitempty"`

	// EnableIPv6 represents whether to enable IPv6
	EnableIPV6 bool `json:"EnableIPv6,omitempty"`

	// IP a m
	IPAM *IPAM `json:"IPAM,omitempty"`

	// ID uniquely identifies a network on a single machine
	ID string `json:"Id,omitempty"`

	// index configs
	IndexConfigs map[string]EndpointResource `json:"IndexConfigs,omitempty"`

	// Internal represents if the network is used internal only
	Internal bool `json:"Internal,omitempty"`

	// Labels holds metadata specific to the network being created
	Labels map[string]string `json:"Labels,omitempty"`

	// Name is the requested name of the network
	Name string `json:"Name,omitempty"`

	// Options holds the network specific options to use for when creating the network
	Options map[string]string `json:"Options,omitempty"`

	// Scope describes the level at which the network exists (e.g. `global` for cluster-wide or `local` for machine level)
	Scope string `json:"Scope,omitempty"`
}

// Validate validates this network resource
func (m *NetworkResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkResource) validateIPAM(formats strfmt.Registry) error {

	if swag.IsZero(m.IPAM) { // not required
		return nil
	}

	if m.IPAM != nil {
		if err := m.IPAM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IPAM")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkResource) validateIndexConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.IndexConfigs) { // not required
		return nil
	}

	for k := range m.IndexConfigs {

		if err := validate.Required("IndexConfigs"+"."+k, "body", m.IndexConfigs[k]); err != nil {
			return err
		}
		if val, ok := m.IndexConfigs[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkResource) UnmarshalBinary(b []byte) error {
	var res NetworkResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
