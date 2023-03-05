package test2openapi

// HeaderOrRefObject represents HeaderObject or ReferenceObject
type HeaderOrRefObject struct {
	*HeaderObject
	*ReferenceObject
}

// HeaderObject represents
// https://spec.openapis.org/oas/v3.1.0#header-object
//
// Field Name	Type							Description
// name			string							REQUIRED. The name of the tag.
// description	string							A description for the tag. CommonMark syntax MAY be used for rich text representation.
// externalDocs	External Documentation Object	Additional external documentation for this tag.
type HeaderObject struct {
	Name         string            `json:"name" yaml:"name"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs ExternalDocObject `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}
