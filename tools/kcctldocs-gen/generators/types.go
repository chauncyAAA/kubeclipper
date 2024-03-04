package generators

type KcctlSpec struct {
	TopLevelCommandGroups []TopLevelCommands `yaml:",omitempty"`
}

type TopLevelCommands struct {
	Group    string            `yaml:",omitempty"`
	Commands []TopLevelCommand `yaml:",omitempty"`
}
type TopLevelCommand struct {
	MainCommand *Command `yaml:",omitempty"`
	SubCommands Commands `yaml:",omitempty"`
}

type Options []*Option
type Option struct {
	Name         string `yaml:",omitempty"`
	Shorthand    string `yaml:",omitempty"`
	DefaultValue string `yaml:"default_value,omitempty"`
	Usage        string `yaml:",omitempty"`
}

type Commands []*Command
type Command struct {
	Name             string   `yaml:",omitempty"`
	Path             string   `yaml:",omitempty"`
	Synopsis         string   `yaml:",omitempty"`
	Description      string   `yaml:",omitempty"`
	Options          Options  `yaml:",omitempty"`
	InheritedOptions Options  `yaml:"inherited_options,omitempty"`
	Example          string   `yaml:",omitempty"`
	SeeAlso          []string `yaml:"see_also,omitempty"`
	Usage            string   `yaml:",omitempty"`
}

type Manifest struct {
	Docs      []Doc  `json:"docs,omitempty"`
	Title     string `json:"title,omitempty"`
	Copyright string `json:"copyright,omitempty"`
}

type Doc struct {
	Filename string `json:"filename,omitempty"`
}

type ToC struct {
	Categories []Category `yaml:",omitempty"`
}

type Category struct {
	Name     string   `yaml:",omitempty"`
	Commands []string `yaml:",omitempty"`
	Include  string   `yaml:",omitempty"`
}
