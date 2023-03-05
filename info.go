package test2openapi

// Info comment ...
// https://spec.openapis.org/oas/v3.1.0#info-object
//
// Field Name		Type			Description
// title			string			REQUIRED. The title of the API.
// summary			string			A short summary of the API.
// description		string			A description of the API. CommonMark syntax MAY be used for rich text representation.
// termsOfService	string			A URL to the Terms of Service for the API. This MUST be in the form of a URL.
// contact			Contact Object	The contact information for the exposed API.
// license			License Object	The license information for the exposed API.
// version			string			REQUIRED. The version of the OpenAPI document (which is distinct from the OpenAPI Specification version or the API implementation version).
type Info struct {
	Title          string   `json:"title" yaml:"title"`
	Summary        *string  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description    *string  `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService *string  `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"`
}

// Contact represents
// https://spec.openapis.org/oas/v3.1.0#contact-object
//
// Field Name	Type	Description
// name			string	The identifying name of the contact person/organization.
// url			string	The URL pointing to the contact information. This MUST be in the form of a URL.
// email		string	The email address of the contact person/organization. This MUST be in the form of an email address.
type Contact struct {
	Name  *string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   *string `json:"url,omitempty" yaml:"url,omitempty"`
	Email *string `json:"email,omitempty" yaml:"email,omitempty"`
}

// License represents
// https://spec.openapis.org/oas/v3.1.0#license-object
//
// Field Name	Type	Description
// name			string	REQUIRED. The license name used for the API.
// identifier	string	An SPDX license expression for the API. The identifier field is mutually exclusive of the url field.
// url			string	A URL to the license used for the API. This MUST be in the form of a URL. The url field is mutually exclusive of the identifier field.
type License struct {
	Name       string  `json:"name" yaml:"name"`
	Identifier *string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	URL        *string `json:"url,omitempty" yaml:"url,omitempty"`
}
