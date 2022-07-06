// Code generated by go-swagger; DO NOT EDIT.

package v1beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// ReplicaSetCondition ReplicaSetCondition describes the state of a replica set at a certain point.
//
// swagger:model ReplicaSetCondition
type ReplicaSetCondition struct {

	// The last time the condition transitioned from one status to another.
	LastTransitionTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastTransitionTime,omitempty"`

	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`

	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	// Required: true
	Status *string `json:"status"`

	// Type of replica set condition.
	// Required: true
	Type *string `json:"type"`
}
