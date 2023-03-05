package test2openapi

type Components []Component

type Component struct {
	Schemas         map[string]Schema         `json:"schemas" yaml:"schemas"`
	Responses       map[string]Response       `json:"responses" yaml:"responses"`
	Parameters      map[string]Parameter      `json:"parameters" yaml:"parameters"`
	Examples        map[string]Example        `json:"examples" yaml:"examples"`
	RequestBodies   map[string]RequestBody    `json:"requestBodies" yaml:"requestBodies"`
	Headers         map[string]Header         `json:"headers" yaml:"headers"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes" yaml:"securitySchemes"`
	Links           map[string]Link           `json:"links" yaml:"links"`
	Callbacks       map[string]Callback       `json:"callbacks" yaml:"callbacks"`
	PathItems       map[string]PathItem       `json:"pathItems" yaml:"pathItems"`
}

type Schema struct {
	ContentType   string            `json:"contentType" yaml:"contentType"`
	Header        map[string]Header `json:"header" yaml:"header"`
	Style         string            `json:"style" yaml:"style"`
	Explode       bool              `json:"explode" yaml:"explode"`
	AllowReserved bool              `json:"allowReserved" yaml:"allowReserved"`
}

type RequestBody struct {
	*RequestBodyObject
	*ReferenceObject
}

type RequestBodyObject struct{}

type Link struct {
	*LinkObject
	*ReferenceObject
}

type Callback struct {
	*CallbackObject
	*ReferenceObject
}

type CallbackObject map[string]PathItem

type Example struct {
	*ExampleObject
	*ReferenceObject
}

type ExampleObject struct {
	Description string
	Content     map[string]MediaType
	Required    bool
}

type SecurityScheme struct {
	*SecuritySchemeObject
	*ReferenceObject
}

type SecuritySchemeObject struct {
	Type             string
	Description      string
	Name             string
	In               string
	Scheme           string
	BearerFormat     string
	Flows            OAuthFlows
	OpenIDConnectURL string
}

type OAuthFlows struct {
	Implicit          OAuthFlow
	Password          OAuthFlow
	ClientCredentials OAuthFlow
	AuthorizationCode OAuthFlow
}

type OAuthFlow struct {
	AuthorizationURL string
	TokenURL         string
	RefreshURL       string
	Scopes           map[string]string
}
