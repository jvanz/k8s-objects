// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// PersistentVolumeClaimCondition PersistentVolumeClaimCondition contails details about state of pvc
//
// swagger:model PersistentVolumeClaimCondition
type PersistentVolumeClaimCondition struct {

	// Last time we probed the condition.
	LastProbeTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastProbeTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	LastTransitionTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// Unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports "ResizeStarted" that means the underlying persistent volume is being resized.
	Reason string `json:"reason,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	//
	//
	//
	// Possible enum values:
	//  - `"FileSystemResizePending"` - controller resize is finished and a file system resize is pending on node
	//  - `"Resizing"` - a user trigger resize of pvc has been started
	// Required: true
	// Enum: [FileSystemResizePending Resizing]
	Type *string `json:"type"`
}
