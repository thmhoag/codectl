package template

type Parameter struct {

	// Name of the parameter
	Name string `yaml:"name"`

	// Prompt is the text to be displayed when prompting
	// the user for the parameter
	Prompt string `yaml:"prompt"`

	// Value of the parameter
	Value string `yaml:"value"`

	// Required indicates if the parameter is required
	Required bool `yaml:"required"`
}