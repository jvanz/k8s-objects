// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

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

func easyjson13d54049DecodeGithubComKubewardenK8sObjectsApiNodeV1alpha1(in *jlexer.Lexer, out *RuntimeClassSpec) {
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
		case "overhead":
			if in.IsNull() {
				in.Skip()
				out.Overhead = nil
			} else {
				if out.Overhead == nil {
					out.Overhead = new(Overhead)
				}
				(*out.Overhead).UnmarshalEasyJSON(in)
			}
		case "runtimeHandler":
			if in.IsNull() {
				in.Skip()
				out.RuntimeHandler = nil
			} else {
				if out.RuntimeHandler == nil {
					out.RuntimeHandler = new(string)
				}
				*out.RuntimeHandler = string(in.String())
			}
		case "scheduling":
			if in.IsNull() {
				in.Skip()
				out.Scheduling = nil
			} else {
				if out.Scheduling == nil {
					out.Scheduling = new(Scheduling)
				}
				easyjson13d54049DecodeGithubComKubewardenK8sObjectsApiNodeV1alpha11(in, out.Scheduling)
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
func easyjson13d54049EncodeGithubComKubewardenK8sObjectsApiNodeV1alpha1(out *jwriter.Writer, in RuntimeClassSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Overhead != nil {
		const prefix string = ",\"overhead\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Overhead).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"runtimeHandler\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.RuntimeHandler == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.RuntimeHandler))
		}
	}
	if in.Scheduling != nil {
		const prefix string = ",\"scheduling\":"
		out.RawString(prefix)
		easyjson13d54049EncodeGithubComKubewardenK8sObjectsApiNodeV1alpha11(out, *in.Scheduling)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RuntimeClassSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson13d54049EncodeGithubComKubewardenK8sObjectsApiNodeV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RuntimeClassSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson13d54049EncodeGithubComKubewardenK8sObjectsApiNodeV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RuntimeClassSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson13d54049DecodeGithubComKubewardenK8sObjectsApiNodeV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RuntimeClassSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson13d54049DecodeGithubComKubewardenK8sObjectsApiNodeV1alpha1(l, v)
}
func easyjson13d54049DecodeGithubComKubewardenK8sObjectsApiNodeV1alpha11(in *jlexer.Lexer, out *Scheduling) {
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
						out.Tolerations = make([]_v1.Toleration, 0, 0)
					} else {
						out.Tolerations = []_v1.Toleration{}
					}
				} else {
					out.Tolerations = (out.Tolerations)[:0]
				}
				for !in.IsDelim(']') {
					var v2 _v1.Toleration
					(v2).UnmarshalEasyJSON(in)
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
func easyjson13d54049EncodeGithubComKubewardenK8sObjectsApiNodeV1alpha11(out *jwriter.Writer, in Scheduling) {
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
	{
		const prefix string = ",\"tolerations\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Tolerations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Tolerations {
				if v4 > 0 {
					out.RawByte(',')
				}
				(v5).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
