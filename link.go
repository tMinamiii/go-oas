package test2openapi

type Link struct {
	*LinkObject
	*ReferenceObject
}

type LinkObject struct {
	OperationRef string            `json:"operationRef" yaml:"operationRef"`
	OperationID  string            `json:"operationId" yaml:"operationId"`
	Parameters   map[string]string `json:"parameters" yaml:"parameters"`
}
