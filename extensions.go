package openapi

type Extensions struct {
	Example string `json:"x-example,omitempty" yaml:"x-example,omitempty"`
	Faker   string `json:"x-faker,omitempty" yaml:"x-faker,omitempty"`
}
