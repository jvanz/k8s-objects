// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_core_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
)

// VolumeAttachmentSource VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
//
// swagger:model VolumeAttachmentSource
type VolumeAttachmentSource struct {

	// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
	InlineVolumeSpec api_core_v1.PersistentVolumeSpec `json:"inlineVolumeSpec,omitempty"`

	// Name of the persistent volume to attach.
	PersistentVolumeName string `json:"persistentVolumeName,omitempty"`
}
