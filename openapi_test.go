package openapi_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neotoolkit/openapi"
)

func TestSchemaError(t *testing.T) {
	got := &openapi.SchemaError{
		Ref: "test",
	}

	require.Equal(t, got.Error(), "unknown schema test")
}

func TestLookupByReference(t *testing.T) {
	api := openapi.OpenAPI{}

	schema, err := api.LookupSchemaByReference("")

	var schemaErr *openapi.SchemaError

	require.Equal(t, openapi.Schema{}, schema)
	require.True(t, errors.As(err, &schemaErr))
}

func TestRequestBodyError(t *testing.T) {
	got := &openapi.RequestBodyError{
		Ref: "test",
	}

	require.Equal(t, got.Error(), "unknown request body test")
}

func TestLookupRequestBodyByReference(t *testing.T) {
	api := openapi.OpenAPI{}

	schema, err := api.LookupRequestBodyByReference("")

	var requestBodyErr *openapi.RequestBodyError

	require.Equal(t, openapi.RequestBody{}, schema)
	require.True(t, errors.As(err, &requestBodyErr))
}
