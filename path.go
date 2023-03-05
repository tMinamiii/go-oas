package test2openapi

type PathItem struct {
	*PathItemObject
	*ReferenceObject
}

// Paths represents PathItem object slice
type Paths []PathItem

// PathItem represent
type PathItemObject struct {
	Summary     string      `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Get         *Operation  `json:"get,omitempty" yaml:"get,omitempty"`
	Put         *Operation  `json:"put,omitempty" yaml:"put,omitempty"`
	Post        *Operation  `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      *Operation  `json:"delete,omitempty" yaml:"delete,omitempty"`
	Patch       *Operation  `json:"patch,omitempty" yaml:"patch,omitempty"`
	Options     *Operation  `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *Operation  `json:"head,omitempty" yaml:"head,omitempty"`
	Trace       *Operation  `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     *Servers    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  *Parameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

type Operation struct {
	Tags         []string                  `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                    `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs ExternalDocumentation     `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  string                    `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   *Parameters               `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  RequestBodyObject         `json:"requrestBody,omitempty" yaml:"requrestBody,omitempty"`
	Responses    Responses                 `json:"responses,omitempty" yaml:"responses,omitempty"`
	Callbacks    map[string]CallbackObject `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   bool                      `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     SecuritySchemeObject      `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      *Servers                  `json:"servers,omitempty" yaml:"servers,omitempty"`
}
