// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// StatusDetails StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.
//
// swagger:model StatusDetails
type StatusDetails struct {

	// The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.
	Causes []*StatusCause `json:"causes,omitempty"`

	// The group attribute of the resource associated with the status StatusReason.
	Group string `json:"group,omitempty"`

	// The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).
	Name string `json:"name,omitempty"`

	// If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.
	RetryAfterSeconds int32 `json:"retryAfterSeconds,omitempty"`

	// UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	UID string `json:"uid,omitempty"`
}
