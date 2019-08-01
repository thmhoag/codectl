package template

type Properties struct {

	// Name of the template
	Name string `yaml:"name"`

	// Version of the template
	Version string `yaml:"version"`

	// Description of the template
	Description string `yaml:"description"`

	// Parameters is the list of parameters expected by the template
	Parameters []Parameter `yaml:"parameters"`

	// Dependencies is a list of the paths of the dependent templates that will be
	// called with this one
	Dependencies []string `yaml:"dependencies"`
}
