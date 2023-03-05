package test2openapi

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
