package test2openapi

type Parameters []Parameter

type Parameter struct {
	*ParameterObject
	*ReferenceObject
}

type ParameterObject struct {
	Name            string `json:"name,omitempty" yaml:"name,omitempty"`
	In              string `json:"in,omitempty" yaml:"in,omitempty"`
	Description     string `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool   `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      bool   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
}
