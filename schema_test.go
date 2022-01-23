package openapi_test

import (
	"fmt"
	"testing"

	"github.com/go-dummy/openapi"

	"github.com/stretchr/testify/require"
)

type schemaContextStub struct{}

func (s schemaContextStub) LookupByReference(ref string) (openapi.Schema, error) {
	userSchema := openapi.Schema{
		Properties: openapi.Schemas{
			"id": &openapi.Schema{
				Type:    "string",
				Format:  "uuid",
				Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
			},
			"firstName": &openapi.Schema{
				Type:    "string",
				Example: "Larry",
			},
			"lastName": &openapi.Schema{
				Type:    "string",
				Example: "Page",
			},
		},
		Type: "object",
	}

	uuidSchema := openapi.Schema{
		Type:    "string",
		Format:  "uuid",
		Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
	}

	switch ref {
	case "#/components/schemas/User":
		return userSchema, nil
	case "#/components/schemas/UUID":
		return uuidSchema, nil
	default:
		return openapi.Schema{}, fmt.Errorf("unknown schema: %q", ref)
	}
}

func TestSchema_ResponseByExample(t *testing.T) {
	type fields struct {
		Properties openapi.Schemas
		Type       string
		Format     string
		Default    interface{}
		Example    interface{}
		Faker      string
		Items      *openapi.Schema
		Reference  string
	}

	type args struct {
		schemaContext openapi.SchemaContext
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes interface{}
		wantErr bool
	}{
		{
			name: "Simple schema",
			fields: fields{
				Properties: openapi.Schemas{
					"id": &openapi.Schema{
						Type:    "string",
						Format:  "uuid",
						Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
					},
					"firstName": &openapi.Schema{
						Type:    "string",
						Example: "Larry",
					},
					"lastName": &openapi.Schema{
						Type:    "string",
						Example: "Page",
					},
				},
				Type: "object",
			},
			args: args{
				schemaContext: schemaContextStub{},
			},
			wantRes: map[string]interface{}{
				"id":        "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
				"firstName": "Larry",
				"lastName":  "Page",
			},
			wantErr: false,
		},
		{
			name: "Simple schema with reference",
			fields: fields{
				Reference: "#/components/schemas/User",
			},
			args: args{
				schemaContext: schemaContextStub{},
			},
			wantRes: map[string]interface{}{
				"id":        "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
				"firstName": "Larry",
				"lastName":  "Page",
			},
			wantErr: false,
		},
		{
			name: "Array schema with reference",
			fields: fields{
				Type: "array",
				Items: &openapi.Schema{
					Reference: "#/components/schemas/User",
				},
			},
			args: args{
				schemaContext: schemaContextStub{},
			},
			wantRes: []interface{}{
				map[string]interface{}{
					"id":        "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
					"firstName": "Larry",
					"lastName":  "Page",
				},
			},
			wantErr: false,
		},
		{
			name: "Schema property with reference",
			fields: fields{
				Properties: openapi.Schemas{
					"id": &openapi.Schema{
						Reference: "#/components/schemas/UUID",
					},
					"firstName": &openapi.Schema{
						Type:    "string",
						Example: "Larry",
					},
					"lastName": &openapi.Schema{
						Type:    "string",
						Example: "Page",
					},
				},
				Type: "object",
			},
			args: args{
				schemaContext: schemaContextStub{},
			},
			wantRes: map[string]interface{}{
				"id":        "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
				"firstName": "Larry",
				"lastName":  "Page",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := openapi.Schema{
				Properties: tc.fields.Properties,
				Type:       tc.fields.Type,
				Format:     tc.fields.Format,
				Default:    tc.fields.Default,
				Example:    tc.fields.Example,
				Faker:      tc.fields.Faker,
				Items:      tc.fields.Items,
				Reference:  tc.fields.Reference,
			}
			gotRes, err := s.ResponseByExample(tc.args.schemaContext)

			require.NoError(t, err)
			require.Equal(t, tc.wantRes, gotRes)
			require.Equal(t, tc.wantErr, err != nil)
		})
	}
}
