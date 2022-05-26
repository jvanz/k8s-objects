// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// EndpointHints EndpointHints provides hints describing how an endpoint should be consumed.
//
// swagger:model EndpointHints
type EndpointHints struct {

	// forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.
	ForZones []*ForZone `json:"forZones"`
}
