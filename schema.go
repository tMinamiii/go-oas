package test2openapi

// OpenAPI represents
// https://spec.openapis.org/oas/v3.1.0#openapi-object
//
// Field Name			Type							Description
// openapi				string							REQUIRED. This string MUST be the version number of the OpenAPI Specification that the OpenAPI document uses. The openapi field SHOULD be used by tooling to interpret the OpenAPI document. This is not related to the API info.version string.
// info	Info 			Object							REQUIRED. Provides metadata about the API. The metadata MAY be used by tooling as required.
// jsonSchemaDialect	string							The default value for the $schema keyword within Schema Objects contained within this OAS document. This MUST be in the form of a URI.
// servers				[Server Object]					An array of Server Objects, which provide connectivity information to a target server. If the servers property is not provided, or is an empty array, the default value would be a Server Object with a url value of /.
// paths				Paths Object					The available paths and operations for the API.
// webhooks				Map[string, Path Item Object | Reference Object] ]	The incoming webhooks that MAY be received as part of this API and that the API consumer MAY choose to implement. Closely related to the callbacks feature, this section describes requests initiated other than by an API call, for example by an out of band registration. The key name is a unique string to refer to each webhook, while the (optionally referenced) Path Item Object describes a request that may be initiated by the API provider and the expected responses. An example is available.
// components			Components Object				An element to hold various schemas for the document.
// security				[Security Requirement Object]	A declaration of which security mechanisms can be used across the API. The list of values includes alternative security requirement objects that can be used. Only one of the security requirement objects need to be satisfied to authorize a request. Individual operations can override this definition. To make security optional, an empty security requirement ({}) can be included in the array.
// tags					[Tag Object]					A list of tags used by the document with additional metadata. The order of the tags can be used to reflect on their order by the parsing tools. Not all tags that are used by the Operation Object must be declared. The tags that are not declared MAY be organized randomly or based on the tools’ logic. Each tag name in the list MUST be unique.
// externalDocs			External Documentation Object	Additional external documentation.
type OpenAPI struct {
	OpenAPIVersion    string                      `json:"openapi" yaml:"openapi"`
	Info              InfoObject                  `json:"info" yaml:"info"`
	JSONSchemaDialect string                      `json:"jsonSchemaDialect,omitempty" yaml:"jsonSchemaDialect,omitempty"`
	Servers           []ServerObject              `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths             PathsObject                 `json:"paths,omitempty" yaml:"paths,omitempty"`
	Webhooks          []PathItemOrRefObject       `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
	Components        ComponentsObject            `json:"components,omitempty" yaml:"components,omitempty"`
	Security          []SecurityRequirementObject `json:"security,omitempty" yaml:"security,omitempty"`
	Tags              []TagObject                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs      ExternalDocObject           `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// InfoObject represents
// https://spec.openapis.org/oas/v3.1.0#info-object
//
// Field Name		Type			Description
// title			string			REQUIRED. The title of the API.
// summary			string			A short summary of the API.
// description		string			A description of the API. CommonMark syntax MAY be used for rich text representation.
// termsOfService	string			A URL to the Terms of Service for the API. This MUST be in the form of a URL.
// contact			Contact Object	The contact information for the exposed API.
// license			License Object	The license information for the exposed API.
// version			string			REQUIRED. The version of the OpenAPI document (which is distinct from the OpenAPI Specification version or the API implementation version).
type InfoObject struct {
	Title          string        `json:"title" yaml:"title"`
	Summary        string        `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description    string        `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string        `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        ContactObject `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        LicenseObject `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string        `json:"version" yaml:"version"`
}

