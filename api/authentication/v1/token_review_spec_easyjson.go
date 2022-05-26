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

func easyjson8cb08f22DecodeGithubComKubewardenK8sObjectsApiAuthenticationV1(in *jlexer.Lexer, out *TokenReviewSpec) {
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
		case "audiences":
			if in.IsNull() {
				in.Skip()
				out.Audiences = nil
			} else {
				in.Delim('[')
				if out.Audiences == nil {
					if !in.IsDelim(']') {
						out.Audiences = make([]string, 0, 4)
					} else {
						out.Audiences = []string{}
					}
				} else {
					out.Audiences = (out.Audiences)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Audiences = append(out.Audiences, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "token":
			out.Token = string(in.String())
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
func easyjson8cb08f22EncodeGithubComKubewardenK8sObjectsApiAuthenticationV1(out *jwriter.Writer, in TokenReviewSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"audiences\":"
		out.RawString(prefix[1:])
		if in.Audiences == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Audiences {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.Token != "" {
		const prefix string = ",\"token\":"
		out.RawString(prefix)
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TokenReviewSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8cb08f22EncodeGithubComKubewardenK8sObjectsApiAuthenticationV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TokenReviewSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8cb08f22EncodeGithubComKubewardenK8sObjectsApiAuthenticationV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TokenReviewSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8cb08f22DecodeGithubComKubewardenK8sObjectsApiAuthenticationV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TokenReviewSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8cb08f22DecodeGithubComKubewardenK8sObjectsApiAuthenticationV1(l, v)
}
