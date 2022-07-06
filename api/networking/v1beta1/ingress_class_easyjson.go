// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

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

func easyjsonBe26a562DecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(in *jlexer.Lexer, out *IngressClass) {
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
					out.Spec = new(IngressClassSpec)
				}
				easyjsonBe26a562DecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta11(in, out.Spec)
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
func easyjsonBe26a562EncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(out *jwriter.Writer, in IngressClass) {
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
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonBe26a562EncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta11(out, *in.Spec)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IngressClass) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBe26a562EncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IngressClass) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBe26a562EncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IngressClass) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBe26a562DecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IngressClass) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBe26a562DecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(l, v)
}
func easyjsonBe26a562DecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta11(in *jlexer.Lexer, out *IngressClassSpec) {
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
		case "controller":
			out.Controller = string(in.String())
		case "parameters":
			if in.IsNull() {
				in.Skip()
				out.Parameters = nil
			} else {
				if out.Parameters == nil {
					out.Parameters = new(IngressClassParametersReference)
				}
				easyjsonBe26a562DecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta12(in, out.Parameters)
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
func easyjsonBe26a562EncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta11(out *jwriter.Writer, in IngressClassSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Controller != "" {
		const prefix string = ",\"controller\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Controller))
	}
	if in.Parameters != nil {
		const prefix string = ",\"parameters\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonBe26a562EncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta12(out, *in.Parameters)
	}
	out.RawByte('}')
}
func easyjsonBe26a562DecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta12(in *jlexer.Lexer, out *IngressClassParametersReference) {
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
		case "apiGroup":
			out.APIGroup = string(in.String())
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(string)
				}
				*out.Kind = string(in.String())
			}
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
		case "namespace":
			out.Namespace = string(in.String())
		case "scope":
			out.Scope = string(in.String())
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
func easyjsonBe26a562EncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta12(out *jwriter.Writer, in IngressClassParametersReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIGroup != "" {
		const prefix string = ",\"apiGroup\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIGroup))
	}
	{
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Kind == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Kind))
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		out.String(string(in.Namespace))
	}
	if in.Scope != "" {
		const prefix string = ",\"scope\":"
		out.RawString(prefix)
		out.String(string(in.Scope))
	}
	out.RawByte('}')
}