// ServerObject represents
// https://spec.openapis.org/oas/v3.1.0#server-object
//
// Field Name	Type								Description
// url			string								REQUIRED. A URL to the target host. This URL supports ServerObject Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenAPI document is being served. Variable substitutions will be made when a variable is named in {brackets}.
// description	string								An optional string describing the host designated by the URL. CommonMark syntax MAY be used for rich text representation.
// variables	Map[string, ServerObject Variable Object]	A map between a variable name and its value. The value is used for substitution in the server’s URL template.
type ServerObject struct {
	URL         string                          `json:"url" yaml:"url"`                                     // REQUIRED
	Description string                          `json:"description,omitempty" yaml:"description,omitempty"` // NULLABLE
	Variables   map[string]ServerVariableObject `json:"variable,omitempty" yaml:"variable,omitempty"`
}

// ServerVariableObject represents
// https://spec.openapis.org/oas/v3.1.0#server-variable-object
//
// Field Name	Type		Description
// enum			[string]	An enumeration of string values to be used if the substitution options are from a limited set. The array MUST NOT be empty.
// default		string		REQUIRED. The default value to use for substitution, which SHALL be sent if an alternate value is not supplied. Note this behavior is different than the Schema Object’s treatment of default values, because in those cases parameter values are optional. If the enum is defined, the value MUST exist in the enum’s values.
// description	string		An optional description for the server variable. CommonMark syntax MAY be used for rich text representation.
type ServerVariableObject struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`               // MUST NOT be empty
	Default     string   `json:"default" yaml:"default"`                             // Required
	Description string   `json:"description,omitempty" yaml:"description,omitempty"` // NULLABLE
}

// PathObject represents
// https://spec.openapis.org/oas/v3.1.0#paths-object
//
// Field Pattern	Type				Description
// /{path}			Path Item Object	A relative path to an individual endpoint. The field name MUST begin with a forward slash (/). The path is appended (no relative URL resolution) to the expanded URL from the Server Object’s url field in order to construct the full URL. Path templating is allowed. When matching URLs, concrete (non-templated) paths would be matched before their templated counterparts. Templated paths with the same hierarchy but different templated names MUST NOT exist as they are identical. In case of ambiguous matching, it’s up to the tooling to decide which one to use.
type PathsObject map[string]PathItemObject

// PathItemOrRefObject represents PathItemObject or ReferenceObject
type PathItemOrRefObject struct {
	*PathItemObject
	*ReferenceObject
}

// PathItem represent
// https://spec.openapis.org/oas/v3.1.0#path-item-object
// Field Name 	Type				Description,
// $ref			string				Allows for a referenced definition of this path item. The referenced structure MUST be in the form of a Path Item Object. In case a Path Item Object field appears both in the defined object and the referenced object, the behavior is undefined. See the rules for resolving Relative References.
// summary		string				An optional, string summary, intended to apply to all operations in this path.
// description	string				An optional, string description, intended to apply to all operations in this path. CommonMark syntax MAY be used for rich text representation.
// get			Operation Object	A definition of a GET operation on this path.
// put			Operation Object	A definition of a PUT operation on this path.
// post			Operation Object	A definition of a POST operation on this path.
// delete		Operation Object	A definition of a DELETE operation on this path.
// options		Operation Object	A definition of a OPTIONS operation on this path.
// head			Operation Object	A definition of a HEAD operation on this path.
// patch		Operation Object	A definition of a PATCH operation on this path.
// trace		Operation Object	A definition of a TRACE operation on this path.
// servers		[Server Object]		An alternative server array to service all operations in this path.
// parameters	[Parameter Object | Reference Object]	A list of parameters that are applicable for all the operations described under this path. These parameters can be overridden at the operation level, but cannot be removed there. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined at the OpenAPI Object’s components/parameters.
type PathItemObject struct {
	Ref         string                 `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Get         OperationObject        `json:"get,omitempty" yaml:"get,omitempty"`
	Put         OperationObject        `json:"put,omitempty" yaml:"put,omitempty"`
	Post        OperationObject        `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      OperationObject        `json:"delete,omitempty" yaml:"delete,omitempty"`
	Patch       OperationObject        `json:"patch,omitempty" yaml:"patch,omitempty"`
	Options     OperationObject        `json:"options,omitempty" yaml:"options,omitempty"`
	Head        OperationObject        `json:"head,omitempty" yaml:"head,omitempty"`
	Trace       OperationObject        `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []ServerObject         `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []ParameterOrRefObject `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// OperationObject represents
