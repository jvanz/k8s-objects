// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// VolumeNodeResources VolumeNodeResources is a set of resource limits for scheduling of volumes.
//
// swagger:model VolumeNodeResources
type VolumeNodeResources struct {

	// Maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is nil, then the supported number of volumes on this node is unbounded.
	Count int32 `json:"count,omitempty"`
}
