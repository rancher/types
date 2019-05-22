package client

const (
	ExampleSpecType               = "exampleSpec"
	ExampleSpecFieldExampleString = "rancherCompose"
)

type ExampleSpec struct {
	ExampleString string `json:"rancherCompose,omitempty" yaml:"rancherCompose,omitempty"`
}