// https://spec.openapis.org/oas/v3.1.0#operation-object
//
// Field Name	Type									Description
// tags			[string]								A list of tags for API documentation control. Tags can be used for logical grouping of operations by resources or any other qualifier.
// summary		string									A short summary of what the operation does.
// description	string									A verbose explanation of the operation behavior. CommonMark syntax MAY be used for rich text representation.
// externalDocs	External Documentation Object			Additional external documentation for this operation.
// operationId	string									Unique string used to identify the operation. The id MUST be unique among all operations described in the API. The operationId value is case-sensitive. Tools and libraries MAY use the operationId to uniquely identify an operation, therefore, it is RECOMMENDED to follow common programming naming conventions.
// parameters	[Parameter Object | Reference Object]	A list of parameters that are applicable for this operation. If a parameter is already defined at the Path Item, the new definition will override it but can never remove it. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined at the OpenAPI Object’s components/parameters.
// requestBody	Request Body Object | Reference Object	The request body applicable for this operation. The requestBody is fully supported in HTTP methods where the HTTP 1.1 specification [RFC7231] has explicitly defined semantics for request bodies. In other cases where the HTTP spec is vague (such as GET, HEAD and DELETE), requestBody is permitted but does not have well-defined semantics and SHOULD be avoided if possible.
// responses	Responses Object						The list of possible responses as they are returned from executing this operation.
// callbacks	Map[string, Callback Object | Reference Object]	A map of possible out-of band callbacks related to the parent operation. The key is a unique identifier for the Callback Object. Each value in the map is a Callback Object that describes a request that may be initiated by the API provider and the expected responses.
// deprecated	boolean									Declares this operation to be deprecated. Consumers SHOULD refrain from usage of the declared operation. Default value is false.
// security		[Security Requirement Object]			A declaration of which security mechanisms can be used for this operation. The list of values includes alternative security requirement objects that can be used. Only one of the security requirement objects need to be satisfied to authorize a request. To make security optional, an empty security requirement ({}) can be included in the array. This definition overrides any declared top-level security. To remove a top-level security declaration, an empty array can be used.
// servers		[Server Object]							An alternative server array to service this operation. If an alternative server object is specified at the Path Item Object or Root level, it will be overridden by this value.
type OperationObject struct {
	Tags         []string               `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs ExternalDocObject      `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []ParameterOrRefObject `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  RequestBodyObject      `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    ResponseOrRefObject    `json:"responses,omitempty" yaml:"responses,omitempty"`
	Callbacks    []CallbackOrRefObject  `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []SecuritySchemeObject `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []ServerObject         `json:"servers,omitempty" yaml:"servers,omitempty"`
}

// ParameterOrRefObject represents ParameterObject or ReferenceObject
type ParameterOrRefObject struct {
	*ParameterObject
	*ReferenceObject
}

