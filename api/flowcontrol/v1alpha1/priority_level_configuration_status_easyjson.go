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

func easyjson211b8a67DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(in *jlexer.Lexer, out *PriorityLevelConfigurationStatus) {
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
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*PriorityLevelConfigurationCondition, 0, 8)
					} else {
						out.Conditions = []*PriorityLevelConfigurationCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *PriorityLevelConfigurationCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(PriorityLevelConfigurationCondition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson211b8a67EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(out *jwriter.Writer, in PriorityLevelConfigurationStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.Conditions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PriorityLevelConfigurationStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson211b8a67EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PriorityLevelConfigurationStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson211b8a67EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PriorityLevelConfigurationStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson211b8a67DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PriorityLevelConfigurationStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson211b8a67DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(l, v)
}
