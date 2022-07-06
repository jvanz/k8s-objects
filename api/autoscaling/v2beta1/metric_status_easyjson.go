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

func easyjsonA79c8f1fDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(in *jlexer.Lexer, out *MetricStatus) {
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
		case "containerResource":
			if in.IsNull() {
				in.Skip()
				out.ContainerResource = nil
			} else {
				if out.ContainerResource == nil {
					out.ContainerResource = new(ContainerResourceMetricStatus)
				}
				(*out.ContainerResource).UnmarshalEasyJSON(in)
			}
		case "external":
			if in.IsNull() {
				in.Skip()
				out.External = nil
			} else {
				if out.External == nil {
					out.External = new(ExternalMetricStatus)
				}
				(*out.External).UnmarshalEasyJSON(in)
			}
		case "object":
			if in.IsNull() {
				in.Skip()
				out.Object = nil
			} else {
				if out.Object == nil {
					out.Object = new(ObjectMetricStatus)
				}
				easyjsonA79c8f1fDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(in, out.Object)
			}
		case "pods":
			if in.IsNull() {
				in.Skip()
				out.Pods = nil
			} else {
				if out.Pods == nil {
					out.Pods = new(PodsMetricStatus)
				}
				easyjsonA79c8f1fDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(in, out.Pods)
			}
		case "resource":
			if in.IsNull() {
				in.Skip()
				out.Resource = nil
			} else {
				if out.Resource == nil {
					out.Resource = new(ResourceMetricStatus)
				}
				easyjsonA79c8f1fDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(in, out.Resource)
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
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
func easyjsonA79c8f1fEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(out *jwriter.Writer, in MetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ContainerResource != nil {
		const prefix string = ",\"containerResource\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ContainerResource).MarshalEasyJSON(out)
	}
	if in.External != nil {
		const prefix string = ",\"external\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.External).MarshalEasyJSON(out)
	}
	if in.Object != nil {
		const prefix string = ",\"object\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonA79c8f1fEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(out, *in.Object)
	}
	if in.Pods != nil {
		const prefix string = ",\"pods\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonA79c8f1fEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(out, *in.Pods)
	}
	if in.Resource != nil {
		const prefix string = ",\"resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonA79c8f1fEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(out, *in.Resource)
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MetricStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA79c8f1fEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MetricStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA79c8f1fEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MetricStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA79c8f1fDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MetricStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA79c8f1fDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(l, v)
}
func easyjsonA79c8f1fDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(in *jlexer.Lexer, out *ResourceMetricStatus) {
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
		case "currentAverageUtilization":
			out.CurrentAverageUtilization = int32(in.Int32())
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
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
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
func easyjsonA79c8f1fEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(out *jwriter.Writer, in ResourceMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CurrentAverageUtilization != 0 {
		const prefix string = ",\"currentAverageUtilization\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.CurrentAverageUtilization))
	}
	{
		const prefix string = ",\"currentAverageValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.CurrentAverageValue == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.CurrentAverageValue))
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}
func easyjsonA79c8f1fDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(in *jlexer.Lexer, out *PodsMetricStatus) {
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
		case "selector":
			if in.IsNull() {
				in.Skip()
				out.Selector = nil
			} else {
				if out.Selector == nil {
					out.Selector = new(_v1.LabelSelector)
				}
				(*out.Selector).UnmarshalEasyJSON(in)
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
func easyjsonA79c8f1fEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(out *jwriter.Writer, in PodsMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"currentAverageValue\":"
		out.RawString(prefix[1:])
		if in.CurrentAverageValue == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.CurrentAverageValue))
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
	if in.Selector != nil {
		const prefix string = ",\"selector\":"
		out.RawString(prefix)
		(*in.Selector).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjsonA79c8f1fDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(in *jlexer.Lexer, out *ObjectMetricStatus) {
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
		case "averageValue":
			if in.IsNull() {
				in.Skip()
				out.AverageValue = nil
			} else {
				if out.AverageValue == nil {
					out.AverageValue = new(resource.Quantity)
				}
				*out.AverageValue = resource.Quantity(in.String())
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
		case "selector":
			if in.IsNull() {
				in.Skip()
				out.Selector = nil
			} else {
				if out.Selector == nil {
					out.Selector = new(_v1.LabelSelector)
				}
				(*out.Selector).UnmarshalEasyJSON(in)
			}
		case "target":
			if in.IsNull() {
				in.Skip()
				out.Target = nil
			} else {
				if out.Target == nil {
					out.Target = new(CrossVersionObjectReference)
				}
				(*out.Target).UnmarshalEasyJSON(in)
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
func easyjsonA79c8f1fEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(out *jwriter.Writer, in ObjectMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AverageValue != nil {
		const prefix string = ",\"averageValue\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.AverageValue))
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
	if in.Selector != nil {
		const prefix string = ",\"selector\":"
		out.RawString(prefix)
		(*in.Selector).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"target\":"
		out.RawString(prefix)
		if in.Target == nil {
			out.RawString("null")
		} else {
			(*in.Target).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}
