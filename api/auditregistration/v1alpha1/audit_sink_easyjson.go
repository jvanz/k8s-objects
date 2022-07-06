// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

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

func easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(in *jlexer.Lexer, out *AuditSink) {
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
					out.Spec = new(AuditSinkSpec)
				}
				easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha11(in, out.Spec)
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
func easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(out *jwriter.Writer, in AuditSink) {
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
		easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha11(out, *in.Spec)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuditSink) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuditSink) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuditSink) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuditSink) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(l, v)
}
func easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha11(in *jlexer.Lexer, out *AuditSinkSpec) {
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
		case "policy":
			if in.IsNull() {
				in.Skip()
				out.Policy = nil
			} else {
				if out.Policy == nil {
					out.Policy = new(Policy)
				}
				easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha12(in, out.Policy)
			}
		case "webhook":
			if in.IsNull() {
				in.Skip()
				out.Webhook = nil
			} else {
				if out.Webhook == nil {
					out.Webhook = new(Webhook)
				}
				easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha13(in, out.Webhook)
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
func easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha11(out *jwriter.Writer, in AuditSinkSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"policy\":"
		out.RawString(prefix[1:])
		if in.Policy == nil {
			out.RawString("null")
		} else {
			easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha12(out, *in.Policy)
		}
	}
	{
		const prefix string = ",\"webhook\":"
		out.RawString(prefix)
		if in.Webhook == nil {
			out.RawString("null")
		} else {
			easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha13(out, *in.Webhook)
		}
	}
	out.RawByte('}')
}
func easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha13(in *jlexer.Lexer, out *Webhook) {
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
		case "clientConfig":
			if in.IsNull() {
				in.Skip()
				out.ClientConfig = nil
			} else {
				if out.ClientConfig == nil {
					out.ClientConfig = new(WebhookClientConfig)
				}
				easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha14(in, out.ClientConfig)
			}
		case "throttle":
			if in.IsNull() {
				in.Skip()
				out.Throttle = nil
			} else {
				if out.Throttle == nil {
					out.Throttle = new(WebhookThrottleConfig)
				}
				easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha15(in, out.Throttle)
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
func easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha13(out *jwriter.Writer, in Webhook) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"clientConfig\":"
		out.RawString(prefix[1:])
		if in.ClientConfig == nil {
			out.RawString("null")
		} else {
			easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha14(out, *in.ClientConfig)
		}
	}
	if in.Throttle != nil {
		const prefix string = ",\"throttle\":"
		out.RawString(prefix)
		easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha15(out, *in.Throttle)
	}
	out.RawByte('}')
}
func easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha15(in *jlexer.Lexer, out *WebhookThrottleConfig) {
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
		case "burst":
			out.Burst = int64(in.Int64())
		case "qps":
			out.QPS = int64(in.Int64())
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
func easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha15(out *jwriter.Writer, in WebhookThrottleConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Burst != 0 {
		const prefix string = ",\"burst\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Burst))
	}
	if in.QPS != 0 {
		const prefix string = ",\"qps\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.QPS))
	}
	out.RawByte('}')
}
func easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha14(in *jlexer.Lexer, out *WebhookClientConfig) {
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
		case "caBundle":
			if in.IsNull() {
				in.Skip()
				out.CaBundle = nil
			} else {
				out.CaBundle = in.Bytes()
			}
		case "service":
			if in.IsNull() {
				in.Skip()
				out.Service = nil
			} else {
				if out.Service == nil {
					out.Service = new(ServiceReference)
				}
				easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha16(in, out.Service)
			}
		case "url":
			out.URL = string(in.String())
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
func easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha14(out *jwriter.Writer, in WebhookClientConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.CaBundle) != 0 {
		const prefix string = ",\"caBundle\":"
		first = false
		out.RawString(prefix[1:])
		out.Base64Bytes(in.CaBundle)
	}
	if in.Service != nil {
		const prefix string = ",\"service\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha16(out, *in.Service)
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	out.RawByte('}')
}
func easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha16(in *jlexer.Lexer, out *ServiceReference) {
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
		case "namespace":
			if in.IsNull() {
				in.Skip()
				out.Namespace = nil
			} else {
				if out.Namespace == nil {
					out.Namespace = new(string)
				}
				*out.Namespace = string(in.String())
			}
		case "path":
			out.Path = string(in.String())
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
func easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha16(out *jwriter.Writer, in ServiceReference) {
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
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		if in.Namespace == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Namespace))
		}
	}
	if in.Path != "" {
		const prefix string = ",\"path\":"
		out.RawString(prefix)
		out.String(string(in.Path))
	}
	out.RawByte('}')
}
func easyjson637368a5DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha12(in *jlexer.Lexer, out *Policy) {
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
		case "level":
			if in.IsNull() {
				in.Skip()
				out.Level = nil
			} else {
				if out.Level == nil {
					out.Level = new(string)
				}
				*out.Level = string(in.String())
			}
		case "stages":
			if in.IsNull() {
				in.Skip()
				out.Stages = nil
			} else {
				in.Delim('[')
				if out.Stages == nil {
					if !in.IsDelim(']') {
						out.Stages = make([]string, 0, 4)
					} else {
						out.Stages = []string{}
					}
				} else {
					out.Stages = (out.Stages)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Stages = append(out.Stages, v4)
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
func easyjson637368a5EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha12(out *jwriter.Writer, in Policy) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"level\":"
		out.RawString(prefix[1:])
		if in.Level == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Level))
		}
	}
	if len(in.Stages) != 0 {
		const prefix string = ",\"stages\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Stages {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
