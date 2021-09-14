// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorModel An error description
//
// swagger:model ErrorModel
type ErrorModel struct {

	// The HTTP-error number
	Code int64 `json:"code,omitempty"`

	// A textual description of the error occured.
	Message string `json:"message,omitempty"`
}

// Validate validates this error model
func (m *ErrorModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error model based on context it is used
func (m *ErrorModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorModel) UnmarshalBinary(b []byte) error {
	var res ErrorModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
