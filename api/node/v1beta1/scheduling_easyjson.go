// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
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

func easyjsonAddc83faDecodeGithubComKubewardenK8sObjectsApiNodeV1beta1(in *jlexer.Lexer, out *Scheduling) {
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
		case "nodeSelector":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.NodeSelector = make(map[string]string)
				} else {
					out.NodeSelector = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.NodeSelector)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "tolerations":
			if in.IsNull() {
				in.Skip()
				out.Tolerations = nil
			} else {
				in.Delim('[')
				if out.Tolerations == nil {
					if !in.IsDelim(']') {
						out.Tolerations = make([]*_v1.Toleration, 0, 8)
					} else {
						out.Tolerations = []*_v1.Toleration{}
					}
				} else {
					out.Tolerations = (out.Tolerations)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *_v1.Toleration
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(_v1.Toleration)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.Tolerations = append(out.Tolerations, v2)
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
func easyjsonAddc83faEncodeGithubComKubewardenK8sObjectsApiNodeV1beta1(out *jwriter.Writer, in Scheduling) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.NodeSelector) != 0 {
		const prefix string = ",\"nodeSelector\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('{')
			v3First := true
			for v3Name, v3Value := range in.NodeSelector {
				if v3First {
					v3First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v3Name))
				out.RawByte(':')
				out.String(string(v3Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.Tolerations) != 0 {
		const prefix string = ",\"tolerations\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v4, v5 := range in.Tolerations {
				if v4 > 0 {
					out.RawByte(',')
				}
				if v5 == nil {
					out.RawString("null")
				} else {
					(*v5).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Scheduling) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAddc83faEncodeGithubComKubewardenK8sObjectsApiNodeV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Scheduling) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAddc83faEncodeGithubComKubewardenK8sObjectsApiNodeV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Scheduling) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAddc83faDecodeGithubComKubewardenK8sObjectsApiNodeV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Scheduling) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAddc83faDecodeGithubComKubewardenK8sObjectsApiNodeV1beta1(l, v)
}
