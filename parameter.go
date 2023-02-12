package test2openapi

type Parameter struct {
	*ParameterObject
	*ReferenceObject
}

type Parameters []Parameter

type ParameterObject struct {
	Name            string `yaml:"name,omitempty"`
	In              string `yaml:"in,omitempty"`
	Description     string `yaml:"description,omitempty"`
	Required        bool   `yaml:"required,omitempty"`
	Deprecated      bool   `yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `yaml:"allowEmptyValue,omitempty"`
}
