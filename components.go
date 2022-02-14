package openapi

// Components -.
type Components struct {
	Schemas       Schemas `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	RequestBodies Schemas `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
}
