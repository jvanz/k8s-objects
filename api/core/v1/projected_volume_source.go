// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ProjectedVolumeSource Represents a projected volume source
//
// swagger:model ProjectedVolumeSource
type ProjectedVolumeSource struct {

	// Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode int32 `json:"defaultMode,omitempty"`

	// list of volume projections
	Sources []*VolumeProjection `json:"sources,omitempty"`
}