// ParameterObject represents
// https://spec.openapis.org/oas/v3.1.0#parameter-object
//
// Field Name		Type	Description
// name				string	REQUIRED. The name of the parameter. Parameter names are case sensitive.  If in is "path", the name field MUST correspond to a template expression occurring within the path field in the Paths Object. See Path Templating for further information.  If in is "header" and the name field is "Accept", "Content-Type" or "Authorization", the parameter definition SHALL be ignored.  For all other cases, the name corresponds to the parameter name used by the in property.
// in				string	REQUIRED. The location of the parameter. Possible values are "query", "header", "path" or "cookie".
// description		string	A brief description of the parameter. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
// required			boolean	Determines whether this parameter is mandatory. If the parameter location is "path", this property is REQUIRED and its value MUST be true. Otherwise, the property MAY be included and its default value is false.
// deprecated		boolean	Specifies that a parameter is deprecated and SHOULD be transitioned out of usage. Default value is false.
// allowEmptyValue	boolean	Sets the ability to pass empty-valued parameters. This is valid only for query parameters and allows sending a parameter with an empty value. Default value is false. If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored. Use of this property is NOT RECOMMENDED, as it is likely to be removed in a later revision.
type ParameterObject struct {
	Name            string `json:"name" yaml:"name"`
	In              string `json:"in" yaml:"in"`
	Description     string `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool   `json:"required" yaml:"required"`
	Deprecated      bool   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
}

// ComponentsObject represents
// https://spec.openapis.org/oas/v3.1.0#components-object
//
// Field Name		Type												Description
// schemas			Map[string, Schema Object]							An object to hold reusable Schema Objects.
// responses		Map[string, Response Object | Reference Object]		An object to hold reusable Response Objects.
// parameters		Map[string, Parameter Object | Reference Object]	An object to hold reusable Parameter Objects.
// examples			Map[string, Example Object | Reference Object]		An object to hold reusable Example Objects.
// requestBodies	Map[string, Request Body Object | Reference Object]	An object to hold reusable Request Body Objects.
// headers			Map[string, Header Object | Reference Object]		An object to hold reusable Header Objects.
// securitySchemes	Map[string, Security Scheme Object | Reference Object]	An object to hold reusable Security Scheme Objects.
// links			Map[string, Link Object | Reference Object]			An object to hold reusable Link Objects.
// callbacks		Map[string, Callback Object | Reference Object]		An object to hold reusable Callback Objects.
// pathItems		Map[string, Path Item Object | Reference Object]	An object to hold reusable Path Item Object.
type ComponentsObject struct {
	Schemas         map[string]SchemaObject                    `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]ResponseOrRefObject             `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]ParameterOrRefObject            `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]ExampleObjectOrRefObject        `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]RequestBodyOrRefObject          `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]HeaderOrRefObject               `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes map[string]SecuritySchemeObjectOrRefObject `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           map[string]LinkOrRefObject                 `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]CallbackOrRefObject             `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	PathItems       map[string]PathItemOrRefObject             `json:"pathItems,omitempty" yaml:"pathItems,omitempty"`
}

// SchemaObject represents
// https://spec.openapis.org/oas/v3.1.0#schema-object
//
// Field Name		Type							Description
// discriminator	Discriminator Object			Adds support for polymorphism. The discriminator is an object name that is used to differentiate between other schemas which may satisfy the payload description. See Composition and Inheritance for more details.
// xml				XML Object						This MAY be used only on properties schemas. It has no effect on root schemas. Adds additional metadata to describe the XML representation of this property.
// externalDocs		External Documentation Object	Additional external documentation for this schema.
// example			Any								A free-form property to include an example of an instance for this schema. To represent examples that cannot be naturally represented in JSON or YAML, a string value can be used to contain the example with escaping where necessary.  Deprecated: The example property has been deprecated in favor of the JSON SchemaObject examples keyword. Use of example is discouraged, and later versions of this specification may remove it.
type SchemaObject struct {
	Discriminator DiscriminatorObject `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	XML           XMLObject           `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExternalDocs  ExternalDocObject   `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Example       bool                `json:"example,omitempty" yaml:"example,omitempty"`
}

// DiscriminatorObject represents
// https://spec.openapis.org/oas/v3.1.0#discriminator-object
//
// Field Name	Type				Description
// propertyName	string				REQUIRED. The name of the property in the payload that will hold the discriminator value.
// mapping		Map[string, string]	An object to hold mappings between payload values and schema names or references.
type DiscriminatorObject struct {
	PropertyName string            `json:"propertyName" yaml:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

// XMLObject represents
// https://spec.openapis.org/oas/v3.1.0#xml-object
//
// Field Name	Type	Description
// name			string	Replaces the name of the element/attribute used for the described schema property. When defined within items, it will affect the name of the individual XML elements within the list. When defined alongside type being array (outside the items), it will affect the wrapping element and only if wrapped is true. If wrapped is false, it will be ignored.
// namespace	string	The URI of the namespace definition. This MUST be in the form of an absolute URI.
// prefix		string	The prefix to be used for the name.
// attribute	boolean	Declares whether the property definition translates to an attribute instead of an element. Default value is false.
// wrapped		boolean	MAY be used only for an array definition. Signifies whether the array is wrapped (for example, <books><book/><book/></books>) or unwrapped (<book/><book/>). Default value is false. The definition takes effect only when defined alongside type being array (outside the items).
type XMLObject struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}

