// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// FSGroupStrategyOptions FSGroupStrategyOptions defines the strategy type and options used to create the strategy.
//
// swagger:model FSGroupStrategyOptions
type FSGroupStrategyOptions struct {

	// ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end. Required for MustRunAs.
	Ranges []*IDRange `json:"ranges"`

	// rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
	Rule string `json:"rule,omitempty"`
}
