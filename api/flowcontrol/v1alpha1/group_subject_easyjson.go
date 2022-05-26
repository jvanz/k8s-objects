// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

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

func easyjsonA80039eaDecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(in *jlexer.Lexer, out *GroupSubject) {
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
func easyjsonA80039eaEncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(out *jwriter.Writer, in GroupSubject) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GroupSubject) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA80039eaEncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GroupSubject) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA80039eaEncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GroupSubject) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA80039eaDecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GroupSubject) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA80039eaDecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(l, v)
}
