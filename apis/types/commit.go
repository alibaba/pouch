// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Commit Commit holds the Git-commit (SHA1) that a binary was built from, as
// reported in the version-string of external tools, such as `containerd`,
// or `runC`.
//
// swagger:model Commit
type Commit struct {

	// Commit ID of external tool expected by pouchd as set at build time.
	//
	Expected string `json:"Expected,omitempty"`

	// Actual commit ID of external tool.
	ID string `json:"ID,omitempty"`
}

// Validate validates this commit
func (m *Commit) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Commit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Commit) UnmarshalBinary(b []byte) error {
	var res Commit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
