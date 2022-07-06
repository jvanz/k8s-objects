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

func easyjson51a8e721DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(in *jlexer.Lexer, out *JSONSchemaProps) {
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
		case "$ref":
			out.DollarRef = string(in.String())
		case "$schema":
			out.DollarSchema = string(in.String())
		case "additionalItems":
			(out.AdditionalItems).UnmarshalEasyJSON(in)
		case "additionalProperties":
			(out.AdditionalProperties).UnmarshalEasyJSON(in)
		case "allOf":
			if in.IsNull() {
				in.Skip()
				out.AllOf = nil
			} else {
				in.Delim('[')
				if out.AllOf == nil {
					if !in.IsDelim(']') {
						out.AllOf = make([]*JSONSchemaProps, 0, 8)
					} else {
						out.AllOf = []*JSONSchemaProps{}
					}
				} else {
					out.AllOf = (out.AllOf)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(JSONSchemaProps)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.AllOf = append(out.AllOf, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "anyOf":
			if in.IsNull() {
				in.Skip()
				out.AnyOf = nil
			} else {
				in.Delim('[')
				if out.AnyOf == nil {
					if !in.IsDelim(']') {
						out.AnyOf = make([]*JSONSchemaProps, 0, 8)
					} else {
						out.AnyOf = []*JSONSchemaProps{}
					}
				} else {
					out.AnyOf = (out.AnyOf)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(JSONSchemaProps)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.AnyOf = append(out.AnyOf, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "default":
			(out.Default).UnmarshalEasyJSON(in)
		case "definitions":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Definitions = make(map[string]*JSONSchemaProps)
				} else {
					out.Definitions = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(JSONSchemaProps)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					(out.Definitions)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "dependencies":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Dependencies = make(map[string]easyjson.RawMessage)
				} else {
					out.Dependencies = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 easyjson.RawMessage
					(v4).UnmarshalEasyJSON(in)
					(out.Dependencies)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "description":
			out.Description = string(in.String())
		case "enum":
			if in.IsNull() {
				in.Skip()
				out.Enum = nil
			} else {
				in.Delim('[')
				if out.Enum == nil {
					if !in.IsDelim(']') {
						out.Enum = make([]easyjson.RawMessage, 0, 2)
					} else {
						out.Enum = []easyjson.RawMessage{}
					}
				} else {
					out.Enum = (out.Enum)[:0]
				}
				for !in.IsDelim(']') {
					var v5 easyjson.RawMessage
					(v5).UnmarshalEasyJSON(in)
					out.Enum = append(out.Enum, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "example":
			(out.Example).UnmarshalEasyJSON(in)
		case "exclusiveMaximum":
			out.ExclusiveMaximum = bool(in.Bool())
		case "exclusiveMinimum":
			out.ExclusiveMinimum = bool(in.Bool())
		case "externalDocs":
			if in.IsNull() {
				in.Skip()
				out.ExternalDocs = nil
			} else {
				if out.ExternalDocs == nil {
					out.ExternalDocs = new(ExternalDocumentation)
				}
				(*out.ExternalDocs).UnmarshalEasyJSON(in)
			}
		case "format":
			out.Format = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "items":
			(out.Items).UnmarshalEasyJSON(in)
		case "maxItems":
			out.MaxItems = int64(in.Int64())
		case "maxLength":
			out.MaxLength = int64(in.Int64())
		case "maxProperties":
			out.MaxProperties = int64(in.Int64())
		case "maximum":
			out.Maximum = float64(in.Float64())
		case "minItems":
			out.MinItems = int64(in.Int64())
		case "minLength":
			out.MinLength = int64(in.Int64())
		case "minProperties":
			out.MinProperties = int64(in.Int64())
		case "minimum":
			out.Minimum = float64(in.Float64())
		case "multipleOf":
			out.MultipleOf = float64(in.Float64())
		case "not":
			if in.IsNull() {
				in.Skip()
				out.Not = nil
			} else {
				if out.Not == nil {
					out.Not = new(JSONSchemaProps)
				}
				(*out.Not).UnmarshalEasyJSON(in)
			}
		case "nullable":
			out.Nullable = bool(in.Bool())
		case "oneOf":
			if in.IsNull() {
				in.Skip()
				out.OneOf = nil
			} else {
				in.Delim('[')
				if out.OneOf == nil {
					if !in.IsDelim(']') {
						out.OneOf = make([]*JSONSchemaProps, 0, 8)
					} else {
						out.OneOf = []*JSONSchemaProps{}
					}
				} else {
					out.OneOf = (out.OneOf)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(JSONSchemaProps)
						}
						(*v6).UnmarshalEasyJSON(in)
					}
					out.OneOf = append(out.OneOf, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pattern":
			out.Pattern = string(in.String())
		case "patternProperties":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.PatternProperties = make(map[string]*JSONSchemaProps)
				} else {
					out.PatternProperties = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(JSONSchemaProps)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					(out.PatternProperties)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		case "properties":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Properties = make(map[string]*JSONSchemaProps)
				} else {
					out.Properties = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v8 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(JSONSchemaProps)
						}
						(*v8).UnmarshalEasyJSON(in)
					}
					(out.Properties)[key] = v8
					in.WantComma()
				}
				in.Delim('}')
			}
		case "required":
			if in.IsNull() {
				in.Skip()
				out.Required = nil
			} else {
				in.Delim('[')
				if out.Required == nil {
					if !in.IsDelim(']') {
						out.Required = make([]string, 0, 4)
					} else {
						out.Required = []string{}
					}
				} else {
					out.Required = (out.Required)[:0]
				}
				for !in.IsDelim(']') {
					var v9 string
					v9 = string(in.String())
					out.Required = append(out.Required, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "title":
			out.Title = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "uniqueItems":
			out.UniqueItems = bool(in.Bool())
		case "x-kubernetes-embedded-resource":
			out.XKubernetesEmbeddedResource = bool(in.Bool())
		case "x-kubernetes-int-or-string":
			out.XKubernetesIntOrString = bool(in.Bool())
		case "x-kubernetes-list-map-keys":
			if in.IsNull() {
				in.Skip()
				out.XKubernetesListMapKeys = nil
			} else {
				in.Delim('[')
				if out.XKubernetesListMapKeys == nil {
					if !in.IsDelim(']') {
						out.XKubernetesListMapKeys = make([]string, 0, 4)
					} else {
						out.XKubernetesListMapKeys = []string{}
					}
				} else {
					out.XKubernetesListMapKeys = (out.XKubernetesListMapKeys)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.XKubernetesListMapKeys = append(out.XKubernetesListMapKeys, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "x-kubernetes-list-type":
			out.XKubernetesListType = string(in.String())
		case "x-kubernetes-map-type":
			out.XKubernetesMapType = string(in.String())
		case "x-kubernetes-preserve-unknown-fields":
			out.XKubernetesPreserveUnknownFields = bool(in.Bool())
		case "x-kubernetes-validations":
			if in.IsNull() {
				in.Skip()
				out.XKubernetesValidations = nil
			} else {
				in.Delim('[')
				if out.XKubernetesValidations == nil {
					if !in.IsDelim(']') {
						out.XKubernetesValidations = make([]*ValidationRule, 0, 8)
					} else {
						out.XKubernetesValidations = []*ValidationRule{}
					}
				} else {
					out.XKubernetesValidations = (out.XKubernetesValidations)[:0]
				}
				for !in.IsDelim(']') {
					var v11 *ValidationRule
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(ValidationRule)
						}
						easyjson51a8e721DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV11(in, v11)
					}
					out.XKubernetesValidations = append(out.XKubernetesValidations, v11)
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
func easyjson51a8e721EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(out *jwriter.Writer, in JSONSchemaProps) {
	out.RawByte('{')
	first := true
	_ = first
	if in.DollarRef != "" {
		const prefix string = ",\"$ref\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.DollarRef))
	}
	if in.DollarSchema != "" {
		const prefix string = ",\"$schema\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DollarSchema))
	}
	if (in.AdditionalItems).IsDefined() {
		const prefix string = ",\"additionalItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.AdditionalItems).MarshalEasyJSON(out)
	}
	if (in.AdditionalProperties).IsDefined() {
		const prefix string = ",\"additionalProperties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.AdditionalProperties).MarshalEasyJSON(out)
	}
	if len(in.AllOf) != 0 {
		const prefix string = ",\"allOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v12, v13 := range in.AllOf {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					(*v13).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.AnyOf) != 0 {
		const prefix string = ",\"anyOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.AnyOf {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if (in.Default).IsDefined() {
		const prefix string = ",\"default\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Default).MarshalEasyJSON(out)
	}
	if len(in.Definitions) != 0 {
		const prefix string = ",\"definitions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v16First := true
			for v16Name, v16Value := range in.Definitions {
				if v16First {
					v16First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v16Name))
				out.RawByte(':')
				if v16Value == nil {
					out.RawString("null")
				} else {
					(*v16Value).MarshalEasyJSON(out)
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Dependencies) != 0 {
		const prefix string = ",\"dependencies\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v17First := true
			for v17Name, v17Value := range in.Dependencies {
				if v17First {
					v17First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v17Name))
				out.RawByte(':')
				(v17Value).MarshalEasyJSON(out)
			}
			out.RawByte('}')
		}
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if len(in.Enum) != 0 {
		const prefix string = ",\"enum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.Enum {
				if v18 > 0 {
					out.RawByte(',')
				}
				(v19).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if (in.Example).IsDefined() {
		const prefix string = ",\"example\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Example).MarshalEasyJSON(out)
	}
	if in.ExclusiveMaximum {
		const prefix string = ",\"exclusiveMaximum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ExclusiveMaximum))
	}
	if in.ExclusiveMinimum {
		const prefix string = ",\"exclusiveMinimum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ExclusiveMinimum))
	}
	if in.ExternalDocs != nil {
		const prefix string = ",\"externalDocs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ExternalDocs).MarshalEasyJSON(out)
	}
	if in.Format != "" {
		const prefix string = ",\"format\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Format))
	}
	if in.ID != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	if (in.Items).IsDefined() {
		const prefix string = ",\"items\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Items).MarshalEasyJSON(out)
	}
	if in.MaxItems != 0 {
		const prefix string = ",\"maxItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MaxItems))
	}
	if in.MaxLength != 0 {
		const prefix string = ",\"maxLength\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MaxLength))
	}
	if in.MaxProperties != 0 {
		const prefix string = ",\"maxProperties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MaxProperties))
	}
	if in.Maximum != 0 {
		const prefix string = ",\"maximum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Maximum))
	}
	if in.MinItems != 0 {
		const prefix string = ",\"minItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MinItems))
	}
	if in.MinLength != 0 {
		const prefix string = ",\"minLength\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MinLength))
	}
	if in.MinProperties != 0 {
		const prefix string = ",\"minProperties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MinProperties))
	}
	if in.Minimum != 0 {
		const prefix string = ",\"minimum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Minimum))
	}
	if in.MultipleOf != 0 {
		const prefix string = ",\"multipleOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.MultipleOf))
	}
	if in.Not != nil {
		const prefix string = ",\"not\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Not).MarshalEasyJSON(out)
	}
	if in.Nullable {
		const prefix string = ",\"nullable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Nullable))
	}
	if len(in.OneOf) != 0 {
		const prefix string = ",\"oneOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.OneOf {
				if v20 > 0 {
					out.RawByte(',')
				}
				if v21 == nil {
					out.RawString("null")
				} else {
					(*v21).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Pattern != "" {
		const prefix string = ",\"pattern\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Pattern))
	}
	if len(in.PatternProperties) != 0 {
		const prefix string = ",\"patternProperties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v22First := true
			for v22Name, v22Value := range in.PatternProperties {
				if v22First {
					v22First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v22Name))
				out.RawByte(':')
				if v22Value == nil {
					out.RawString("null")
				} else {
					(*v22Value).MarshalEasyJSON(out)
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Properties) != 0 {
		const prefix string = ",\"properties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v23First := true
			for v23Name, v23Value := range in.Properties {
				if v23First {
					v23First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v23Name))
				out.RawByte(':')
				if v23Value == nil {
					out.RawString("null")
				} else {
					(*v23Value).MarshalEasyJSON(out)
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Required) != 0 {
		const prefix string = ",\"required\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v24, v25 := range in.Required {
				if v24 > 0 {
					out.RawByte(',')
				}
				out.String(string(v25))
			}
			out.RawByte(']')
		}
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
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
	if in.UniqueItems {
		const prefix string = ",\"uniqueItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.UniqueItems))
	}
	if in.XKubernetesEmbeddedResource {
		const prefix string = ",\"x-kubernetes-embedded-resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.XKubernetesEmbeddedResource))
	}
	if in.XKubernetesIntOrString {
		const prefix string = ",\"x-kubernetes-int-or-string\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.XKubernetesIntOrString))
	}
	if len(in.XKubernetesListMapKeys) != 0 {
		const prefix string = ",\"x-kubernetes-list-map-keys\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v26, v27 := range in.XKubernetesListMapKeys {
				if v26 > 0 {
					out.RawByte(',')
				}
				out.String(string(v27))
			}
			out.RawByte(']')
		}
	}
	if in.XKubernetesListType != "" {
		const prefix string = ",\"x-kubernetes-list-type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.XKubernetesListType))
	}
	if in.XKubernetesMapType != "" {
		const prefix string = ",\"x-kubernetes-map-type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.XKubernetesMapType))
	}
	if in.XKubernetesPreserveUnknownFields {
		const prefix string = ",\"x-kubernetes-preserve-unknown-fields\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.XKubernetesPreserveUnknownFields))
	}
	if len(in.XKubernetesValidations) != 0 {
		const prefix string = ",\"x-kubernetes-validations\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v28, v29 := range in.XKubernetesValidations {
				if v28 > 0 {
					out.RawByte(',')
				}
				if v29 == nil {
					out.RawString("null")
				} else {
					easyjson51a8e721EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV11(out, *v29)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JSONSchemaProps) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson51a8e721EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JSONSchemaProps) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson51a8e721EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JSONSchemaProps) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson51a8e721DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JSONSchemaProps) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson51a8e721DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(l, v)
}
func easyjson51a8e721DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV11(in *jlexer.Lexer, out *ValidationRule) {
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
		case "message":
			out.Message = string(in.String())
		case "rule":
			if in.IsNull() {
				in.Skip()
				out.Rule = nil
			} else {
				if out.Rule == nil {
					out.Rule = new(string)
				}
				*out.Rule = string(in.String())
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
func easyjson51a8e721EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV11(out *jwriter.Writer, in ValidationRule) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"rule\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Rule == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Rule))
		}
	}
	out.RawByte('}')
}
