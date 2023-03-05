package test2openapi

// Headers comment...
type Headers []Header

// Header comment...
type Header struct {
	*HeaderObject
	*ReferenceObject
}

// HeaderObject comment...
type HeaderObject struct {
	Name         string                `json:"name" yaml:"name"`
	Description  string                `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}
