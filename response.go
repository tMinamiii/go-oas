package test2openapi

// ResponseOrRefObject represents ResponseObject or ReferenceObject
type ResponseOrRefObject struct {
	*ResponseObject
	*ReferenceObject
}

// ResponsesObject
// https://spec.openapis.org/oas/v3.1.0#responses-object
//
// Field Name	Type								Description
// default		Response Object | Reference Object	The documentation of responses other than the ones declared for specific HTTP response codes. Use this field to cover undeclared responses.
type ResponsesObject struct {
	Default ResponseOrRefObject
}

// ResponseObject represents
// https://spec.openapis.org/oas/v3.1.0#response-object
//
// Field Name	Type											Description
// description	string											REQUIRED. A description of the response. CommonMark syntax MAY be used for rich text representation.
// headers		Map[string, Header Object | Reference Object]	Maps a header name to its definition. [RFC7230] states header names are case insensitive. If a response header is defined with the name "Content-Type", it SHALL be ignored.
// content		Map[string, Media Type Object]					A map containing descriptions of potential response payloads. The key is a media type or media type range and the value describes it. For responses that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/*
// links		Map[string, Link Object | Reference Object]		A map of operations links that can be followed from the response. The key of the map is a short name for the link, following the naming constraints of the names for Component Objects.
type ResponseObject struct {
	Description string                       `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]HeaderOrRefObject `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]MediaTypeObject   `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]LinkOrRefObject   `json:"links,omitempty" yaml:"links,omitempty"`
}

// MediaTypeObject represents
// https://spec.openapis.org/oas/v3.1.0#media-type-object
//
// Field Name	Type											Description
// schema		Schema Object									The schema defining the content of the request, response, or parameter.
// example		Any												Example of the media type. The example object SHOULD be in the correct format as specified by the media type. The example field is mutually exclusive of the examples field. Furthermore, if referencing a schema which contains an example, the example value SHALL override the example provided by the schema.
// examples		Map[string, Example Object | Reference Object]	Examples of the media type. Each example object SHOULD match the media type and specified schema if present. The examples field is mutually exclusive of the example field. Furthermore, if referencing a schema which contains an example, the examples value SHALL override the example provided by the schema.
// encoding		Map[string, Encoding Object]					A map between a property name and its encoding information. The key, being the property name, MUST exist in the schema as a property. The encoding object SHALL only apply to requestBody objects when the media type is multipart or application/x-www-form-urlencoded.
type MediaTypeObject struct {
	Schema   SchemaObject                        `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  string                              `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]ExampleObjectOrRefObject `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]EncodingObject           `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// EncodingObject represents
// https://spec.openapis.org/oas/v3.1.0#encoding-object
//
// Field Name	Type											Description
// contentType	string											The Content-Type for encoding a specific property. Default value depends on the property type: for object - application/json; for array – the default is defined based on the inner type; for all other cases the default is application/octet-stream. The value can be a specific media type (e.g. application/json), a wildcard media type (e.g. image/*), or a comma-separated list of the two types.
// headers		Map[string, Header Object | Reference Object]	A map allowing additional information to be provided as headers, for example Content-Disposition. Content-Type is described separately and SHALL be ignored in this section. This property SHALL be ignored if the request body media type is not a multipart.
// style		string											Describes how a specific property value will be serialized depending on its type. See Parameter Object for details on the style property. The behavior follows the same values as query parameters, including default values. This property SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
// explode		boolean											When this is true, property values of type array or object generate separate parameters for each value of the array, or key-value-pair of the map. For other types of properties this property has no effect. When style is form, the default value is true. For all other styles, the default value is false. This property SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
// allowReserved	boolean	Determines whether the parameter value SHOULD allow reserved characters, as defined by [RFC3986] :/?#[]@!$&'()*+,;= to be included without percent-encoding. The default value is false. This property SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
type EncodingObject struct {
	ContentType   string                       `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]HeaderOrRefObject `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string                       `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool                         `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                         `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}
