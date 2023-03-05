package test2openapi

type OpenAPI[Resp Response, Param Parameter] struct {
	OpenAPIVersion    string                `json:"openapi" yaml:"openapi"`
	Info              Info                  `json:"info" yaml:"info"`
	JSONSchemaDialect string                `json:"jsonSchemaDialect,omitempty" yaml:"jsonSchemaDialect,omitempty"`
	Servers           Servers               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths             Paths                 `json:"paths,omitempty" yaml:"paths,omitempty"`
	Webhooks          map[string]PathItem   `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
	Components        Components            `json:"components,omitempty" yaml:"components,omitempty"`
	Security          SecurityRequirement   `json:"security,omitempty" yaml:"security,omitempty"`
	Tags              Tags                  `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs      ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

type SecurityRequirement map[string][]string

type Tags []Tag

type Tag struct {
	Name         string                `json:"name" yaml:"name"`
	Description  string                `json:"description" yaml:"description"`
	ExternalDocs ExternalDocumentation `json:"externalDocs" yaml:"externalDocs"`
}
