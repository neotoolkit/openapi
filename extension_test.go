package openapi_test

import (
	"github.com/neotoolkit/openapi"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetExtensions(t *testing.T) {
	s := openapi.Extensions{}

	s.Extensions = map[string]interface{}{
		"test": "test",
	}

	e, err := s.GetExtensions()
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "test", e["test"])
}
