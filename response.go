package test2openapi

type Response struct {
	*ResponseObject
	*ReferenceObject
}

type Responses map[string]Response

type ResponseObject struct {
	Description string                `yaml:"description,omitempty"`
	Headers     map[string]Header     `yaml:"headers,omitempty"`
	Content     map[string]MediaType  `yaml:"content,omitempty"`
	Links       map[string]LinkObject `yaml:"links,omitempty"`
}

type Header struct {
	*HeaderObject
	*ReferenceObject
}

type Headers map[string]Header

type HeaderObject struct{}

type MediaType struct{}
type LinkObject struct{}
