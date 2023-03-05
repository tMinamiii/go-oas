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
	Schema   Schema              `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  string              `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]string   `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

type Encoding struct {
	ContentType   string            `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]Header `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string            `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool              `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool              `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}
