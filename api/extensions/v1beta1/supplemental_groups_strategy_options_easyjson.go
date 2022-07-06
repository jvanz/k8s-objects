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

func easyjsonBb83f54DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *SupplementalGroupsStrategyOptions) {
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
		case "ranges":
			if in.IsNull() {
				in.Skip()
				out.Ranges = nil
			} else {
				in.Delim('[')
				if out.Ranges == nil {
					if !in.IsDelim(']') {
						out.Ranges = make([]*IDRange, 0, 8)
					} else {
						out.Ranges = []*IDRange{}
					}
				} else {
					out.Ranges = (out.Ranges)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *IDRange
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(IDRange)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Ranges = append(out.Ranges, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "rule":
			out.Rule = string(in.String())
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
func easyjsonBb83f54EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in SupplementalGroupsStrategyOptions) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Ranges) != 0 {
		const prefix string = ",\"ranges\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.Ranges {
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
	if in.Rule != "" {
		const prefix string = ",\"rule\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Rule))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SupplementalGroupsStrategyOptions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBb83f54EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SupplementalGroupsStrategyOptions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBb83f54EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SupplementalGroupsStrategyOptions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBb83f54DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SupplementalGroupsStrategyOptions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBb83f54DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}
