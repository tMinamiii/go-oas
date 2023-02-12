package test2openapi

type Info struct {
	Title          string  `yaml:"title,omitempty"`
	Summary        string  `yaml:"summary,omitempty"`
	Description    string  `yaml:"description,omitempty"`
	TermsOfService string  `yaml:"termsOfService,omitempty"`
	Contact        Contact `yaml:"contact,omitempty"`
	License        License `yaml:"license,omitempty"`
	Version        string  `yaml:"version,omitempty"`
}

type Contact struct {
	Name  string `yaml:"name,omitempty"`
	URL   string `yaml:"url,omitempty"`
	Email string `yaml:"email,omitempty"`
}

type License struct {
	name       string `yaml:"name,omitempty"`
	Identifier string `yaml:"identifier,omitempty"`
	URL        string `yaml:"url,omitempty"`
}
