// Code generated by go-swagger; DO NOT EDIT.

package v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Policy Policy defines the configuration of how audit events are logged
//
// swagger:model Policy
type Policy struct {

	// The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required
	// Required: true
	Level *string `json:"level"`

	// Stages is a list of stages for which events are created.
	Stages []string `json:"stages,omitempty"`
}
