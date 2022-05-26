// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2alpha1

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

func easyjson761e00a6DecodeGithubComKubewardenK8sObjectsApiBatchV2alpha1(in *jlexer.Lexer, out *JobTemplateSpec) {
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
		case "metadata":
			(out.Metadata).UnmarshalEasyJSON(in)
		case "spec":
			(out.Spec).UnmarshalEasyJSON(in)
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
func easyjson761e00a6EncodeGithubComKubewardenK8sObjectsApiBatchV2alpha1(out *jwriter.Writer, in JobTemplateSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"metadata\":"
		first = false
		out.RawString(prefix[1:])
		(in.Metadata).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Spec).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JobTemplateSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson761e00a6EncodeGithubComKubewardenK8sObjectsApiBatchV2alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JobTemplateSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson761e00a6EncodeGithubComKubewardenK8sObjectsApiBatchV2alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JobTemplateSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson761e00a6DecodeGithubComKubewardenK8sObjectsApiBatchV2alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JobTemplateSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson761e00a6DecodeGithubComKubewardenK8sObjectsApiBatchV2alpha1(l, v)
}
