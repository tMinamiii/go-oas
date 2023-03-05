package test2openapi

type Parameters []Parameter

type Parameter struct {
	*ParameterObject
	*ReferenceObject
}

type ParameterObject struct {
	Name            string `json:"name" yaml:"name"`
	In              string `json:"in" yaml:"in"`
	Description     string `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool   `json:"required" yaml:"required"`
	Deprecated      bool   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
}