// RequestBodyOrRefObject represents RequestBodyObject or ReferenceObject
type RequestBodyOrRefObject struct {
	*RequestBodyObject
	*ReferenceObject
}

// HeaderOrRefObject represents HeaderObject or ReferenceObject
type HeaderOrRefObject struct {
	*HeaderObject
	*ReferenceObject
}

// HeaderObject represents
// https://spec.openapis.org/oas/v3.1.0#header-object
//
// Field Name	Type							Description
// name			string							REQUIRED. The name of the tag.
// description	string							A description for the tag. CommonMark syntax MAY be used for rich text representation.
// externalDocs	External Documentation Object	Additional external documentation for this tag.
type HeaderObject struct {
	Name         string            `json:"name" yaml:"name"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs ExternalDocObject `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// RequestBodyObject represents
// https://spec.openapis.org/oas/v3.1.0#request-body-object
//
// Field Name	Type							Description
// description	string							A brief description of the request body. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
// content		Map[string, Media Type Object]	REQUIRED. The content of the request body. The key is a media type or media type range and the value describes it. For requests that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/*
// required		boolean							Determines if the request body is required in the request. Defaults to false.
type RequestBodyObject struct {
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]MediaTypeObject `json:"content" yaml:"content"`
	Required    bool                       `json:"required,omitempty" yaml:"required,omitempty"`
}

// CallbackOrRefObject represents CallbackObject or ReferenceObject
type CallbackOrRefObject struct {
	*CallbackObject
	*ReferenceObject
}

// CallbackObject represents
// https://spec.openapis.org/oas/v3.1.0#callback-object
//
// Field Pattern	Type									Description
// {expression}		Path Item Object | Reference Object		A Path Item Object, or a reference to one, used to define a callback request and expected responses. A complete example is available.
//
// Expression					Value
// $url							https://example.org/subscribe/myevent?queryUrl=https://clientdomain.com/stillrunning
// $method						POST
// $request.path.eventType		myevent
// $request.query.queryUrl		https://clientdomain.com/stillrunning
// $request.header.content-Type	application/json
// $request.body#/failedUrl		https://clientdomain.com/failed
// $request.body#/successUrls/2	https://clientdomain.com/medium
// $response.header.Location	https://example.org/subscription/1
type CallbackObject map[string]PathItemOrRefObject

// ExampleObjectOrRefObject represents ExampleObject or ReferenceObject
type ExampleObjectOrRefObject struct {
	*ExampleObject
	*ReferenceObject
}

// ExampleObject represents
// https://spec.openapis.org/oas/v3.1.0#example-object
//
// Field Name		Type	Description
// summary			string	Short description for the example.
// description		string	Long description for the example. CommonMark syntax MAY be used for rich text representation.
// value			Any		Embedded literal example. The value field and externalValue field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON or YAML, use a string value to contain the example, escaping where necessary.
// externalValue	string	A URI that points to the literal example. This provides the capability to reference examples that cannot easily be included in JSON or YAML documents. The value field and externalValue field are mutually exclusive. See the rules for resolving Relative References.
type ExampleObject struct {
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]MediaTypeObject `json:"content,omitempty" yaml:"content,omitempty"`
	Required    bool                       `json:"required,omitempty" yaml:"required,omitempty"`
}

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

// LinkOrRefObject represents LinkObject or ReferenceObject
type LinkOrRefObject struct {
	*LinkObject
	*ReferenceObject
}

