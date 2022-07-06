// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2beta1

import (
	json "encoding/json"
	resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonA168a7c1DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(in *jlexer.Lexer, out *ExternalMetricStatus) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "currentAverageValue":
			if in.IsNull() {
				in.Skip()
				out.CurrentAverageValue = nil
			} else {
				if out.CurrentAverageValue == nil {
					out.CurrentAverageValue = new(resource.Quantity)
				}
				*out.CurrentAverageValue = resource.Quantity(in.String())
			}
		case "currentValue":
			if in.IsNull() {
				in.Skip()
				out.CurrentValue = nil
			} else {
				if out.CurrentValue == nil {
					out.CurrentValue = new(resource.Quantity)
				}
				*out.CurrentValue = resource.Quantity(in.String())
			}
		case "metricName":
			if in.IsNull() {
				in.Skip()
				out.MetricName = nil
			} else {
				if out.MetricName == nil {
					out.MetricName = new(string)
				}
				*out.MetricName = string(in.String())
			}
		case "metricSelector":
			if in.IsNull() {
				in.Skip()
				out.MetricSelector = nil
			} else {
				if out.MetricSelector == nil {
					out.MetricSelector = new(_v1.LabelSelector)
				}
				(*out.MetricSelector).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA168a7c1EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(out *jwriter.Writer, in ExternalMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CurrentAverageValue != nil {
		const prefix string = ",\"currentAverageValue\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.CurrentAverageValue))
	}
	{
		const prefix string = ",\"currentValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.CurrentValue == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.CurrentValue))
		}
	}
	{
		const prefix string = ",\"metricName\":"
		out.RawString(prefix)
		if in.MetricName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.MetricName))
		}
	}
	if in.MetricSelector != nil {
		const prefix string = ",\"metricSelector\":"
		out.RawString(prefix)
		(*in.MetricSelector).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ExternalMetricStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA168a7c1EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ExternalMetricStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA168a7c1EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ExternalMetricStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA168a7c1DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ExternalMetricStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA168a7c1DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(l, v)
}
