// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
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

func easyjsonEfd2f10eDecodeGithubComKubewardenK8sObjectsApiAuthenticationV1(in *jlexer.Lexer, out *TokenReview) {
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
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(_v1.ObjectMeta)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
			}
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(TokenReviewSpec)
				}
				easyjsonEfd2f10eDecodeGithubComKubewardenK8sObjectsApiAuthenticationV11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(TokenReviewStatus)
				}
				easyjsonEfd2f10eDecodeGithubComKubewardenK8sObjectsApiAuthenticationV12(in, out.Status)
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
func easyjsonEfd2f10eEncodeGithubComKubewardenK8sObjectsApiAuthenticationV1(out *jwriter.Writer, in TokenReview) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kind))
	}
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Metadata).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Spec == nil {
			out.RawString("null")
		} else {
			easyjsonEfd2f10eEncodeGithubComKubewardenK8sObjectsApiAuthenticationV11(out, *in.Spec)
		}
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		easyjsonEfd2f10eEncodeGithubComKubewardenK8sObjectsApiAuthenticationV12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TokenReview) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEfd2f10eEncodeGithubComKubewardenK8sObjectsApiAuthenticationV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TokenReview) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEfd2f10eEncodeGithubComKubewardenK8sObjectsApiAuthenticationV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TokenReview) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEfd2f10eDecodeGithubComKubewardenK8sObjectsApiAuthenticationV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TokenReview) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEfd2f10eDecodeGithubComKubewardenK8sObjectsApiAuthenticationV1(l, v)
}
func easyjsonEfd2f10eDecodeGithubComKubewardenK8sObjectsApiAuthenticationV12(in *jlexer.Lexer, out *TokenReviewStatus) {
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
		case "authenticated":
			out.Authenticated = bool(in.Bool())
		case "error":
			out.Error = string(in.String())
		case "user":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(UserInfo)
				}
				easyjsonEfd2f10eDecodeGithubComKubewardenK8sObjectsApiAuthenticationV13(in, out.User)
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
func easyjsonEfd2f10eEncodeGithubComKubewardenK8sObjectsApiAuthenticationV12(out *jwriter.Writer, in TokenReviewStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Audiences) != 0 {
		const prefix string = ",\"audiences\":"
		first = false
		out.RawString(prefix[1:])
		{
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
	if in.Authenticated {
		const prefix string = ",\"authenticated\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Authenticated))
	}
	if in.Error != "" {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error))
	}
	if in.User != nil {
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonEfd2f10eEncodeGithubComKubewardenK8sObjectsApiAuthenticationV13(out, *in.User)
	}
	out.RawByte('}')
}
func easyjsonEfd2f10eDecodeGithubComKubewardenK8sObjectsApiAuthenticationV13(in *jlexer.Lexer, out *UserInfo) {
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
		case "extra":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Extra = make(map[string][]string)
				} else {
					out.Extra = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 []string
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						in.Delim('[')
						if v4 == nil {
							if !in.IsDelim(']') {
								v4 = make([]string, 0, 4)
							} else {
								v4 = []string{}
							}
						} else {
							v4 = (v4)[:0]
						}
						for !in.IsDelim(']') {
							var v5 string
							v5 = string(in.String())
							v4 = append(v4, v5)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Extra)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "groups":
			if in.IsNull() {
				in.Skip()
				out.Groups = nil
			} else {
				in.Delim('[')
				if out.Groups == nil {
					if !in.IsDelim(']') {
						out.Groups = make([]string, 0, 4)
					} else {
						out.Groups = []string{}
					}
				} else {
					out.Groups = (out.Groups)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.Groups = append(out.Groups, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "uid":
			out.UID = string(in.String())
		case "username":
			out.Username = string(in.String())
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
func easyjsonEfd2f10eEncodeGithubComKubewardenK8sObjectsApiAuthenticationV13(out *jwriter.Writer, in UserInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Extra) != 0 {
		const prefix string = ",\"extra\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('{')
			v7First := true
			for v7Name, v7Value := range in.Extra {
				if v7First {
					v7First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v7Name))
				out.RawByte(':')
				if v7Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v8, v9 := range v7Value {
						if v8 > 0 {
							out.RawByte(',')
						}
						out.String(string(v9))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Groups) != 0 {
		const prefix string = ",\"groups\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.Groups {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	if in.UID != "" {
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UID))
	}
	if in.Username != "" {
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	out.RawByte('}')
}
func easyjsonEfd2f10eDecodeGithubComKubewardenK8sObjectsApiAuthenticationV11(in *jlexer.Lexer, out *TokenReviewSpec) {
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
					var v12 string
					v12 = string(in.String())
					out.Audiences = append(out.Audiences, v12)
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
func easyjsonEfd2f10eEncodeGithubComKubewardenK8sObjectsApiAuthenticationV11(out *jwriter.Writer, in TokenReviewSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Audiences) != 0 {
		const prefix string = ",\"audiences\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v13, v14 := range in.Audiences {
				if v13 > 0 {
					out.RawByte(',')
				}
				out.String(string(v14))
			}
			out.RawByte(']')
		}
	}
	if in.Token != "" {
		const prefix string = ",\"token\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Token))
	}
	out.RawByte('}')
}
