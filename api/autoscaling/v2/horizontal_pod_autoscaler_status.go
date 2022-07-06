// Code generated by go-swagger; DO NOT EDIT.

package v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// HorizontalPodAutoscalerStatus HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
//
// swagger:model HorizontalPodAutoscalerStatus
type HorizontalPodAutoscalerStatus struct {

	// conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
	Conditions []*HorizontalPodAutoscalerCondition `json:"conditions,omitempty"`

	// currentMetrics is the last read state of the metrics used by this autoscaler.
	CurrentMetrics []*MetricStatus `json:"currentMetrics,omitempty"`

	// currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
	CurrentReplicas int32 `json:"currentReplicas,omitempty"`

	// desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
	// Required: true
	DesiredReplicas *int32 `json:"desiredReplicas"`

	// lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
	LastScaleTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastScaleTime,omitempty"`

	// observedGeneration is the most recent generation observed by this autoscaler.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}
