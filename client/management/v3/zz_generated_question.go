package client

const (
	QuestionType                   = "question"
	QuestionFieldDefault           = "default"
	QuestionFieldDescription       = "description"
	QuestionFieldGroup             = "group"
	QuestionFieldInvalidChars      = "invalidChars"
	QuestionFieldLabel             = "label"
	QuestionFieldMax               = "max"
	QuestionFieldMaxLength         = "maxLength"
	QuestionFieldMin               = "min"
	QuestionFieldMinLength         = "minLength"
	QuestionFieldOptions           = "options"
	QuestionFieldRequired          = "required"
	QuestionFieldShowIf            = "showIf"
	QuestionFieldShowSubquestionIf = "showSubquestionIf"
	QuestionFieldSubquestions      = "subquestions"
	QuestionFieldType              = "type"
	QuestionFieldValidChars        = "validChars"
	QuestionFieldVariable          = "variable"
)

type Question struct {
	Default           string        `json:"default,omitempty" yaml:"default,omitempty"`
	Description       string        `json:"description,omitempty" yaml:"description,omitempty"`
	Group             string        `json:"group,omitempty" yaml:"group,omitempty"`
	InvalidChars      string        `json:"invalidChars,omitempty" yaml:"invalid_chars,omitempty"`
	Label             string        `json:"label,omitempty" yaml:"label,omitempty"`
	Max               int64         `json:"max,omitempty" yaml:"max,omitempty"`
	MaxLength         int64         `json:"maxLength,omitempty" yaml:"max_length,omitempty"`
	Min               int64         `json:"min,omitempty" yaml:"min,omitempty"`
	MinLength         int64         `json:"minLength,omitempty" yaml:"min_length,omitempty"`
	Options           []string      `json:"options,omitempty" yaml:"options,omitempty"`
	Required          bool          `json:"required,omitempty" yaml:"required,omitempty"`
	ShowIf            string        `json:"showIf,omitempty" yaml:"show_if,omitempty"`
	ShowSubquestionIf string        `json:"showSubquestionIf,omitempty" yaml:"show_subquestion_if,omitempty"`
	Subquestions      []SubQuestion `json:"subquestions,omitempty" yaml:"subquestions,omitempty"`
	Type              string        `json:"type,omitempty" yaml:"type,omitempty"`
	ValidChars        string        `json:"validChars,omitempty" yaml:"valid_chars,omitempty"`
	Variable          string        `json:"variable,omitempty" yaml:"variable,omitempty"`
}
