package template

// Overrides defines replacement text for names of items in the templates
type Overrides struct {
	//Paths contains the file/folder names and the override values
	Paths map[string]string `yaml:"paths"`
}