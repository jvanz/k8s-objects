// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_core_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
)

// IngressClassSpec IngressClassSpec provides information about the class of an Ingress.
//
// swagger:model IngressClassSpec
type IngressClassSpec struct {

	// Controller refers to the name of the controller that should handle this class. This allows for different "flavors" that are controlled by the same controller. For example, you may have different Parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. "acme.io/ingress-controller". This field is immutable.
	Controller string `json:"controller,omitempty"`

	// Parameters is a link to a custom resource containing additional configuration for the controller. This is optional if the controller does not require extra parameters.
	Parameters api_core_v1.TypedLocalObjectReference `json:"parameters,omitempty"`
}
