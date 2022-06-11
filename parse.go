package openapi

import (
	"github.com/ghodss/yaml"
)

func Parse(data []byte) (OpenAPI, error) {
	var openapi OpenAPI

	err := yaml.Unmarshal(data, &openapi)
	if err != nil {
		return OpenAPI{}, err
	}

	return openapi, nil
}
