package test2openapi

type Headers []Header

type Header struct {
	*HeaderObject
	*ReferenceObject
}

type HeaderObject struct {
	Name         string                `json:"name" yaml:"name"`
	Description  string                `json:"description" yaml:"description"`
	ExternalDocs ExternalDocumentation `json:"externalDocs" yaml:"externalDocs"`
}
