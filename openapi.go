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

// OpenAPI Object
// See specification https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	OpenAPI    string     `json:"openapi" yaml:"openapi"`
	Info       Info       `json:"info" yaml:"info"`
	Servers    Servers    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths      Paths      `json:"paths" yaml:"paths"`
	Components Components `json:"components,omitempty" yaml:"components,omitempty"`
	Security   []Security `json:"security,omitempty" yaml:"security,omitempty"`
	Tags       Tags       `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// LookupByReference -.
func (api OpenAPI) LookupByReference(ref string) (Schema, error) {
	schema, ok := api.Components.Schemas[schemaKey(ref)]
	if ok {
		return *schema, nil
	}

	schema, ok = api.Components.Schemas[requestBodiesKey(ref)]
	if ok {
		return *schema, nil
	}

	return Schema{}, &SchemaError{Ref: ref}
}

func schemaKey(ref string) string {
	const prefix = "#/components/schemas/"
	return strings.TrimPrefix(ref, prefix)
}

func requestBodiesKey(ref string) string {
	const prefix = "#/components/requestBodies/"
	return strings.TrimPrefix(ref, prefix)
}
