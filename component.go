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

type RequestBodyObject struct {
	Description string               `json:"description" yaml:"description"`
	Content     map[string]MediaType `json:"content" yaml:"content"`
	Required    bool                 `json:"required" yaml:"required"`
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
	Description string               `json:"description" yaml:"description"`
	Content     map[string]MediaType `json:"content" yaml:"content"`
	Required    bool                 `json:"required" yaml:"required"`
}

type SecurityScheme struct {
	*SecuritySchemeObject
	*ReferenceObject
}

type SecuritySchemeObject struct {
	Type             string     `json:"type" yaml:"type"`
	Description      string     `json:"description" yaml:"description"`
	Name             string     `json:"name" yaml:"name"`
	In               string     `json:"in" yaml:"in"`
	Scheme           string     `json:"scheme" yaml:"scheme"`
	BearerFormat     string     `json:"bearerFormat" yaml:"bearerFormat"`
	Flows            OAuthFlows `json:"flows" yaml:"flows"`
	OpenIDConnectURL string     `json:"openIdConnectUrl" yaml:"openIdConnectUrl"`
}

type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit" yaml:"implicit"`
	Password          *OAuthFlow `json:"password" yaml:"password"`
	ClientCredentials *OAuthFlow `json:"clientCredentials" yaml:"clientCredentials"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode" yaml:"authorizationCode"`
}

type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl" yaml:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl" yaml:"tokenUrl"`
	RefreshURL       string            `json:"refreshUrl" yaml:"refreshUrl"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}
