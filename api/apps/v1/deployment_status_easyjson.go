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

func easyjson7381fb58DecodeGithubComKubewardenK8sObjectsApiAppsV1(in *jlexer.Lexer, out *DeploymentStatus) {
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
		case "availableReplicas":
			out.AvailableReplicas = int32(in.Int32())
		case "collisionCount":
			out.CollisionCount = int32(in.Int32())
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*DeploymentCondition, 0, 8)
					} else {
						out.Conditions = []*DeploymentCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *DeploymentCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(DeploymentCondition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "observedGeneration":
			out.ObservedGeneration = int64(in.Int64())
		case "readyReplicas":
			out.ReadyReplicas = int32(in.Int32())
		case "replicas":
			out.Replicas = int32(in.Int32())
		case "unavailableReplicas":
			out.UnavailableReplicas = int32(in.Int32())
		case "updatedReplicas":
			out.UpdatedReplicas = int32(in.Int32())
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
func easyjson7381fb58EncodeGithubComKubewardenK8sObjectsApiAppsV1(out *jwriter.Writer, in DeploymentStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AvailableReplicas != 0 {
		const prefix string = ",\"availableReplicas\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AvailableReplicas))
	}
	if in.CollisionCount != 0 {
		const prefix string = ",\"collisionCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CollisionCount))
	}
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Conditions {
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
	if in.ObservedGeneration != 0 {
		const prefix string = ",\"observedGeneration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ObservedGeneration))
	}
	if in.ReadyReplicas != 0 {
		const prefix string = ",\"readyReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ReadyReplicas))
	}
	if in.Replicas != 0 {
		const prefix string = ",\"replicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Replicas))
	}
	if in.UnavailableReplicas != 0 {
		const prefix string = ",\"unavailableReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.UnavailableReplicas))
	}
	if in.UpdatedReplicas != 0 {
		const prefix string = ",\"updatedReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.UpdatedReplicas))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeploymentStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7381fb58EncodeGithubComKubewardenK8sObjectsApiAppsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeploymentStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7381fb58EncodeGithubComKubewardenK8sObjectsApiAppsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeploymentStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7381fb58DecodeGithubComKubewardenK8sObjectsApiAppsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeploymentStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7381fb58DecodeGithubComKubewardenK8sObjectsApiAppsV1(l, v)
}
