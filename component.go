package test2openapi

type Components []Component

// Component comment...
type Component struct {
	Schemas         map[string]Schema         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]Response       `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]Example        `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]RequestBody    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]Header         `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           map[string]Link           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]Callback       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	PathItems       map[string]PathItem       `json:"pathItems,omitempty" yaml:"pathItems,omitempty"`
}

// Schema comment...
type Schema struct {
	Discriminator Discriminator         `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	XML           XML                   `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExternalDocs  ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Example       bool                  `json:"example,omitempty" yaml:"example,omitempty"`
}

// Discriminator comment...
type Discriminator struct {
	PropertyName string            `json:"propertyName" yaml:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

// XML comment...
type XML struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}

type RequestBody struct {
	*RequestBodyObject
	*ReferenceObject
}

// RequestBodyObject comment...
type RequestBodyObject struct {
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]MediaType `json:"content" yaml:"content"`
	Required    bool                 `json:"required,omitempty" yaml:"required,omitempty"`
}

// Callback comment...
type Callback struct {
	*CallbackObject
	*ReferenceObject
}

// CallbackObject comment...
type CallbackObject map[string]PathItem

// Example comment...
type Example struct {
	*ExampleObject
	*ReferenceObject
}

// ExampleObject comment...
type ExampleObject struct {
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Required    bool                 `json:"required,omitempty" yaml:"required,omitempty"`
}

// SecurityScheme comment...
type SecurityScheme struct {
	*SecuritySchemeObject
	*ReferenceObject
}

// SecuritySchemeObject comment...
type SecuritySchemeObject struct {
	Type             string     `json:"type" yaml:"type"`
	Description      string     `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string     `json:"name" yaml:"name"`
	In               string     `json:"in" yaml:"in"`
	Scheme           string     `json:"scheme" yaml:"scheme"`
	BearerFormat     string     `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            OAuthFlows `json:"flows" yaml:"flows"`
	OpenIDConnectURL string     `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
}

// OAuthFlows comment...
type OAuthFlows struct {
	Implicit          OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

// OAuthFlow comment...
type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl" yaml:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl" yaml:"tokenUrl"`
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}
