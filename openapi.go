package openapi

import (
	"strings"
)

// SchemaError -.
type SchemaError struct {
	Ref string
}

// Error -.
func (e *SchemaError) Error() string {
	return "unknown schema " + e.Ref
}

// RequestBodyError -.
type RequestBodyError struct {
	Ref string
}

// Error -.
func (e *RequestBodyError) Error() string {
	return "unknown request body " + e.Ref
}

// OpenAPI Object
// See specification https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	Paths      Paths      `json:"paths" yaml:"paths"`
	Components Components `json:"components,omitempty" yaml:"components,omitempty"`
	OpenAPI    string     `json:"openapi" yaml:"openapi"`
	Servers    Servers    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Security   []Security `json:"security,omitempty" yaml:"security,omitempty"`
	Tags       Tags       `json:"tags,omitempty" yaml:"tags,omitempty"`
	Info       Info       `json:"info" yaml:"info"`
}

// LookupSchemaByReference -.
func (api OpenAPI) LookupSchemaByReference(ref string) (Schema, error) {
	schema, ok := api.Components.Schemas[schemaKey(ref)]
	if ok {
		return *schema, nil
	}

	return Schema{}, &SchemaError{Ref: ref}
}

func (api OpenAPI) LookupRequestBodyByReference(ref string) (RequestBody, error) {
	requestBody, ok := api.Components.RequestBodies[requestBodiesKey(ref)]
	if ok {
		return *requestBody, nil
	}

	return RequestBody{}, &RequestBodyError{Ref: ref}
}

func schemaKey(ref string) string {
	const prefix = "#/components/schemas/"
	return strings.TrimPrefix(ref, prefix)
}

func requestBodiesKey(ref string) string {
	const prefix = "#/components/requestBodies/"
	return strings.TrimPrefix(ref, prefix)
}
