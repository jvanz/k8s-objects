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

func easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *Event) {
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
		case "action":
			out.Action = string(in.String())
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "count":
			out.Count = int32(in.Int32())
		case "eventTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EventTime).UnmarshalJSON(data))
			}
		case "firstTimestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.FirstTimestamp).UnmarshalJSON(data))
			}
		case "involvedObject":
			if in.IsNull() {
				in.Skip()
				out.InvolvedObject = nil
			} else {
				if out.InvolvedObject == nil {
					out.InvolvedObject = new(ObjectReference)
				}
				easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.InvolvedObject)
			}
		case "kind":
			out.Kind = string(in.String())
		case "lastTimestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastTimestamp).UnmarshalJSON(data))
			}
		case "message":
			out.Message = string(in.String())
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
		case "reason":
			out.Reason = string(in.String())
		case "related":
			if in.IsNull() {
				in.Skip()
				out.Related = nil
			} else {
				if out.Related == nil {
					out.Related = new(ObjectReference)
				}
				easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.Related)
			}
		case "reportingComponent":
			out.ReportingComponent = string(in.String())
		case "reportingInstance":
			out.ReportingInstance = string(in.String())
		case "series":
			if in.IsNull() {
				in.Skip()
				out.Series = nil
			} else {
				if out.Series == nil {
					out.Series = new(EventSeries)
				}
				easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.Series)
			}
		case "source":
			if in.IsNull() {
				in.Skip()
				out.Source = nil
			} else {
				if out.Source == nil {
					out.Source = new(EventSource)
				}
				easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV13(in, out.Source)
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
func easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Action != "" {
		const prefix string = ",\"action\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Action))
	}
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.APIVersion))
	}
	if in.Count != 0 {
		const prefix string = ",\"count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Count))
	}
	if true {
		const prefix string = ",\"eventTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.EventTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"firstTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.FirstTimestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"involvedObject\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.InvolvedObject == nil {
			out.RawString("null")
		} else {
			easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.InvolvedObject)
		}
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	if true {
		const prefix string = ",\"lastTimestamp\":"
		out.RawString(prefix)
		out.Raw((in.LastTimestamp).MarshalJSON())
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		if in.Metadata == nil {
			out.RawString("null")
		} else {
			(*in.Metadata).MarshalEasyJSON(out)
		}
	}
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		out.RawString(prefix)
		out.String(string(in.Reason))
	}
	if in.Related != nil {
		const prefix string = ",\"related\":"
		out.RawString(prefix)
		easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.Related)
	}
	if in.ReportingComponent != "" {
		const prefix string = ",\"reportingComponent\":"
		out.RawString(prefix)
		out.String(string(in.ReportingComponent))
	}
	if in.ReportingInstance != "" {
		const prefix string = ",\"reportingInstance\":"
		out.RawString(prefix)
		out.String(string(in.ReportingInstance))
	}
	if in.Series != nil {
		const prefix string = ",\"series\":"
		out.RawString(prefix)
		easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.Series)
	}
	if in.Source != nil {
		const prefix string = ",\"source\":"
		out.RawString(prefix)
		easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV13(out, *in.Source)
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV13(in *jlexer.Lexer, out *EventSource) {
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
		case "component":
			out.Component = string(in.String())
		case "host":
			out.Host = string(in.String())
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
func easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV13(out *jwriter.Writer, in EventSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Component != "" {
		const prefix string = ",\"component\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Component))
	}
	if in.Host != "" {
		const prefix string = ",\"host\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Host))
	}
	out.RawByte('}')
}
func easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *EventSeries) {
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
		case "count":
			out.Count = int32(in.Int32())
		case "lastObservedTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastObservedTime).UnmarshalJSON(data))
			}
		case "state":
			out.State = string(in.String())
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
func easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in EventSeries) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Count != 0 {
		const prefix string = ",\"count\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Count))
	}
	if true {
		const prefix string = ",\"lastObservedTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastObservedTime).MarshalJSON())
	}
	if in.State != "" {
		const prefix string = ",\"state\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.State))
	}
	out.RawByte('}')
}
func easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *ObjectReference) {
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
		case "fieldPath":
			out.FieldPath = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
		case "resourceVersion":
			out.ResourceVersion = string(in.String())
		case "uid":
			out.UID = string(in.String())
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
func easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in ObjectReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	if in.FieldPath != "" {
		const prefix string = ",\"fieldPath\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FieldPath))
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
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
	if in.ResourceVersion != "" {
		const prefix string = ",\"resourceVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ResourceVersion))
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
	out.RawByte('}')
}
