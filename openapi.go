package test2openapi

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
	OpenAPIVersion    string                `json:"openapi" yaml:"openapi"`
	Info              Info                  `json:"info" yaml:"info"`
	JSONSchemaDialect string                `json:"jsonSchemaDialect,omitempty" yaml:"jsonSchemaDialect,omitempty"`
	Servers           []ServerObject        `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths             PathsObject           `json:"paths,omitempty" yaml:"paths,omitempty"`
	Webhooks          []PathItemOrRefObject `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
	Components        ComponentsObject      `json:"components,omitempty" yaml:"components,omitempty"`
	Security          []SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
	Tags              []Tag                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs      ExternalDocObject     `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// SecurityRequirement represents
// https://spec.openapis.org/oas/v3.1.0#security-requirement-object
// Field Pattern	Type	Description
// {name}	[string]	Each name MUST correspond to a security scheme which is declared in the Security Schemes under the Components Object. If the security scheme is of type "oauth2" or "openIdConnect", then the value is a list of scope names required for the execution, and the list MAY be empty if authorization does not require a specified scope. For other security scheme types, the array MAY contain a list of role names which are required for the execution, but are not otherwise defined or exchanged in-band.
//
// api_key: []
type SecurityRequirement map[string][]string

type Tag struct {
	Name         string            `json:"name" yaml:"name"`
	Description  string            `json:"description" yaml:"description"`
	ExternalDocs ExternalDocObject `json:"externalDocs" yaml:"externalDocs"`
}
