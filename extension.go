package openapi

type ExtensionProps struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`
}
