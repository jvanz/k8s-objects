// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
	intstr "github.com/kubewarden/k8s-objects/apimachinery/pkg/util/intstr"
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

func easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV1(in *jlexer.Lexer, out *NetworkPolicy) {
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
					out.Spec = new(NetworkPolicySpec)
				}
				easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV11(in, out.Spec)
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
func easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV1(out *jwriter.Writer, in NetworkPolicy) {
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
		easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV11(out, *in.Spec)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NetworkPolicy) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NetworkPolicy) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NetworkPolicy) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NetworkPolicy) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV1(l, v)
}
func easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV11(in *jlexer.Lexer, out *NetworkPolicySpec) {
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
		case "egress":
			if in.IsNull() {
				in.Skip()
				out.Egress = nil
			} else {
				in.Delim('[')
				if out.Egress == nil {
					if !in.IsDelim(']') {
						out.Egress = make([]*NetworkPolicyEgressRule, 0, 8)
					} else {
						out.Egress = []*NetworkPolicyEgressRule{}
					}
				} else {
					out.Egress = (out.Egress)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *NetworkPolicyEgressRule
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(NetworkPolicyEgressRule)
						}
						easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV12(in, v1)
					}
					out.Egress = append(out.Egress, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ingress":
			if in.IsNull() {
				in.Skip()
				out.Ingress = nil
			} else {
				in.Delim('[')
				if out.Ingress == nil {
					if !in.IsDelim(']') {
						out.Ingress = make([]*NetworkPolicyIngressRule, 0, 8)
					} else {
						out.Ingress = []*NetworkPolicyIngressRule{}
					}
				} else {
					out.Ingress = (out.Ingress)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *NetworkPolicyIngressRule
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(NetworkPolicyIngressRule)
						}
						easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV13(in, v2)
					}
					out.Ingress = append(out.Ingress, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "podSelector":
			if in.IsNull() {
				in.Skip()
				out.PodSelector = nil
			} else {
				if out.PodSelector == nil {
					out.PodSelector = new(_v1.LabelSelector)
				}
				(*out.PodSelector).UnmarshalEasyJSON(in)
			}
		case "policyTypes":
			if in.IsNull() {
				in.Skip()
				out.PolicyTypes = nil
			} else {
				in.Delim('[')
				if out.PolicyTypes == nil {
					if !in.IsDelim(']') {
						out.PolicyTypes = make([]string, 0, 4)
					} else {
						out.PolicyTypes = []string{}
					}
				} else {
					out.PolicyTypes = (out.PolicyTypes)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.PolicyTypes = append(out.PolicyTypes, v3)
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
func easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV11(out *jwriter.Writer, in NetworkPolicySpec) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Egress) != 0 {
		const prefix string = ",\"egress\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v4, v5 := range in.Egress {
				if v4 > 0 {
					out.RawByte(',')
				}
				if v5 == nil {
					out.RawString("null")
				} else {
					easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV12(out, *v5)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Ingress) != 0 {
		const prefix string = ",\"ingress\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.Ingress {
				if v6 > 0 {
					out.RawByte(',')
				}
				if v7 == nil {
					out.RawString("null")
				} else {
					easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV13(out, *v7)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"podSelector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.PodSelector == nil {
			out.RawString("null")
		} else {
			(*in.PodSelector).MarshalEasyJSON(out)
		}
	}
	if len(in.PolicyTypes) != 0 {
		const prefix string = ",\"policyTypes\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v8, v9 := range in.PolicyTypes {
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
func easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV13(in *jlexer.Lexer, out *NetworkPolicyIngressRule) {
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
		case "from":
			if in.IsNull() {
				in.Skip()
				out.From = nil
			} else {
				in.Delim('[')
				if out.From == nil {
					if !in.IsDelim(']') {
						out.From = make([]*NetworkPolicyPeer, 0, 8)
					} else {
						out.From = []*NetworkPolicyPeer{}
					}
				} else {
					out.From = (out.From)[:0]
				}
				for !in.IsDelim(']') {
					var v10 *NetworkPolicyPeer
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(NetworkPolicyPeer)
						}
						easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV14(in, v10)
					}
					out.From = append(out.From, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ports":
			if in.IsNull() {
				in.Skip()
				out.Ports = nil
			} else {
				in.Delim('[')
				if out.Ports == nil {
					if !in.IsDelim(']') {
						out.Ports = make([]*NetworkPolicyPort, 0, 8)
					} else {
						out.Ports = []*NetworkPolicyPort{}
					}
				} else {
					out.Ports = (out.Ports)[:0]
				}
				for !in.IsDelim(']') {
					var v11 *NetworkPolicyPort
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(NetworkPolicyPort)
						}
						easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV15(in, v11)
					}
					out.Ports = append(out.Ports, v11)
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
func easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV13(out *jwriter.Writer, in NetworkPolicyIngressRule) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.From) != 0 {
		const prefix string = ",\"from\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v12, v13 := range in.From {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV14(out, *v13)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Ports) != 0 {
		const prefix string = ",\"ports\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.Ports {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV15(out, *v15)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV15(in *jlexer.Lexer, out *NetworkPolicyPort) {
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
		case "endPort":
			out.EndPort = int32(in.Int32())
		case "port":
			if in.IsNull() {
				in.Skip()
				out.Port = nil
			} else {
				if out.Port == nil {
					out.Port = new(intstr.IntOrString)
				}
				*out.Port = intstr.IntOrString(in.String())
			}
		case "protocol":
			out.Protocol = string(in.String())
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
func easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV15(out *jwriter.Writer, in NetworkPolicyPort) {
	out.RawByte('{')
	first := true
	_ = first
	if in.EndPort != 0 {
		const prefix string = ",\"endPort\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.EndPort))
	}
	if in.Port != nil {
		const prefix string = ",\"port\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Port))
	}
	if in.Protocol != "" {
		const prefix string = ",\"protocol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Protocol))
	}
	out.RawByte('}')
}
func easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV14(in *jlexer.Lexer, out *NetworkPolicyPeer) {
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
		case "ipBlock":
			if in.IsNull() {
				in.Skip()
				out.IPBlock = nil
			} else {
				if out.IPBlock == nil {
					out.IPBlock = new(IPBlock)
				}
				(*out.IPBlock).UnmarshalEasyJSON(in)
			}
		case "namespaceSelector":
			if in.IsNull() {
				in.Skip()
				out.NamespaceSelector = nil
			} else {
				if out.NamespaceSelector == nil {
					out.NamespaceSelector = new(_v1.LabelSelector)
				}
				(*out.NamespaceSelector).UnmarshalEasyJSON(in)
			}
		case "podSelector":
			if in.IsNull() {
				in.Skip()
				out.PodSelector = nil
			} else {
				if out.PodSelector == nil {
					out.PodSelector = new(_v1.LabelSelector)
				}
				(*out.PodSelector).UnmarshalEasyJSON(in)
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
func easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV14(out *jwriter.Writer, in NetworkPolicyPeer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.IPBlock != nil {
		const prefix string = ",\"ipBlock\":"
		first = false
		out.RawString(prefix[1:])
		(*in.IPBlock).MarshalEasyJSON(out)
	}
	if in.NamespaceSelector != nil {
		const prefix string = ",\"namespaceSelector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.NamespaceSelector).MarshalEasyJSON(out)
	}
	if in.PodSelector != nil {
		const prefix string = ",\"podSelector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.PodSelector).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV12(in *jlexer.Lexer, out *NetworkPolicyEgressRule) {
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
		case "ports":
			if in.IsNull() {
				in.Skip()
				out.Ports = nil
			} else {
				in.Delim('[')
				if out.Ports == nil {
					if !in.IsDelim(']') {
						out.Ports = make([]*NetworkPolicyPort, 0, 8)
					} else {
						out.Ports = []*NetworkPolicyPort{}
					}
				} else {
					out.Ports = (out.Ports)[:0]
				}
				for !in.IsDelim(']') {
					var v16 *NetworkPolicyPort
					if in.IsNull() {
						in.Skip()
						v16 = nil
					} else {
						if v16 == nil {
							v16 = new(NetworkPolicyPort)
						}
						easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV15(in, v16)
					}
					out.Ports = append(out.Ports, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "to":
			if in.IsNull() {
				in.Skip()
				out.To = nil
			} else {
				in.Delim('[')
				if out.To == nil {
					if !in.IsDelim(']') {
						out.To = make([]*NetworkPolicyPeer, 0, 8)
					} else {
						out.To = []*NetworkPolicyPeer{}
					}
				} else {
					out.To = (out.To)[:0]
				}
				for !in.IsDelim(']') {
					var v17 *NetworkPolicyPeer
					if in.IsNull() {
						in.Skip()
						v17 = nil
					} else {
						if v17 == nil {
							v17 = new(NetworkPolicyPeer)
						}
						easyjsonF1367303DecodeGithubComKubewardenK8sObjectsApiNetworkingV14(in, v17)
					}
					out.To = append(out.To, v17)
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
func easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV12(out *jwriter.Writer, in NetworkPolicyEgressRule) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Ports) != 0 {
		const prefix string = ",\"ports\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v18, v19 := range in.Ports {
				if v18 > 0 {
					out.RawByte(',')
				}
				if v19 == nil {
					out.RawString("null")
				} else {
					easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV15(out, *v19)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.To) != 0 {
		const prefix string = ",\"to\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.To {
				if v20 > 0 {
					out.RawByte(',')
				}
				if v21 == nil {
					out.RawString("null")
				} else {
					easyjsonF1367303EncodeGithubComKubewardenK8sObjectsApiNetworkingV14(out, *v21)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
