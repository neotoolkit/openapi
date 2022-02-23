package openapi

// RequestBody -.
type RequestBody struct {
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Content     Content `json:"content,omitempty" yaml:"content,omitempty"`
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// IsRef -.
func (body RequestBody) IsRef() bool {
	return body.Ref != ""
}

// RequestBodies -.
type RequestBodies map[string]*RequestBody
