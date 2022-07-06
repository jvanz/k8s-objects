// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2

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

func easyjson320b3e1DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(in *jlexer.Lexer, out *HPAScalingRules) {
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
		case "policies":
			if in.IsNull() {
				in.Skip()
				out.Policies = nil
			} else {
				in.Delim('[')
				if out.Policies == nil {
					if !in.IsDelim(']') {
						out.Policies = make([]*HPAScalingPolicy, 0, 8)
					} else {
						out.Policies = []*HPAScalingPolicy{}
					}
				} else {
					out.Policies = (out.Policies)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *HPAScalingPolicy
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(HPAScalingPolicy)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Policies = append(out.Policies, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "selectPolicy":
			out.SelectPolicy = string(in.String())
		case "stabilizationWindowSeconds":
			out.StabilizationWindowSeconds = int32(in.Int32())
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
func easyjson320b3e1EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(out *jwriter.Writer, in HPAScalingRules) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Policies) != 0 {
		const prefix string = ",\"policies\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.Policies {
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
	if in.SelectPolicy != "" {
		const prefix string = ",\"selectPolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SelectPolicy))
	}
	if in.StabilizationWindowSeconds != 0 {
		const prefix string = ",\"stabilizationWindowSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.StabilizationWindowSeconds))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HPAScalingRules) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson320b3e1EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HPAScalingRules) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson320b3e1EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HPAScalingRules) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson320b3e1DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HPAScalingRules) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson320b3e1DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(l, v)
}
