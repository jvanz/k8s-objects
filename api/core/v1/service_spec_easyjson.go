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

func easyjson9d06a251DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *ServiceSpec) {
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
		case "allocateLoadBalancerNodePorts":
			out.AllocateLoadBalancerNodePorts = bool(in.Bool())
		case "clusterIP":
			out.ClusterIP = string(in.String())
		case "clusterIPs":
			if in.IsNull() {
				in.Skip()
				out.ClusterIPs = nil
			} else {
				in.Delim('[')
				if out.ClusterIPs == nil {
					if !in.IsDelim(']') {
						out.ClusterIPs = make([]string, 0, 4)
					} else {
						out.ClusterIPs = []string{}
					}
				} else {
					out.ClusterIPs = (out.ClusterIPs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.ClusterIPs = append(out.ClusterIPs, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "externalIPs":
			if in.IsNull() {
				in.Skip()
				out.ExternalIPs = nil
			} else {
				in.Delim('[')
				if out.ExternalIPs == nil {
					if !in.IsDelim(']') {
						out.ExternalIPs = make([]string, 0, 4)
					} else {
						out.ExternalIPs = []string{}
					}
				} else {
					out.ExternalIPs = (out.ExternalIPs)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.ExternalIPs = append(out.ExternalIPs, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "externalName":
			out.ExternalName = string(in.String())
		case "externalTrafficPolicy":
			out.ExternalTrafficPolicy = string(in.String())
		case "healthCheckNodePort":
			out.HealthCheckNodePort = int32(in.Int32())
		case "internalTrafficPolicy":
			out.InternalTrafficPolicy = string(in.String())
		case "ipFamilies":
			if in.IsNull() {
				in.Skip()
				out.IPFamilies = nil
			} else {
				in.Delim('[')
				if out.IPFamilies == nil {
					if !in.IsDelim(']') {
						out.IPFamilies = make([]string, 0, 4)
					} else {
						out.IPFamilies = []string{}
					}
				} else {
					out.IPFamilies = (out.IPFamilies)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.IPFamilies = append(out.IPFamilies, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ipFamilyPolicy":
			out.IPFamilyPolicy = string(in.String())
		case "loadBalancerClass":
			out.LoadBalancerClass = string(in.String())
		case "loadBalancerIP":
			out.LoadBalancerIP = string(in.String())
		case "loadBalancerSourceRanges":
			if in.IsNull() {
				in.Skip()
				out.LoadBalancerSourceRanges = nil
			} else {
				in.Delim('[')
				if out.LoadBalancerSourceRanges == nil {
					if !in.IsDelim(']') {
						out.LoadBalancerSourceRanges = make([]string, 0, 4)
					} else {
						out.LoadBalancerSourceRanges = []string{}
					}
				} else {
					out.LoadBalancerSourceRanges = (out.LoadBalancerSourceRanges)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.LoadBalancerSourceRanges = append(out.LoadBalancerSourceRanges, v4)
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
						out.Ports = make([]*ServicePort, 0, 8)
					} else {
						out.Ports = []*ServicePort{}
					}
				} else {
					out.Ports = (out.Ports)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *ServicePort
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(ServicePort)
						}
						(*v5).UnmarshalEasyJSON(in)
					}
					out.Ports = append(out.Ports, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "publishNotReadyAddresses":
			out.PublishNotReadyAddresses = bool(in.Bool())
		case "selector":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Selector = make(map[string]string)
				} else {
					out.Selector = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v6 string
					v6 = string(in.String())
					(out.Selector)[key] = v6
					in.WantComma()
				}
				in.Delim('}')
			}
		case "sessionAffinity":
			out.SessionAffinity = string(in.String())
		case "sessionAffinityConfig":
			if in.IsNull() {
				in.Skip()
				out.SessionAffinityConfig = nil
			} else {
				if out.SessionAffinityConfig == nil {
					out.SessionAffinityConfig = new(SessionAffinityConfig)
				}
				easyjson9d06a251DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.SessionAffinityConfig)
			}
		case "topologyKeys":
			if in.IsNull() {
				in.Skip()
				out.TopologyKeys = nil
			} else {
				in.Delim('[')
				if out.TopologyKeys == nil {
					if !in.IsDelim(']') {
						out.TopologyKeys = make([]string, 0, 4)
					} else {
						out.TopologyKeys = []string{}
					}
				} else {
					out.TopologyKeys = (out.TopologyKeys)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.TopologyKeys = append(out.TopologyKeys, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "type":
			out.Type = string(in.String())
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
func easyjson9d06a251EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in ServiceSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllocateLoadBalancerNodePorts {
		const prefix string = ",\"allocateLoadBalancerNodePorts\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.AllocateLoadBalancerNodePorts))
	}
	if in.ClusterIP != "" {
		const prefix string = ",\"clusterIP\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClusterIP))
	}
	if len(in.ClusterIPs) != 0 {
		const prefix string = ",\"clusterIPs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.ClusterIPs {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if len(in.ExternalIPs) != 0 {
		const prefix string = ",\"externalIPs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.ExternalIPs {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	if in.ExternalName != "" {
		const prefix string = ",\"externalName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExternalName))
	}
	if in.ExternalTrafficPolicy != "" {
		const prefix string = ",\"externalTrafficPolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExternalTrafficPolicy))
	}
	if in.HealthCheckNodePort != 0 {
		const prefix string = ",\"healthCheckNodePort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.HealthCheckNodePort))
	}
	if in.InternalTrafficPolicy != "" {
		const prefix string = ",\"internalTrafficPolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InternalTrafficPolicy))
	}
	if len(in.IPFamilies) != 0 {
		const prefix string = ",\"ipFamilies\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v12, v13 := range in.IPFamilies {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	if in.IPFamilyPolicy != "" {
		const prefix string = ",\"ipFamilyPolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IPFamilyPolicy))
	}
	if in.LoadBalancerClass != "" {
		const prefix string = ",\"loadBalancerClass\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LoadBalancerClass))
	}
	if in.LoadBalancerIP != "" {
		const prefix string = ",\"loadBalancerIP\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LoadBalancerIP))
	}
	if len(in.LoadBalancerSourceRanges) != 0 {
		const prefix string = ",\"loadBalancerSourceRanges\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.LoadBalancerSourceRanges {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
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
			for v16, v17 := range in.Ports {
				if v16 > 0 {
					out.RawByte(',')
				}
				if v17 == nil {
					out.RawString("null")
				} else {
					(*v17).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.PublishNotReadyAddresses {
		const prefix string = ",\"publishNotReadyAddresses\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.PublishNotReadyAddresses))
	}
	if len(in.Selector) != 0 {
		const prefix string = ",\"selector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v18First := true
			for v18Name, v18Value := range in.Selector {
				if v18First {
					v18First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v18Name))
				out.RawByte(':')
				out.String(string(v18Value))
			}
			out.RawByte('}')
		}
	}
	if in.SessionAffinity != "" {
		const prefix string = ",\"sessionAffinity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SessionAffinity))
	}
	if in.SessionAffinityConfig != nil {
		const prefix string = ",\"sessionAffinityConfig\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9d06a251EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.SessionAffinityConfig)
	}
	if len(in.TopologyKeys) != 0 {
		const prefix string = ",\"topologyKeys\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v19, v20 := range in.TopologyKeys {
				if v19 > 0 {
					out.RawByte(',')
				}
				out.String(string(v20))
			}
			out.RawByte(']')
		}
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9d06a251EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9d06a251EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9d06a251DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9d06a251DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson9d06a251DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *SessionAffinityConfig) {
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
		case "clientIP":
			if in.IsNull() {
				in.Skip()
				out.ClientIP = nil
			} else {
				if out.ClientIP == nil {
					out.ClientIP = new(ClientIPConfig)
				}
				(*out.ClientIP).UnmarshalEasyJSON(in)
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
func easyjson9d06a251EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in SessionAffinityConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ClientIP != nil {
		const prefix string = ",\"clientIP\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ClientIP).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
