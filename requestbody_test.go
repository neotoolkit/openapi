package openapi_test

import (
	"testing"

	"github.com/neotoolkit/openapi"
	"github.com/stretchr/testify/require"
)

func TestRequestBody_IsRef(t *testing.T) {
	type testCase struct {
		name       string
		refValue   string
		wantResult bool
	}

	tests := []testCase{
		{
			name:       "Present ref value",
			refValue:   "Some text",
			wantResult: true,
		},
		{
			name:       "Empty ref value",
			wantResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			requestBody := openapi.RequestBody{
				Ref: tc.refValue,
			}

			actualResult := requestBody.IsRef()
			require.Equal(t, tc.wantResult, actualResult)
		})
	}
}
