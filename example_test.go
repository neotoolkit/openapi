package openapi_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/openapi"
)

func TestExamples_GetKeys(t *testing.T) {
	e := openapi.Examples{
		"first_example":  openapi.Example{},
		"second_example": openapi.Example{},
	}

	res := e.GetKeys()

	require.Equal(t, len(e), len(res))
}

func TestExampleToResponse(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		want interface{}
	}{
		{
			name: "",
			data: nil,
			want: nil,
		},
		{
			name: "",
			data: map[string]interface{}{},
			want: map[string]interface{}{},
		},
		{
			name: "",
			data: []interface{}{},
			want: []map[string]interface{}{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := openapi.ExampleToResponse(tc.data)

			require.Equal(t, tc.want, got)
		})
	}
}
