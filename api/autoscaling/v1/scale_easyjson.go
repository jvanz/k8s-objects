// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
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

func easyjsonF06ff468DecodeGithubComKubewardenK8sObjectsApiAutoscalingV1(in *jlexer.Lexer, out *Scale) {
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
			(out.Metadata).UnmarshalEasyJSON(in)
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(ScaleSpec)
				}
				easyjsonF06ff468DecodeGithubComKubewardenK8sObjectsApiAutoscalingV11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(ScaleStatus)
				}
				easyjsonF06ff468DecodeGithubComKubewardenK8sObjectsApiAutoscalingV12(in, out.Status)
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
func easyjsonF06ff468EncodeGithubComKubewardenK8sObjectsApiAutoscalingV1(out *jwriter.Writer, in Scale) {
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
	if true {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Metadata).MarshalEasyJSON(out)
	}
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonF06ff468EncodeGithubComKubewardenK8sObjectsApiAutoscalingV11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonF06ff468EncodeGithubComKubewardenK8sObjectsApiAutoscalingV12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Scale) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF06ff468EncodeGithubComKubewardenK8sObjectsApiAutoscalingV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Scale) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF06ff468EncodeGithubComKubewardenK8sObjectsApiAutoscalingV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Scale) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF06ff468DecodeGithubComKubewardenK8sObjectsApiAutoscalingV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Scale) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF06ff468DecodeGithubComKubewardenK8sObjectsApiAutoscalingV1(l, v)
}
func easyjsonF06ff468DecodeGithubComKubewardenK8sObjectsApiAutoscalingV12(in *jlexer.Lexer, out *ScaleStatus) {
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
		case "replicas":
			if in.IsNull() {
				in.Skip()
				out.Replicas = nil
			} else {
				if out.Replicas == nil {
					out.Replicas = new(int32)
				}
				*out.Replicas = int32(in.Int32())
			}
		case "selector":
			out.Selector = string(in.String())
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
func easyjsonF06ff468EncodeGithubComKubewardenK8sObjectsApiAutoscalingV12(out *jwriter.Writer, in ScaleStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"replicas\":"
		out.RawString(prefix[1:])
		if in.Replicas == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Replicas))
		}
	}
	if in.Selector != "" {
		const prefix string = ",\"selector\":"
		out.RawString(prefix)
		out.String(string(in.Selector))
	}
	out.RawByte('}')
}
func easyjsonF06ff468DecodeGithubComKubewardenK8sObjectsApiAutoscalingV11(in *jlexer.Lexer, out *ScaleSpec) {
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
		case "replicas":
			out.Replicas = int32(in.Int32())
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
func easyjsonF06ff468EncodeGithubComKubewardenK8sObjectsApiAutoscalingV11(out *jwriter.Writer, in ScaleSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Replicas != 0 {
		const prefix string = ",\"replicas\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Replicas))
	}
	out.RawByte('}')
}
