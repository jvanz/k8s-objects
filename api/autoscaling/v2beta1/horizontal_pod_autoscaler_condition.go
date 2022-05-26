// Code generated by go-swagger; DO NOT EDIT.

package v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// HorizontalPodAutoscalerCondition HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.
//
// swagger:model HorizontalPodAutoscalerCondition
type HorizontalPodAutoscalerCondition struct {

	// lastTransitionTime is the last time the condition transitioned from one status to another
	LastTransitionTime apimachinery_pkg_apis_meta_v1.Time `json:"lastTransitionTime,omitempty"`

	// message is a human-readable explanation containing details about the transition
	Message string `json:"message,omitempty"`

	// reason is the reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// status is the status of the condition (True, False, Unknown)
	// Required: true
	Status *string `json:"status"`

	// type describes the current condition
	// Required: true
	Type *string `json:"type"`
}
