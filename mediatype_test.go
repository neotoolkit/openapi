package openapi_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neotoolkit/openapi"
)

func TestMediaType_ResponseByExample(t *testing.T) {
	m := openapi.MediaType{
		Example: []interface{}{},
	}

	require.IsType(t, []map[string]interface{}{}, m.ResponseByExample())
}

func TestMediaType_ResponseByExamplesKey(t *testing.T) {
	const key = "key"

	m := openapi.MediaType{
		Examples: openapi.Examples{
			key: openapi.Example{
				Value: map[string]interface{}{
					"key": "value",
				},
			},
		},
	}

	require.IsType(t, map[string]interface{}{"key": "value"}, m.ResponseByExamplesKey(key))
}
