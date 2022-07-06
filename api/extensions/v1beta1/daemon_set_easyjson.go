// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
	_v11 "github.com/kubewarden/k8s-objects/api/core/v1"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
	intstr "github.com/kubewarden/k8s-objects/apimachinery/pkg/util/intstr"
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

func easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *DaemonSet) {
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
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(_v1.ObjectMeta)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
			}
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(DaemonSetSpec)
				}
				easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(DaemonSetStatus)
				}
				easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(in, out.Status)
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
func easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in DaemonSet) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kind))
	}
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Metadata).MarshalEasyJSON(out)
	}
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DaemonSet) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DaemonSet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DaemonSet) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DaemonSet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}
func easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(in *jlexer.Lexer, out *DaemonSetStatus) {
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
		case "collisionCount":
			out.CollisionCount = int32(in.Int32())
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*DaemonSetCondition, 0, 8)
					} else {
						out.Conditions = []*DaemonSetCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *DaemonSetCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(DaemonSetCondition)
						}
						easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(in, v1)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "currentNumberScheduled":
			if in.IsNull() {
				in.Skip()
				out.CurrentNumberScheduled = nil
			} else {
				if out.CurrentNumberScheduled == nil {
					out.CurrentNumberScheduled = new(int32)
				}
				*out.CurrentNumberScheduled = int32(in.Int32())
			}
		case "desiredNumberScheduled":
			if in.IsNull() {
				in.Skip()
				out.DesiredNumberScheduled = nil
			} else {
				if out.DesiredNumberScheduled == nil {
					out.DesiredNumberScheduled = new(int32)
				}
				*out.DesiredNumberScheduled = int32(in.Int32())
			}
		case "numberAvailable":
			out.NumberAvailable = int32(in.Int32())
		case "numberMisscheduled":
			if in.IsNull() {
				in.Skip()
				out.NumberMisscheduled = nil
			} else {
				if out.NumberMisscheduled == nil {
					out.NumberMisscheduled = new(int32)
				}
				*out.NumberMisscheduled = int32(in.Int32())
			}
		case "numberReady":
			if in.IsNull() {
				in.Skip()
				out.NumberReady = nil
			} else {
				if out.NumberReady == nil {
					out.NumberReady = new(int32)
				}
				*out.NumberReady = int32(in.Int32())
			}
		case "numberUnavailable":
			out.NumberUnavailable = int32(in.Int32())
		case "observedGeneration":
			out.ObservedGeneration = int64(in.Int64())
		case "updatedNumberScheduled":
			out.UpdatedNumberScheduled = int32(in.Int32())
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
func easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(out *jwriter.Writer, in DaemonSetStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CollisionCount != 0 {
		const prefix string = ",\"collisionCount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.CollisionCount))
	}
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Conditions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"currentNumberScheduled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.CurrentNumberScheduled == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.CurrentNumberScheduled))
		}
	}
	{
		const prefix string = ",\"desiredNumberScheduled\":"
		out.RawString(prefix)
		if in.DesiredNumberScheduled == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.DesiredNumberScheduled))
		}
	}
	if in.NumberAvailable != 0 {
		const prefix string = ",\"numberAvailable\":"
		out.RawString(prefix)
		out.Int32(int32(in.NumberAvailable))
	}
	{
		const prefix string = ",\"numberMisscheduled\":"
		out.RawString(prefix)
		if in.NumberMisscheduled == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.NumberMisscheduled))
		}
	}
	{
		const prefix string = ",\"numberReady\":"
		out.RawString(prefix)
		if in.NumberReady == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.NumberReady))
		}
	}
	if in.NumberUnavailable != 0 {
		const prefix string = ",\"numberUnavailable\":"
		out.RawString(prefix)
		out.Int32(int32(in.NumberUnavailable))
	}
	if in.ObservedGeneration != 0 {
		const prefix string = ",\"observedGeneration\":"
		out.RawString(prefix)
		out.Int64(int64(in.ObservedGeneration))
	}
	if in.UpdatedNumberScheduled != 0 {
		const prefix string = ",\"updatedNumberScheduled\":"
		out.RawString(prefix)
		out.Int32(int32(in.UpdatedNumberScheduled))
	}
	out.RawByte('}')
}
func easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(in *jlexer.Lexer, out *DaemonSetCondition) {
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
		case "lastTransitionTime":
			if in.IsNull() {
				in.Skip()
				out.LastTransitionTime = nil
			} else {
				if out.LastTransitionTime == nil {
					out.LastTransitionTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastTransitionTime).UnmarshalJSON(data))
				}
			}
		case "message":
			out.Message = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(string)
				}
				*out.Status = string(in.String())
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
func easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(out *jwriter.Writer, in DaemonSetCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastTransitionTime != nil {
		const prefix string = ",\"lastTransitionTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.LastTransitionTime).MarshalJSON())
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Reason))
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Status == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Status))
		}
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}
func easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in *jlexer.Lexer, out *DaemonSetSpec) {
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
		case "minReadySeconds":
			out.MinReadySeconds = int32(in.Int32())
		case "revisionHistoryLimit":
			out.RevisionHistoryLimit = int32(in.Int32())
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
		case "template":
			if in.IsNull() {
				in.Skip()
				out.Template = nil
			} else {
				if out.Template == nil {
					out.Template = new(_v11.PodTemplateSpec)
				}
				(*out.Template).UnmarshalEasyJSON(in)
			}
		case "templateGeneration":
			out.TemplateGeneration = int64(in.Int64())
		case "updateStrategy":
			if in.IsNull() {
				in.Skip()
				out.UpdateStrategy = nil
			} else {
				if out.UpdateStrategy == nil {
					out.UpdateStrategy = new(DaemonSetUpdateStrategy)
				}
				easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta14(in, out.UpdateStrategy)
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
func easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out *jwriter.Writer, in DaemonSetSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MinReadySeconds != 0 {
		const prefix string = ",\"minReadySeconds\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.MinReadySeconds))
	}
	if in.RevisionHistoryLimit != 0 {
		const prefix string = ",\"revisionHistoryLimit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.RevisionHistoryLimit))
	}
	if in.Selector != nil {
		const prefix string = ",\"selector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Selector).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"template\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Template == nil {
			out.RawString("null")
		} else {
			(*in.Template).MarshalEasyJSON(out)
		}
	}
	if in.TemplateGeneration != 0 {
		const prefix string = ",\"templateGeneration\":"
		out.RawString(prefix)
		out.Int64(int64(in.TemplateGeneration))
	}
	if in.UpdateStrategy != nil {
		const prefix string = ",\"updateStrategy\":"
		out.RawString(prefix)
		easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta14(out, *in.UpdateStrategy)
	}
	out.RawByte('}')
}
func easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta14(in *jlexer.Lexer, out *DaemonSetUpdateStrategy) {
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
		case "rollingUpdate":
			if in.IsNull() {
				in.Skip()
				out.RollingUpdate = nil
			} else {
				if out.RollingUpdate == nil {
					out.RollingUpdate = new(RollingUpdateDaemonSet)
				}
				easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta15(in, out.RollingUpdate)
			}
		case "type":
			out.Type = string(in.String())
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
func easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta14(out *jwriter.Writer, in DaemonSetUpdateStrategy) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RollingUpdate != nil {
		const prefix string = ",\"rollingUpdate\":"
		first = false
		out.RawString(prefix[1:])
		easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta15(out, *in.RollingUpdate)
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	out.RawByte('}')
}
func easyjson21f59f41DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta15(in *jlexer.Lexer, out *RollingUpdateDaemonSet) {
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
		case "maxUnavailable":
			if in.IsNull() {
				in.Skip()
				out.MaxUnavailable = nil
			} else {
				if out.MaxUnavailable == nil {
					out.MaxUnavailable = new(intstr.IntOrString)
				}
				*out.MaxUnavailable = intstr.IntOrString(in.String())
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
func easyjson21f59f41EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta15(out *jwriter.Writer, in RollingUpdateDaemonSet) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MaxUnavailable != nil {
		const prefix string = ",\"maxUnavailable\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.MaxUnavailable))
	}
	out.RawByte('}')
}
