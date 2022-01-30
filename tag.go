package openapi

// Tag -.
type Tag struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// Tags -.
type Tags []*Tag
