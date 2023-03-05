package test2openapi

type OpenAPI[Resp Response, Param Parameter] struct {
	OpenAPIVersion    string                `json:"openapi" yaml:"openapi"`
	Info              Info                  `json:"info,omitempty" yaml:"info,omitempty"`
	JSONSchemaDialect string                `json:"jsonSchemaDialect" yaml:"jsonSchemaDialect"`
	Servers           Servers               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths             Paths                 `json:"paths,omitempty" yaml:"paths,omitempty"`
	Webhooks          map[string]PathItem   `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
	Components        Components            `json:"components,omitempty" yaml:"components,omitempty"`
	Security          SecurityRequirement   `json:"security,omitempty" yaml:"security,omitempty"`
	Tags              Tags                  `json:"tags" yaml:"tags"`
	ExternalDocs      ExternalDocumentation `json:"externalDocs" yaml:"externalDocs"`
}

type SecurityRequirement map[string][]string

type Tags []Tag

type Tag struct {
	Name         string
	Description  string
	ExternalDocs ExternalDocumentation
}
