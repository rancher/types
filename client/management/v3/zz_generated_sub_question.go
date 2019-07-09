package client

const (
	SubQuestionType              = "subQuestion"
	SubQuestionFieldDefault      = "default"
	SubQuestionFieldDescription  = "description"
	SubQuestionFieldGroup        = "group"
	SubQuestionFieldInvalidChars = "invalidChars"
	SubQuestionFieldLabel        = "label"
	SubQuestionFieldMax          = "max"
	SubQuestionFieldMaxLength    = "maxLength"
	SubQuestionFieldMin          = "min"
	SubQuestionFieldMinLength    = "minLength"
	SubQuestionFieldOptions      = "options"
	SubQuestionFieldRequired     = "required"
	SubQuestionFieldShowIf       = "showIf"
	SubQuestionFieldType         = "type"
	SubQuestionFieldValidChars   = "validChars"
	SubQuestionFieldVariable     = "variable"
)

type SubQuestion struct {
	Default      string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description  string   `json:"description,omitempty" yaml:"description,omitempty"`
	Group        string   `json:"group,omitempty" yaml:"group,omitempty"`
	InvalidChars string   `json:"invalidChars,omitempty" yaml:"invalid_chars,omitempty"`
	Label        string   `json:"label,omitempty" yaml:"label,omitempty"`
	Max          int64    `json:"max,omitempty" yaml:"max,omitempty"`
	MaxLength    int64    `json:"maxLength,omitempty" yaml:"max_length,omitempty"`
	Min          int64    `json:"min,omitempty" yaml:"min,omitempty"`
	MinLength    int64    `json:"minLength,omitempty" yaml:"min_length,omitempty"`
	Options      []string `json:"options,omitempty" yaml:"options,omitempty"`
	Required     bool     `json:"required,omitempty" yaml:"required,omitempty"`
	ShowIf       string   `json:"showIf,omitempty" yaml:"show_if,omitempty"`
	Type         string   `json:"type,omitempty" yaml:"type,omitempty"`
	ValidChars   string   `json:"validChars,omitempty" yaml:"valid_chars,omitempty"`
	Variable     string   `json:"variable,omitempty" yaml:"variable,omitempty"`
}
