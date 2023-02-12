package test2openapi

type Servers []Server

type Server struct {
	URL         string                    `yaml:"url"`                   // REQUIRED
	Description string                    `yaml:"description,omitempty"` // NULLABLE
	Variables   map[string]ServerVariable `yaml:"variable,omitempty"`
}

type ServerVariable struct {
	Enum        []string `yaml:"enum"`                  // MUST NOT be empty
	Default     string   `yaml:"default"`               // Required
	Description string   `yaml:"description,omitempty"` // NULLABLE
}
