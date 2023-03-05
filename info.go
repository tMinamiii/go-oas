package test2openapi

type Info struct {
	Title          string  `json:"title,omitempty" yaml:"title,omitempty"`
	Summary        string  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description    string  `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string  `json:"version,omitempty" yaml:"version,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

type License struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	URL        string `json:"url,omitempty" yaml:"url,omitempty"`
}
