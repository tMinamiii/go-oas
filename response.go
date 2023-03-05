package test2openapi

type Response struct {
	*ResponseObject
	*ReferenceObject
}

type Responses map[string]Response

type ResponseObject struct {
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]Header    `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]Link      `json:"links,omitempty" yaml:"links,omitempty"`
}

type MediaType struct {
	Schema   Schema              `json:"schema" yaml:"schema"`
	Example  string              `json:"example" yaml:"example"`
	Examples map[string]string   `json:"examples" yaml:"examples"`
	Encoding map[string]Encoding `json:"encoding" yaml:"encoding"`
}

type Encoding struct {
	ContentType   string            `json:"contentType" yaml:"contentType"`
	Headers       map[string]Header `json:"headers" yaml:"headers"`
	Style         string            `json:"style" yaml:"style"`
	Explode       bool              `json:"explode" yaml:"explode"`
	AllowReserved bool              `json:"allowReserved" yaml:"allowReserved"`
}
