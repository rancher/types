package client

const (
	ValuesType                  = "values"
	ValuesFieldBoolValue        = "boolValue"
	ValuesFieldIntValue         = "intValue"
	ValuesFieldStringSliceValue = "stringSliceValue"
	ValuesFieldStringValue      = "stringValue"
)

type Values struct {
	BoolValue        bool     `json:"boolValue,omitempty" yaml:"bool_value,omitempty"`
	IntValue         int64    `json:"intValue,omitempty" yaml:"int_value,omitempty"`
	StringSliceValue []string `json:"stringSliceValue,omitempty" yaml:"string_slice_value,omitempty"`
	StringValue      string   `json:"stringValue,omitempty" yaml:"string_value,omitempty"`
}
