package test2openapi

// ServerObject represents
// https://spec.openapis.org/oas/v3.1.0#server-object
//
// Field Name	Type								Description
// url			string								REQUIRED. A URL to the target host. This URL supports ServerObject Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenAPI document is being served. Variable substitutions will be made when a variable is named in {brackets}.
// description	string								An optional string describing the host designated by the URL. CommonMark syntax MAY be used for rich text representation.
// variables	Map[string, ServerObject Variable Object]	A map between a variable name and its value. The value is used for substitution in the server’s URL template.
type ServerObject struct {
	URL         string                    `json:"url" yaml:"url"`                                     // REQUIRED
	Description string                    `json:"description,omitempty" yaml:"description,omitempty"` // NULLABLE
	Variables   map[string]ServerVariable `json:"variable,omitempty" yaml:"variable,omitempty"`
}

// ServerVariable represents
// https://spec.openapis.org/oas/v3.1.0#server-variable-object
//
// Field Name	Type		Description
// enum			[string]	An enumeration of string values to be used if the substitution options are from a limited set. The array MUST NOT be empty.
// default		string		REQUIRED. The default value to use for substitution, which SHALL be sent if an alternate value is not supplied. Note this behavior is different than the Schema Object’s treatment of default values, because in those cases parameter values are optional. If the enum is defined, the value MUST exist in the enum’s values.
// description	string		An optional description for the server variable. CommonMark syntax MAY be used for rich text representation.
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`               // MUST NOT be empty
	Default     string   `json:"default" yaml:"default"`                             // Required
	Description string   `json:"description,omitempty" yaml:"description,omitempty"` // NULLABLE
}
