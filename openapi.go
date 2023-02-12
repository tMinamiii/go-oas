package test2openapi

type Example struct {
	*ExampleObject
	*ReferenceObject
}

type RequestBody struct {
	*RequestBodyObject
	*ReferenceObject
}

type SecurityScheme struct {
	*SecuritySchemeObject
	*ReferenceObject
}

type Link struct {
	*LinkObject
	*ReferenceObject
}

type Callback struct {
	*CallbackObject
	*ReferenceObject
}

type OpenAPI[Resp Response, Param Parameter] struct {
	OpenAPIVersion string              `yaml:"openapi"`
	Info           Info                `yaml:"info,omitempty"`
	Servers        Servers             `yaml:"servers,omitempty"`
	Paths          Paths               `yaml:"paths,omitempty"`
	Webhooks       map[string]PathItem `yaml:"webhooks,omitempty"`
	Components     Components          `yaml:"components,omitempty"`
	Security       SecurityRequirement `yaml:"security,omitempty"`
}

type ReferenceObject struct{}
type ExampleObject struct{}
type CallbackObject struct{}
type SecuritySchemeObject struct{}
type RequestBodyObject struct{}

type ExternalDocumentation struct{}

type Components []Component

type Component struct{}

type SecurityRequirement map[string][]string
