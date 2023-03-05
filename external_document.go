package test2openapi

// ExternalDocObject comment...
// https://spec.openapis.org/oas/v3.1.0#external-documentation-object
//
// Field Name	Type	Description
// description	string	A description of the target documentation. CommonMark syntax MAY be used for rich text representation.
// url			string	REQUIRED. The URL for the target documentation. This MUST be in the form of a URL.
type ExternalDocObject struct {
	Description string `json:"description,omitempty" json:"description,omitempty"`
	URL         string `json:"url" yaml:"url"`
}
