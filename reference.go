package test2openapi

// ReferenceObject represents
// https://spec.openapis.org/oas/v3.1.0#reference-object
//
// Field Name	Type	Description
// $ref			string	REQUIRED. The reference identifier. This MUST be in the form of a URI.
// summary		string	A short summary which by default SHOULD override that of the referenced component. If the referenced object-type does not allow a summary field, then this field has no effect.
// description	string	A description which by default SHOULD override that of the referenced component. CommonMark syntax MAY be used for rich text representation. If the referenced object-type does not allow a description field, then this field has no effect.
type ReferenceObject struct {
	Ref         string `json:"$ref" yaml:"$ref"`
	Summary     string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
