package form

import (
	"fmt"
	"strings"
)

type inputType string

func (o *inputType) UnmarshalJSON(data []byte) error {
	strVar := strings.Trim(string(data), `"`)

	switch strVar {
	case "button", "checkbox", "color",
		"date", "datetime-local", "email",
		"file", "hidden", "image", "month",
		"number", "password", "radio", "range",
		"reset", "search", "submit", "tel", "time",
		"url", "week", "select", "textarea":
	default:
		return fmt.Errorf("unknown input type %q", strVar)
	}

	return nil
}

const (
	InputTypeButton        inputType = "button"
	InputTypeCheckbox      inputType = "checkbox"
	InputTypeColor         inputType = "color"
	InputTypeDate          inputType = "date"
	InputTypeDatetimeLocal inputType = "datetime-local"
	InputTypeEmail         inputType = "email"
	InputTypeFile          inputType = "file"
	InputTypeHidden        inputType = "hidden"
	InputTypeImage         inputType = "image"
	InputTypeMonth         inputType = "month"
	InputTypeNumber        inputType = "number"
	InputTypePassword      inputType = "password"
	InputTypeRadio         inputType = "radio"
	InputTypeRange         inputType = "range"
	InputTypeReset         inputType = "reset"
	InputTypeSearch        inputType = "search"
	InputTypeSubmit        inputType = "submit"
	InputTypeTel           inputType = "tel"
	InputTypeText          inputType = "text"
	InputTypeTime          inputType = "time"
	InputTypeUrl           inputType = "url"
	InputTypeWeek          inputType = "week"

	// non-standard
	InputTypeSelect   inputType = "select"
	InputTypeTextarea inputType = "textarea"
)
