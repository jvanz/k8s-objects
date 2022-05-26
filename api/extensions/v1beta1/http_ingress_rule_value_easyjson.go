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

func easyjson658ab631DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *HTTPIngressRuleValue) {
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
		case "paths":
			if in.IsNull() {
				in.Skip()
				out.Paths = nil
			} else {
				in.Delim('[')
				if out.Paths == nil {
					if !in.IsDelim(']') {
						out.Paths = make([]*HTTPIngressPath, 0, 8)
					} else {
						out.Paths = []*HTTPIngressPath{}
					}
				} else {
					out.Paths = (out.Paths)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *HTTPIngressPath
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(HTTPIngressPath)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Paths = append(out.Paths, v1)
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
func easyjson658ab631EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in HTTPIngressRuleValue) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"paths\":"
		out.RawString(prefix[1:])
		if in.Paths == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Paths {
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
func (v HTTPIngressRuleValue) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson658ab631EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HTTPIngressRuleValue) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson658ab631EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HTTPIngressRuleValue) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson658ab631DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HTTPIngressRuleValue) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson658ab631DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}
