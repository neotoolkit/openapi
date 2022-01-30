package openapi_test

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/go-dummy/openapi"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		path string
		want openapi.OpenAPI
		err  error
	}{
		{
			name: "wrong yml",
			path: "./testdata/wrong-openapi.yml",
			want: openapi.OpenAPI{},
			err:  errors.New("[1:1] string was used where mapping is expected\n>  1 | openapi\n       ^\n"),
		},
		{
			name: "",
			path: "./testdata/openapi.yml",
			want: openapi.OpenAPI{
				OpenAPI: "3.0.3",
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data, err := ioutil.ReadFile(tc.path)
			if err != nil {
				t.Fatal(err)
			}
			got, err := openapi.Parse(data)
			if err != nil {
				require.EqualError(t, err, tc.err.Error())
			}
			require.Equal(t, tc.want, got)
		})
	}

}
