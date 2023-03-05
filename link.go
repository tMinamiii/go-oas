package test2openapi

type Link struct {
	*LinkObject
	*ReferenceObject
}

type LinkObject struct {
	OperationRef string            `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationID  string            `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  string            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty"`
	Server       Server            `json:"server,omitempty" yaml:"server,omitempty"`
}
