package test2openapi

type Headers []Header

type Header struct {
	*HeaderObject
	*ReferenceObject
}

type HeaderObject struct {
	Name         string
	Description  string
	ExternalDocs ExternalDocumentation
}
