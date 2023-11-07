package form

// Field represents a form field
// T is the type of "output" strings
// The easy way to use this is to use string
// Another possibility is to use *i18n.Message for internationalization
type Field[T any] struct {
	ID       string    `json:"id,omitempty"`
	Label    T         `json:"label,omitempty"`
	HelpText T         `json:"help_text,omitempty"`
	Type     inputType `json:"kind,omitempty"`

	Default      string         `json:"default,omitempty"`
	Required     bool           `json:"required,omitempty"`
	Multiple     bool           `json:"multiple,omitempty"`
	Min          string         `json:"min,omitempty"`
	Max          string         `json:"max,omitempty"`
	Step         string         `json:"step,omitempty"`
	MinLength    int            `json:"min_length,omitempty"`
	MaxLength    int            `json:"max_length,omitempty"`
	Pattern      string         `json:"pattern,omitempty"`
	Options      []SelectOption `json:"options,omitempty"`
	List         string         `json:"list,omitempty"`
	Placeholder  T              `json:"placeholder,omitempty"`
	Title        T              `json:"title,omitempty"`
	Autocomplete string         `json:"autocomplete,omitempty"`

	// Used in templates
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// SelectOption represents the <option> tag
type SelectOption struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Selected bool   `json:"selected,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}
