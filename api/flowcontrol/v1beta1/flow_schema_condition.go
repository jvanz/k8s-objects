// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// FlowSchemaCondition FlowSchemaCondition describes conditions for a FlowSchema.
//
// swagger:model FlowSchemaCondition
type FlowSchemaCondition struct {

	// `lastTransitionTime` is the last time the condition transitioned from one status to another.
	LastTransitionTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastTransitionTime,omitempty"`

	// `message` is a human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// `status` is the status of the condition. Can be True, False, Unknown. Required.
	Status string `json:"status,omitempty"`

	// `type` is the type of the condition. Required.
	Type string `json:"type,omitempty"`
}
