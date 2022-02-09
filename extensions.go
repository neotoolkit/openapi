package openapi

import "github.com/goccy/go-yaml"

// Extensions -.
type Extensions struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`
}

// GetExtensions -.
func (e Extensions) GetExtensions() (map[string]interface{}, error) {
	bytes, err := yaml.Marshal(e.Extensions)
	if err != nil {
		return nil, err
	}

	var extensions map[string]interface{}

	if err := yaml.Unmarshal(bytes, &extensions); err != nil {
		return nil, err
	}

	return extensions, nil
}
