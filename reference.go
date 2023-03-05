package test2openapi

type ReferenceObject struct {
	Ref         string `json:"$ref" yaml:"$ref"`
	Summary     string `json:"summary" yaml:"summary"`
	Description string `json:"description" yaml:"description"`
}
