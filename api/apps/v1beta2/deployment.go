// Code generated by go-swagger; DO NOT EDIT.

package v1beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// Deployment DEPRECATED - This group version of Deployment is deprecated by apps/v1/Deployment. See the release notes for more information. Deployment enables declarative updates for Pods and ReplicaSets.
//
// swagger:model Deployment
type Deployment struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// Standard object metadata.
	Metadata *apimachinery_pkg_apis_meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of the Deployment.
	Spec *DeploymentSpec `json:"spec,omitempty"`

	// Most recently observed status of the Deployment.
	Status *DeploymentStatus `json:"status,omitempty"`
}
