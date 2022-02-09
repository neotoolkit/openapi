package openapi_test

import (
	"github.com/neotoolkit/openapi"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetExtensions(t *testing.T) {
	tests := []struct {
		name       string
		extensions map[string]interface{}
		want       interface{}
		err        bool
	}{
		{
			name:       "",
			extensions: nil,
			want:       map[string]interface{}{},
			err:        false,
		},
		{
			name: "",
			extensions: map[string]interface{}{
				"chan": make(chan byte),
			},
			want: map[string]interface{}(nil),
			err:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := openapi.Extensions{}

			s.Extensions = tc.extensions

			e, err := s.GetExtensions()

			require.Equal(t, tc.want, e)
			require.True(t, err != nil == tc.err)
		})
	}
}