// LinkObject represents
// https://spec.openapis.org/oas/v3.1.0#link-object
//
// Field Name	Type					Description
// operationRef	string					A relative or absolute URI reference to an OAS operation. This field is mutually exclusive of the operationId field, and MUST point to an Operation Object. Relative operationRef values MAY be used to locate an existing Operation Object in the OpenAPI definition. See the rules for resolving Relative References.
// operationId	string					The name of an existing, resolvable OAS operation, as defined with a unique operationId. This field is mutually exclusive of the operationRef field.
// parameters	Map[string, Any | {expression}]	A map representing parameters to pass to an operation as specified with operationId or identified via operationRef. The key is the parameter name to be used, whereas the value can be a constant or an expression to be evaluated and passed to the linked operation. The parameter name can be qualified using the parameter location [{in}.]{name} for operations that use the same parameter name in different locations (e.g. path.id).
// requestBody	Any | {expression}		A literal value or {expression} to use as a request body when calling the target operation.
// description	string					A description of the link. CommonMark syntax MAY be used for rich text representation.
// server		Server Object			A server object to be used by the target operation.
type LinkObject struct {
	OperationRef string            `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationID  string            `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  string            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty"`
	Server       ServerObject      `json:"server,omitempty" yaml:"server,omitempty"`
}

// SecuritySchemeObjectOrRefObject represents SecuritySchemeObject or ReferenceObject
type SecuritySchemeObjectOrRefObject struct {
	*SecuritySchemeObject
	*ReferenceObject
}

// SecuritySchemeObject represents
// https://spec.openapis.org/oas/v3.1.0#security-scheme-object
//
// Field Name		Type				Applies To		Description
// type				string				Any				REQUIRED. The type of the security scheme. Valid values are "apiKey", "http", "mutualTLS", "oauth2", "openIdConnect".
// description		string				Any				A description for security scheme. CommonMark syntax MAY be used for rich text representation.
// name				string				apiKey			REQUIRED. The name of the header, query or cookie parameter to be used.
// in				string				apiKey			REQUIRED. The location of the API key. Valid values are "query", "header" or "cookie".
// scheme			string				http			REQUIRED. The name of the HTTP Authorization scheme to be used in the Authorization header as defined in [RFC7235]. The values used SHOULD be registered in the IANA Authentication Scheme registry.
// bearerFormat		string				http ("bearer")	A hint to the client to identify how the bearer token is formatted. Bearer tokens are usually generated by an authorization server, so this information is primarily for documentation purposes.
// flows			OAuth Flows Object	oauth2			REQUIRED. An object containing configuration information for the flow types supported.
// openIdConnectUrl	string				openIdConnect	REQUIRED. OpenId Connect URL to discover OAuth2 configuration values. This MUST be in the form of a URL. The OpenID Connect standard requires the use of TLS.
type SecuritySchemeObject struct {
	Type             string           `json:"type" yaml:"type"`
	Description      string           `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string           `json:"name" yaml:"name"`
	In               string           `json:"in" yaml:"in"`
	Scheme           string           `json:"scheme" yaml:"scheme"`
	BearerFormat     string           `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            OAuthFlowsObject `json:"flows" yaml:"flows"`
	OpenIDConnectURL string           `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
}

// OAuthFlowsObject represents
// https://spec.openapis.org/oas/v3.1.0#oauth-flows-object
//
// Field Name			Type				Description
// implicit				OAuth Flow Object	Configuration for the OAuth Implicit flow
// password				OAuth Flow Object	Configuration for the OAuth Resource Owner Password flow
// clientCredentials	OAuth Flow Object	Configuration for the OAuth Client Credentials flow. Previously called application in OpenAPI 2.0.
// authorizationCode	OAuth Flow Object	Configuration for the OAuth Authorization Code flow. Previously called accessCode in OpenAPI 2.0.
type OAuthFlowsObject struct {
	Implicit          OAuthFlowObject `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          OAuthFlowObject `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials OAuthFlowObject `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode OAuthFlowObject `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

// OAuthFlowObject represents
// https://spec.openapis.org/oas/v3.1.0#oauth-flow-object
//
// authorizationUrl	string				oauth2 ("implicit", "authorizationCode")	REQUIRED. The authorization URL to be used for this flow. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
// tokenUrl			string				oauth2 ("password", "clientCredentials", "authorizationCode")	REQUIRED. The token URL to be used for this flow. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
// refreshUrl		string				oauth2	The URL to be used for obtaining refresh tokens. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
// scopes			Map[string, string]	oauth2	REQUIRED. The available scopes for the OAuth2 security scheme. A map between the scope name and a short description for it. The map MAY be empty.
type OAuthFlowObject struct {
	AuthorizationURL string            `json:"authorizationUrl" yaml:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl" yaml:"tokenUrl"`
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}

