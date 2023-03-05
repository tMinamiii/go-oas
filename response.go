package test2openapi

type Response struct {
	*ResponseObject
	*ReferenceObject
}

type Responses map[string]Response

type ResponseObject struct {
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]Header     `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]MediaType  `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]LinkObject `json:"links,omitempty" yaml:"links,omitempty"`
}

type MediaType struct {
	Schema   Schema
	Example  string
	Examples map[string]string
	encoding map[string]Encoding
}

type Encoding struct {
	ContentType   string
	Headers       map[string]Header
	Style         string
	explode       bool
	allowReserved bool
}

type LinkObject struct {
	OperationRef string
	OperationID  string
	Parameters   map[string]string
}
