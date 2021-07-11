## Form

This package provides an `Option` type which can be used to represent a HTML form making it easy to use forms in templates.

```go
type Option struct {
	ID       string     `json:"id,omitempty"`
	Label    string     `json:"label,omitempty"`
	HelpText string     `json:"help_text,omitempty"`
	Kind     optionKind `json:"kind,omitempty"`

	Default   string         `json:"default,omitempty"`
	Required  bool           `json:"required,omitempty"`
	Min       string         `json:"min,omitempty"`
	Max       string         `json:"max,omitempty"`
	MinLength int            `json:"min_length,omitempty"`
	MaxLength int            `json:"max_length,omitempty"`
	Pattern   string         `json:"pattern,omitempty"`
	Options   []SelectOption `json:"options,omitempty"`
	List      []string       `json:"list,omitempty"`

	// Used in templates
	FormName  string `json:"-"`
	FormValue string `json:"-"`
}
```

