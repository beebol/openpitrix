// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteRepoResponse openpitrix delete repo response
// swagger:model openpitrixDeleteRepoResponse
type OpenpitrixDeleteRepoResponse struct {

	// repo
	Repo *OpenpitrixRepo `json:"repo,omitempty"`
}

// Validate validates this openpitrix delete repo response
func (m *OpenpitrixDeleteRepoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteRepoResponse) validateRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.Repo) { // not required
		return nil
	}

	if m.Repo != nil {

		if err := m.Repo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteRepoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteRepoResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteRepoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}