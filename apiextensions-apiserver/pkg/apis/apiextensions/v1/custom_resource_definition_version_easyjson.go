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

func easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(in *jlexer.Lexer, out *CustomResourceDefinitionVersion) {
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
		case "additionalPrinterColumns":
			if in.IsNull() {
				in.Skip()
				out.AdditionalPrinterColumns = nil
			} else {
				in.Delim('[')
				if out.AdditionalPrinterColumns == nil {
					if !in.IsDelim(']') {
						out.AdditionalPrinterColumns = make([]*CustomResourceColumnDefinition, 0, 8)
					} else {
						out.AdditionalPrinterColumns = []*CustomResourceColumnDefinition{}
					}
				} else {
					out.AdditionalPrinterColumns = (out.AdditionalPrinterColumns)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *CustomResourceColumnDefinition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(CustomResourceColumnDefinition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.AdditionalPrinterColumns = append(out.AdditionalPrinterColumns, v1)
					in.WantComma()
				}
				in.Delim(']')
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
		case "schema":
			if in.IsNull() {
				in.Skip()
				out.Schema = nil
			} else {
				if out.Schema == nil {
					out.Schema = new(CustomResourceValidation)
				}
				easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV11(in, out.Schema)
			}
		case "served":
			if in.IsNull() {
				in.Skip()
				out.Served = nil
			} else {
				if out.Served == nil {
					out.Served = new(bool)
				}
				*out.Served = bool(in.Bool())
			}
		case "storage":
			if in.IsNull() {
				in.Skip()
				out.Storage = nil
			} else {
				if out.Storage == nil {
					out.Storage = new(bool)
				}
				*out.Storage = bool(in.Bool())
			}
		case "subresources":
			if in.IsNull() {
				in.Skip()
				out.Subresources = nil
			} else {
				if out.Subresources == nil {
					out.Subresources = new(CustomResourceSubresources)
				}
				easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV12(in, out.Subresources)
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
func easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(out *jwriter.Writer, in CustomResourceDefinitionVersion) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.AdditionalPrinterColumns) != 0 {
		const prefix string = ",\"additionalPrinterColumns\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.AdditionalPrinterColumns {
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
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	if in.Schema != nil {
		const prefix string = ",\"schema\":"
		out.RawString(prefix)
		easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV11(out, *in.Schema)
	}
	{
		const prefix string = ",\"served\":"
		out.RawString(prefix)
		if in.Served == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Served))
		}
	}
	{
		const prefix string = ",\"storage\":"
		out.RawString(prefix)
		if in.Storage == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Storage))
		}
	}
	if in.Subresources != nil {
		const prefix string = ",\"subresources\":"
		out.RawString(prefix)
		easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV12(out, *in.Subresources)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CustomResourceDefinitionVersion) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CustomResourceDefinitionVersion) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CustomResourceDefinitionVersion) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CustomResourceDefinitionVersion) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(l, v)
}
func easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV12(in *jlexer.Lexer, out *CustomResourceSubresources) {
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
		case "scale":
			if in.IsNull() {
				in.Skip()
				out.Scale = nil
			} else {
				if out.Scale == nil {
					out.Scale = new(CustomResourceSubresourceScale)
				}
				easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV13(in, out.Scale)
			}
		case "status":
			(out.Status).UnmarshalEasyJSON(in)
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
func easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV12(out *jwriter.Writer, in CustomResourceSubresources) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Scale != nil {
		const prefix string = ",\"scale\":"
		first = false
		out.RawString(prefix[1:])
		easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV13(out, *in.Scale)
	}
	if (in.Status).IsDefined() {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Status).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV13(in *jlexer.Lexer, out *CustomResourceSubresourceScale) {
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
		case "labelSelectorPath":
			out.LabelSelectorPath = string(in.String())
		case "specReplicasPath":
			if in.IsNull() {
				in.Skip()
				out.SpecReplicasPath = nil
			} else {
				if out.SpecReplicasPath == nil {
					out.SpecReplicasPath = new(string)
				}
				*out.SpecReplicasPath = string(in.String())
			}
		case "statusReplicasPath":
			if in.IsNull() {
				in.Skip()
				out.StatusReplicasPath = nil
			} else {
				if out.StatusReplicasPath == nil {
					out.StatusReplicasPath = new(string)
				}
				*out.StatusReplicasPath = string(in.String())
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
func easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV13(out *jwriter.Writer, in CustomResourceSubresourceScale) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LabelSelectorPath != "" {
		const prefix string = ",\"labelSelectorPath\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.LabelSelectorPath))
	}
	{
		const prefix string = ",\"specReplicasPath\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.SpecReplicasPath == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.SpecReplicasPath))
		}
	}
	{
		const prefix string = ",\"statusReplicasPath\":"
		out.RawString(prefix)
		if in.StatusReplicasPath == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.StatusReplicasPath))
		}
	}
	out.RawByte('}')
}
func easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV11(in *jlexer.Lexer, out *CustomResourceValidation) {
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
		case "openAPIV3Schema":
			if in.IsNull() {
				in.Skip()
				out.OpenAPIV3Schema = nil
			} else {
				if out.OpenAPIV3Schema == nil {
					out.OpenAPIV3Schema = new(JSONSchemaProps)
				}
				easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(in, out.OpenAPIV3Schema)
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
func easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV11(out *jwriter.Writer, in CustomResourceValidation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.OpenAPIV3Schema != nil {
		const prefix string = ",\"openAPIV3Schema\":"
		first = false
		out.RawString(prefix[1:])
		easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(out, *in.OpenAPIV3Schema)
	}
	out.RawByte('}')
}
func easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(in *jlexer.Lexer, out *JSONSchemaProps) {
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
					var v4 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(JSONSchemaProps)
						}
						easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(in, v4)
					}
					out.AllOf = append(out.AllOf, v4)
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
					var v5 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(JSONSchemaProps)
						}
						easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(in, v5)
					}
					out.AnyOf = append(out.AnyOf, v5)
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
					var v6 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(JSONSchemaProps)
						}
						easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(in, v6)
					}
					(out.Definitions)[key] = v6
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
					var v7 easyjson.RawMessage
					(v7).UnmarshalEasyJSON(in)
					(out.Dependencies)[key] = v7
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
					var v8 easyjson.RawMessage
					(v8).UnmarshalEasyJSON(in)
					out.Enum = append(out.Enum, v8)
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
				easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV15(in, out.ExternalDocs)
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
				easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(in, out.Not)
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
					var v9 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(JSONSchemaProps)
						}
						easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(in, v9)
					}
					out.OneOf = append(out.OneOf, v9)
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
					var v10 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(JSONSchemaProps)
						}
						easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(in, v10)
					}
					(out.PatternProperties)[key] = v10
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
					var v11 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(JSONSchemaProps)
						}
						easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(in, v11)
					}
					(out.Properties)[key] = v11
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
					var v12 string
					v12 = string(in.String())
					out.Required = append(out.Required, v12)
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
					var v13 string
					v13 = string(in.String())
					out.XKubernetesListMapKeys = append(out.XKubernetesListMapKeys, v13)
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
func easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(out *jwriter.Writer, in JSONSchemaProps) {
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
			for v14, v15 := range in.AllOf {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(out, *v15)
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
			for v16, v17 := range in.AnyOf {
				if v16 > 0 {
					out.RawByte(',')
				}
				if v17 == nil {
					out.RawString("null")
				} else {
					easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(out, *v17)
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
			v18First := true
			for v18Name, v18Value := range in.Definitions {
				if v18First {
					v18First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v18Name))
				out.RawByte(':')
				if v18Value == nil {
					out.RawString("null")
				} else {
					easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(out, *v18Value)
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
			v19First := true
			for v19Name, v19Value := range in.Dependencies {
				if v19First {
					v19First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v19Name))
				out.RawByte(':')
				(v19Value).MarshalEasyJSON(out)
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
			for v20, v21 := range in.Enum {
				if v20 > 0 {
					out.RawByte(',')
				}
				(v21).MarshalEasyJSON(out)
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
		easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV15(out, *in.ExternalDocs)
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
		easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(out, *in.Not)
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
			for v22, v23 := range in.OneOf {
				if v22 > 0 {
					out.RawByte(',')
				}
				if v23 == nil {
					out.RawString("null")
				} else {
					easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(out, *v23)
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
			v24First := true
			for v24Name, v24Value := range in.PatternProperties {
				if v24First {
					v24First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v24Name))
				out.RawByte(':')
				if v24Value == nil {
					out.RawString("null")
				} else {
					easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(out, *v24Value)
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
			v25First := true
			for v25Name, v25Value := range in.Properties {
				if v25First {
					v25First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v25Name))
				out.RawByte(':')
				if v25Value == nil {
					out.RawString("null")
				} else {
					easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV14(out, *v25Value)
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
			for v26, v27 := range in.Required {
				if v26 > 0 {
					out.RawByte(',')
				}
				out.String(string(v27))
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
			for v28, v29 := range in.XKubernetesListMapKeys {
				if v28 > 0 {
					out.RawByte(',')
				}
				out.String(string(v29))
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
	out.RawByte('}')
}
func easyjson81183069DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV15(in *jlexer.Lexer, out *ExternalDocumentation) {
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
		case "description":
			out.Description = string(in.String())
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
func easyjson81183069EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV15(out *jwriter.Writer, in ExternalDocumentation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Description != "" {
		const prefix string = ",\"description\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Description))
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
