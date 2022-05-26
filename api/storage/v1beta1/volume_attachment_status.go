// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// VolumeAttachmentStatus VolumeAttachmentStatus is the status of a VolumeAttachment request.
//
// swagger:model VolumeAttachmentStatus
type VolumeAttachmentStatus struct {

	// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachError *VolumeError `json:"attachError,omitempty"`

	// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	// Required: true
	Attached *bool `json:"attached"`

	// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachmentMetadata map[string]string `json:"attachmentMetadata,omitempty"`

	// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
	DetachError *VolumeError `json:"detachError,omitempty"`
}
