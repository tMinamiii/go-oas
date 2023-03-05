package test2openapi

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
