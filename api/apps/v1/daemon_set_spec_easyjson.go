// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjson143303c7DecodeGithubComKubewardenK8sObjectsApiAppsV1(in *jlexer.Lexer, out *DaemonSetSpec) {
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
		case "updateStrategy":
			if in.IsNull() {
				in.Skip()
				out.UpdateStrategy = nil
			} else {
				if out.UpdateStrategy == nil {
					out.UpdateStrategy = new(DaemonSetUpdateStrategy)
				}
				easyjson143303c7DecodeGithubComKubewardenK8sObjectsApiAppsV11(in, out.UpdateStrategy)
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
func easyjson143303c7EncodeGithubComKubewardenK8sObjectsApiAppsV1(out *jwriter.Writer, in DaemonSetSpec) {
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
	{
		const prefix string = ",\"selector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Selector == nil {
			out.RawString("null")
		} else {
			(*in.Selector).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"template\":"
		out.RawString(prefix)
		if in.Template == nil {
			out.RawString("null")
		} else {
			(*in.Template).MarshalEasyJSON(out)
		}
	}
	if in.UpdateStrategy != nil {
		const prefix string = ",\"updateStrategy\":"
		out.RawString(prefix)
		easyjson143303c7EncodeGithubComKubewardenK8sObjectsApiAppsV11(out, *in.UpdateStrategy)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DaemonSetSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson143303c7EncodeGithubComKubewardenK8sObjectsApiAppsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DaemonSetSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson143303c7EncodeGithubComKubewardenK8sObjectsApiAppsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DaemonSetSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson143303c7DecodeGithubComKubewardenK8sObjectsApiAppsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DaemonSetSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson143303c7DecodeGithubComKubewardenK8sObjectsApiAppsV1(l, v)
}
func easyjson143303c7DecodeGithubComKubewardenK8sObjectsApiAppsV11(in *jlexer.Lexer, out *DaemonSetUpdateStrategy) {
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
				easyjson143303c7DecodeGithubComKubewardenK8sObjectsApiAppsV12(in, out.RollingUpdate)
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
func easyjson143303c7EncodeGithubComKubewardenK8sObjectsApiAppsV11(out *jwriter.Writer, in DaemonSetUpdateStrategy) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RollingUpdate != nil {
		const prefix string = ",\"rollingUpdate\":"
		first = false
		out.RawString(prefix[1:])
		easyjson143303c7EncodeGithubComKubewardenK8sObjectsApiAppsV12(out, *in.RollingUpdate)
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
func easyjson143303c7DecodeGithubComKubewardenK8sObjectsApiAppsV12(in *jlexer.Lexer, out *RollingUpdateDaemonSet) {
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
		case "maxSurge":
			if in.IsNull() {
				in.Skip()
				out.MaxSurge = nil
			} else {
				if out.MaxSurge == nil {
					out.MaxSurge = new(intstr.IntOrString)
				}
				*out.MaxSurge = intstr.IntOrString(in.String())
			}
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
func easyjson143303c7EncodeGithubComKubewardenK8sObjectsApiAppsV12(out *jwriter.Writer, in RollingUpdateDaemonSet) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MaxSurge != nil {
		const prefix string = ",\"maxSurge\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.MaxSurge))
	}
	if in.MaxUnavailable != nil {
		const prefix string = ",\"maxUnavailable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.MaxUnavailable))
	}
	out.RawByte('}')
}
