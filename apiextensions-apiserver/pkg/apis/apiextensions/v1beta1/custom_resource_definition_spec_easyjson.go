// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

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

func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(in *jlexer.Lexer, out *CustomResourceDefinitionSpec) {
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
		case "conversion":
			if in.IsNull() {
				in.Skip()
				out.Conversion = nil
			} else {
				if out.Conversion == nil {
					out.Conversion = new(CustomResourceConversion)
				}
				(*out.Conversion).UnmarshalEasyJSON(in)
			}
		case "group":
			if in.IsNull() {
				in.Skip()
				out.Group = nil
			} else {
				if out.Group == nil {
					out.Group = new(string)
				}
				*out.Group = string(in.String())
			}
		case "names":
			if in.IsNull() {
				in.Skip()
				out.Names = nil
			} else {
				if out.Names == nil {
					out.Names = new(CustomResourceDefinitionNames)
				}
				(*out.Names).UnmarshalEasyJSON(in)
			}
		case "scope":
			if in.IsNull() {
				in.Skip()
				out.Scope = nil
			} else {
				if out.Scope == nil {
					out.Scope = new(string)
				}
				*out.Scope = string(in.String())
			}
		case "subresources":
			if in.IsNull() {
				in.Skip()
				out.Subresources = nil
			} else {
				if out.Subresources == nil {
					out.Subresources = new(CustomResourceSubresources)
				}
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, out.Subresources)
			}
		case "validation":
			if in.IsNull() {
				in.Skip()
				out.Validation = nil
			} else {
				if out.Validation == nil {
					out.Validation = new(CustomResourceValidation)
				}
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(in, out.Validation)
			}
		case "version":
			out.Version = string(in.String())
		case "versions":
			if in.IsNull() {
				in.Skip()
				out.Versions = nil
			} else {
				in.Delim('[')
				if out.Versions == nil {
					if !in.IsDelim(']') {
						out.Versions = make([]*CustomResourceDefinitionVersion, 0, 8)
					} else {
						out.Versions = []*CustomResourceDefinitionVersion{}
					}
				} else {
					out.Versions = (out.Versions)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *CustomResourceDefinitionVersion
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(CustomResourceDefinitionVersion)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta13(in, v2)
					}
					out.Versions = append(out.Versions, v2)
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(out *jwriter.Writer, in CustomResourceDefinitionSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.AdditionalPrinterColumns) != 0 {
		const prefix string = ",\"additionalPrinterColumns\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v3, v4 := range in.AdditionalPrinterColumns {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					(*v4).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Conversion != nil {
		const prefix string = ",\"conversion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Conversion).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"group\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Group == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Group))
		}
	}
	{
		const prefix string = ",\"names\":"
		out.RawString(prefix)
		if in.Names == nil {
			out.RawString("null")
		} else {
			(*in.Names).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"scope\":"
		out.RawString(prefix)
		if in.Scope == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Scope))
		}
	}
	if in.Subresources != nil {
		const prefix string = ",\"subresources\":"
		out.RawString(prefix)
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *in.Subresources)
	}
	if in.Validation != nil {
		const prefix string = ",\"validation\":"
		out.RawString(prefix)
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(out, *in.Validation)
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	if len(in.Versions) != 0 {
		const prefix string = ",\"versions\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Versions {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta13(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CustomResourceDefinitionSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CustomResourceDefinitionSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CustomResourceDefinitionSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CustomResourceDefinitionSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(l, v)
}
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta13(in *jlexer.Lexer, out *CustomResourceDefinitionVersion) {
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
					var v7 *CustomResourceColumnDefinition
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(CustomResourceColumnDefinition)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.AdditionalPrinterColumns = append(out.AdditionalPrinterColumns, v7)
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
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(in, out.Schema)
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
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, out.Subresources)
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta13(out *jwriter.Writer, in CustomResourceDefinitionVersion) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.AdditionalPrinterColumns) != 0 {
		const prefix string = ",\"additionalPrinterColumns\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v8, v9 := range in.AdditionalPrinterColumns {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
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
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(out, *in.Schema)
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
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *in.Subresources)
	}
	out.RawByte('}')
}
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(in *jlexer.Lexer, out *CustomResourceValidation) {
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
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, out.OpenAPIV3Schema)
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(out *jwriter.Writer, in CustomResourceValidation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.OpenAPIV3Schema != nil {
		const prefix string = ",\"openAPIV3Schema\":"
		first = false
		out.RawString(prefix[1:])
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *in.OpenAPIV3Schema)
	}
	out.RawByte('}')
}
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in *jlexer.Lexer, out *JSONSchemaProps) {
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
					var v10 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(JSONSchemaProps)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, v10)
					}
					out.AllOf = append(out.AllOf, v10)
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
					var v11 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(JSONSchemaProps)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, v11)
					}
					out.AnyOf = append(out.AnyOf, v11)
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
					var v12 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v12 = nil
					} else {
						if v12 == nil {
							v12 = new(JSONSchemaProps)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, v12)
					}
					(out.Definitions)[key] = v12
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
					var v13 easyjson.RawMessage
					(v13).UnmarshalEasyJSON(in)
					(out.Dependencies)[key] = v13
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
					var v14 easyjson.RawMessage
					(v14).UnmarshalEasyJSON(in)
					out.Enum = append(out.Enum, v14)
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
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta15(in, out.ExternalDocs)
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
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, out.Not)
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
					var v15 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v15 = nil
					} else {
						if v15 == nil {
							v15 = new(JSONSchemaProps)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, v15)
					}
					out.OneOf = append(out.OneOf, v15)
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
					var v16 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v16 = nil
					} else {
						if v16 == nil {
							v16 = new(JSONSchemaProps)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, v16)
					}
					(out.PatternProperties)[key] = v16
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
					var v17 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v17 = nil
					} else {
						if v17 == nil {
							v17 = new(JSONSchemaProps)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, v17)
					}
					(out.Properties)[key] = v17
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
					var v18 string
					v18 = string(in.String())
					out.Required = append(out.Required, v18)
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out *jwriter.Writer, in JSONSchemaProps) {
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
			for v19, v20 := range in.AllOf {
				if v19 > 0 {
					out.RawByte(',')
				}
				if v20 == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *v20)
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
			for v21, v22 := range in.AnyOf {
				if v21 > 0 {
					out.RawByte(',')
				}
				if v22 == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *v22)
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
			v23First := true
			for v23Name, v23Value := range in.Definitions {
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
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *v23Value)
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
			v24First := true
			for v24Name, v24Value := range in.Dependencies {
				if v24First {
					v24First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v24Name))
				out.RawByte(':')
				(v24Value).MarshalEasyJSON(out)
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
			for v25, v26 := range in.Enum {
				if v25 > 0 {
					out.RawByte(',')
				}
				(v26).MarshalEasyJSON(out)
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
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta15(out, *in.ExternalDocs)
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
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *in.Not)
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
			for v27, v28 := range in.OneOf {
				if v27 > 0 {
					out.RawByte(',')
				}
				if v28 == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *v28)
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
			v29First := true
			for v29Name, v29Value := range in.PatternProperties {
				if v29First {
					v29First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v29Name))
				out.RawByte(':')
				if v29Value == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *v29Value)
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
			v30First := true
			for v30Name, v30Value := range in.Properties {
				if v30First {
					v30First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v30Name))
				out.RawByte(':')
				if v30Value == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *v30Value)
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
			for v31, v32 := range in.Required {
				if v31 > 0 {
					out.RawByte(',')
				}
				out.String(string(v32))
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
	out.RawByte('}')
}
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta15(in *jlexer.Lexer, out *ExternalDocumentation) {
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta15(out *jwriter.Writer, in ExternalDocumentation) {
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
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in *jlexer.Lexer, out *CustomResourceSubresources) {
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
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta16(in, out.Scale)
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out *jwriter.Writer, in CustomResourceSubresources) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Scale != nil {
		const prefix string = ",\"scale\":"
		first = false
		out.RawString(prefix[1:])
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta16(out, *in.Scale)
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
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta16(in *jlexer.Lexer, out *CustomResourceSubresourceScale) {
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta16(out *jwriter.Writer, in CustomResourceSubresourceScale) {
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
