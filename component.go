package test2openapi

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
