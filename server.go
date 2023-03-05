package test2openapi

type Servers []Server

type Server struct {
	URL         string                    `json:"url" yaml:"url"`                                     // REQUIRED
	Description string                    `json:"description,omitempty" yaml:"description,omitempty"` // NULLABLE
	Variables   map[string]ServerVariable `json:"variable,omitempty" yaml:"variable,omitempty"`
}

type ServerVariable struct {
	Enum        []string `json:"enum" yaml:"enum"`                                   // MUST NOT be empty
	Default     string   `json:"default" yaml:"default"`                             // Required
	Description string   `json:"description,omitempty" yaml:"description,omitempty"` // NULLABLE
}