// ContactObject represents
// https://spec.openapis.org/oas/v3.1.0#contact-object
//
// Field Name	Type	Description
// name			string	The identifying name of the contact person/organization.
// url			string	The URL pointing to the contact information. This MUST be in the form of a URL.
// email		string	The email address of the contact person/organization. This MUST be in the form of an email address.
type ContactObject struct {
	Name  *string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   *string `json:"url,omitempty" yaml:"url,omitempty"`
	Email *string `json:"email,omitempty" yaml:"email,omitempty"`
}

// LicenseObject represents
// https://spec.openapis.org/oas/v3.1.0#license-object
//
// Field Name	Type	Description
// name			string	REQUIRED. The license name used for the API.
// identifier	string	An SPDX license expression for the API. The identifier field is mutually exclusive of the url field.
// url			string	A URL to the license used for the API. This MUST be in the form of a URL. The url field is mutually exclusive of the identifier field.
type LicenseObject struct {
	Name       string  `json:"name" yaml:"name"`
	Identifier *string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	URL        *string `json:"url,omitempty" yaml:"url,omitempty"`
}

// SecurityRequirementObject represents
// https://spec.openapis.org/oas/v3.1.0#security-requirement-object
// Field Pattern	Type	Description
// {name}	[string]	Each name MUST correspond to a security scheme which is declared in the Security Schemes under the Components Object. If the security scheme is of type "oauth2" or "openIdConnect", then the value is a list of scope names required for the execution, and the list MAY be empty if authorization does not require a specified scope. For other security scheme types, the array MAY contain a list of role names which are required for the execution, but are not otherwise defined or exchanged in-band.
//
// api_key: []
type SecurityRequirementObject map[string][]string

type TagObject struct {
	Name         string            `json:"name" yaml:"name"`
	Description  string            `json:"description" yaml:"description"`
	ExternalDocs ExternalDocObject `json:"externalDocs" yaml:"externalDocs"`
}

// ReferenceObject represents
// https://spec.openapis.org/oas/v3.1.0#reference-object
//
// Field Name	Type	Description
// $ref			string	REQUIRED. The reference identifier. This MUST be in the form of a URI.
// summary		string	A short summary which by default SHOULD override that of the referenced component. If the referenced object-type does not allow a summary field, then this field has no effect.
// description	string	A description which by default SHOULD override that of the referenced component. CommonMark syntax MAY be used for rich text representation. If the referenced object-type does not allow a description field, then this field has no effect.
type ReferenceObject struct {
	Ref         string `json:"$ref" yaml:"$ref"`
	Summary     string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// ExternalDocObject comment...
// https://spec.openapis.org/oas/v3.1.0#external-documentation-object
//
// Field Name	Type	Description
// description	string	A description of the target documentation. CommonMark syntax MAY be used for rich text representation.
// url			string	REQUIRED. The URL for the target documentation. This MUST be in the form of a URL.
type ExternalDocObject struct {
	Description string `json:"description,omitempty" json:"description,omitempty"`
	URL         string `json:"url" yaml:"url"`
}
