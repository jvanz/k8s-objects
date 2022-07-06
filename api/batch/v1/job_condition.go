// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// JobCondition JobCondition describes current state of a job.
//
// swagger:model JobCondition
type JobCondition struct {

	// Last time the condition was checked.
	LastProbeTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastProbeTime,omitempty"`

	// Last time the condition transit from one status to another.
	LastTransitionTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastTransitionTime,omitempty"`

	// Human readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// (brief) reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	// Required: true
	Status *string `json:"status"`

	// Type of job condition, Complete or Failed.
	//
	// Possible enum values:
	//  - `"Complete"` means the job has completed its execution.
	//  - `"Failed"` means the job has failed its execution.
	//  - `"Suspended"` means the job has been suspended.
	// Required: true
	// Enum: [Complete Failed Suspended]
	Type *string `json:"type"`
}
