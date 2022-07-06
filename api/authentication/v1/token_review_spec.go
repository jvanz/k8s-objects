// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// TokenReviewSpec TokenReviewSpec is a description of the token authentication request.
//
// swagger:model TokenReviewSpec
type TokenReviewSpec struct {

	// Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
	Audiences []string `json:"audiences,omitempty"`

	// Token is the opaque bearer token.
	Token string `json:"token,omitempty"`
}
