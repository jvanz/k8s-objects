// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// CustomResourceDefinitionCondition CustomResourceDefinitionCondition contains details for the current condition of this pod.
//
// swagger:model CustomResourceDefinitionCondition
type CustomResourceDefinitionCondition struct {

	// Last time the condition transitioned from one status to another.
	LastTransitionTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status is the status of the condition. Can be True, False, Unknown.
	// Required: true
	Status *string `json:"status"`

	// Type is the type of the condition.
	// Required: true
	Type *string `json:"type"`
}
