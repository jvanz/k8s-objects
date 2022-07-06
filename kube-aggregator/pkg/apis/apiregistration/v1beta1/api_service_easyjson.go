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

func easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta1(in *jlexer.Lexer, out *APIService) {
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
					out.Spec = new(APIServiceSpec)
				}
				easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(APIServiceStatus)
				}
				easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta12(in, out.Status)
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
func easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta1(out *jwriter.Writer, in APIService) {
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
		easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v APIService) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v APIService) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *APIService) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *APIService) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta1(l, v)
}
func easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta12(in *jlexer.Lexer, out *APIServiceStatus) {
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
						out.Conditions = make([]*APIServiceCondition, 0, 8)
					} else {
						out.Conditions = []*APIServiceCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *APIServiceCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(APIServiceCondition)
						}
						easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta13(in, v1)
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
func easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta12(out *jwriter.Writer, in APIServiceStatus) {
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
					easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta13(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta13(in *jlexer.Lexer, out *APIServiceCondition) {
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
		case "lastTransitionTime":
			if in.IsNull() {
				in.Skip()
				out.LastTransitionTime = nil
			} else {
				if out.LastTransitionTime == nil {
					out.LastTransitionTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastTransitionTime).UnmarshalJSON(data))
				}
			}
		case "message":
			out.Message = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(string)
				}
				*out.Status = string(in.String())
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
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
func easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta13(out *jwriter.Writer, in APIServiceCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastTransitionTime != nil {
		const prefix string = ",\"lastTransitionTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.LastTransitionTime).MarshalJSON())
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Reason))
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Status == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Status))
		}
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}
func easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta11(in *jlexer.Lexer, out *APIServiceSpec) {
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
		case "group":
			out.Group = string(in.String())
		case "groupPriorityMinimum":
			if in.IsNull() {
				in.Skip()
				out.GroupPriorityMinimum = nil
			} else {
				if out.GroupPriorityMinimum == nil {
					out.GroupPriorityMinimum = new(int32)
				}
				*out.GroupPriorityMinimum = int32(in.Int32())
			}
		case "insecureSkipTLSVerify":
			out.InsecureSkipTLSVerify = bool(in.Bool())
		case "service":
			if in.IsNull() {
				in.Skip()
				out.Service = nil
			} else {
				if out.Service == nil {
					out.Service = new(ServiceReference)
				}
				easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta14(in, out.Service)
			}
		case "version":
			out.Version = string(in.String())
		case "versionPriority":
			if in.IsNull() {
				in.Skip()
				out.VersionPriority = nil
			} else {
				if out.VersionPriority == nil {
					out.VersionPriority = new(int32)
				}
				*out.VersionPriority = int32(in.Int32())
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
func easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta11(out *jwriter.Writer, in APIServiceSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.CaBundle) != 0 {
		const prefix string = ",\"caBundle\":"
		first = false
		out.RawString(prefix[1:])
		out.Base64Bytes(in.CaBundle)
	}
	if in.Group != "" {
		const prefix string = ",\"group\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Group))
	}
	{
		const prefix string = ",\"groupPriorityMinimum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.GroupPriorityMinimum == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.GroupPriorityMinimum))
		}
	}
	if in.InsecureSkipTLSVerify {
		const prefix string = ",\"insecureSkipTLSVerify\":"
		out.RawString(prefix)
		out.Bool(bool(in.InsecureSkipTLSVerify))
	}
	if in.Service != nil {
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta14(out, *in.Service)
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"versionPriority\":"
		out.RawString(prefix)
		if in.VersionPriority == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.VersionPriority))
		}
	}
	out.RawByte('}')
}
func easyjson4f973ab6DecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta14(in *jlexer.Lexer, out *ServiceReference) {
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
			out.Name = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
		case "port":
			out.Port = int32(in.Int32())
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
func easyjson4f973ab6EncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1beta14(out *jwriter.Writer, in ServiceReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Namespace))
	}
	if in.Port != 0 {
		const prefix string = ",\"port\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Port))
	}
	out.RawByte('}')
}
