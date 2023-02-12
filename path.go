package test2openapi

type PathItem struct {
	*PathItemObject
	*ReferenceObject
}

// Paths represents PathItem object slice
type Paths []PathItem

// PathItem represent
type PathItemObject struct {
	Summary     string      `yaml:"summary,omitempty"`
	Description string      `yaml:"description,omitempty"`
	Get         *Operation  `yaml:"get,omitempty"`
	Put         *Operation  `yaml:"put,omitempty"`
	Post        *Operation  `yaml:"post,omitempty"`
	Delete      *Operation  `yaml:"delete,omitempty"`
	Patch       *Operation  `yaml:"patch,omitempty"`
	Options     *Operation  `yaml:"options,omitempty"`
	Head        *Operation  `yaml:"head,omitempty"`
	Trace       *Operation  `yaml:"trace,omitempty"`
	Servers     *Servers    `yaml:"servers,omitempty"`
	Parameters  *Parameters `yaml:"parameters,omitempty"`
}

type Operation struct {
	Tags         []string                  `yaml:"tags,omitempty"`
	Summary      string                    `yaml:"summary,omitempty"`
	Description  string                    `yaml:"description,omitempty"`
	ExternalDocs ExternalDocumentation     `yaml:"externalDocs,omitempty"`
	OperationID  string                    `yaml:"operationId,omitempty"`
	Parameters   *Parameters               `yaml:"parameters,omitempty"`
	RequestBody  RequestBodyObject         `yaml:"requrestBody,omitempty"`
	Responses    Responses                 `yaml:"responses,omitempty"`
	Callbacks    map[string]CallbackObject `yaml:"callbacks,omitempty"`
	Deprecated   bool                      `yaml:"deprecated,omitempty"`
	Security     SecuritySchemeObject      `yaml:"security,omitempty"`
	Servers      *Servers                  `yaml:"servers,omitempty"`
}
