package form

import (
	"errors"
	"fmt"
	"strings"
)

// Option represents a form field
//
// Deprecated: Use [Field] instead
type Option struct {
	ID       string     `json:"id,omitempty"`
	Label    string     `json:"label,omitempty"`
	HelpText string     `json:"help_text,omitempty"`
	Kind     optionKind `json:"kind,omitempty"`

	Default      string         `json:"default,omitempty"`
	Required     bool           `json:"required,omitempty"`
	Min          string         `json:"min,omitempty"`
	Max          string         `json:"max,omitempty"`
	Step         string         `json:"step,omitempty"`
	MinLength    int            `json:"min_length,omitempty"`
	MaxLength    int            `json:"max_length,omitempty"`
	Pattern      string         `json:"pattern,omitempty"`
	Options      []SelectOption `json:"options,omitempty"`
	List         []string       `json:"list,omitempty"`
	Placeholder  string         `json:"placeholder,omitempty"`
	Title        string         `json:"title,omitempty"`
	Autocomplete string         `json:"autocomplete,omitempty"`

	// Used in templates
	FormName  string `json:"-"`
	FormValue string `json:"-"`
}

func (o Option) WithFormName(name string) Option {
	o.FormName = name
	return o
}

func (o Option) WithFormValue(val string) Option {
	o.FormValue = val
	return o
}

// To be able to satisfy an interface in case
// other custom options are needed
func (o Option) GetOption() Option {
	return o
}

type optionKind string

func (o *optionKind) UnmarshalJSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("unknown option kind")
		}
	}()

	strVar := strings.Trim(string(data), `"`)

	*o = MustKind(strVar)
	return
}

// MustKind panics if the given string is not a valid kind
func MustKind(name string) optionKind {
	theKind := optionKind(name)
	for _, k := range OptionKinds {
		if theKind == k {
			return optionKind(name)
		}
	}

	panic(fmt.Errorf("%q is not a valid option kind", name))
}

var OptionKinds = []optionKind{
	"button",
	"checkbox",
	"color",
	"date",
	"datetime-local",
	"email",
	"file",
	"hidden",
	"image",
	"month",
	"number",
	"password",
	"radio",
	"range",
	"reset",
	"search",
	"submit",
	"tel",
	"text",
	"time",
	"url",
	"week",

	// non-standard
	"select",
	"textarea",
}
