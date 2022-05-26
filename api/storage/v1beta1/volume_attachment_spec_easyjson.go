// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

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

func easyjsonB50f7d14DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(in *jlexer.Lexer, out *VolumeAttachmentSpec) {
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
		case "attacher":
			if in.IsNull() {
				in.Skip()
				out.Attacher = nil
			} else {
				if out.Attacher == nil {
					out.Attacher = new(string)
				}
				*out.Attacher = string(in.String())
			}
		case "nodeName":
			if in.IsNull() {
				in.Skip()
				out.NodeName = nil
			} else {
				if out.NodeName == nil {
					out.NodeName = new(string)
				}
				*out.NodeName = string(in.String())
			}
		case "source":
			if in.IsNull() {
				in.Skip()
				out.Source = nil
			} else {
				if out.Source == nil {
					out.Source = new(VolumeAttachmentSource)
				}
				(*out.Source).UnmarshalEasyJSON(in)
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
func easyjsonB50f7d14EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(out *jwriter.Writer, in VolumeAttachmentSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"attacher\":"
		out.RawString(prefix[1:])
		if in.Attacher == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Attacher))
		}
	}
	{
		const prefix string = ",\"nodeName\":"
		out.RawString(prefix)
		if in.NodeName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.NodeName))
		}
	}
	{
		const prefix string = ",\"source\":"
		out.RawString(prefix)
		if in.Source == nil {
			out.RawString("null")
		} else {
			(*in.Source).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VolumeAttachmentSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB50f7d14EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VolumeAttachmentSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB50f7d14EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VolumeAttachmentSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB50f7d14DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VolumeAttachmentSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB50f7d14DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(l, v)
}
