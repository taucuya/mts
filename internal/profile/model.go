package profile

type Profile struct {
	User    string `yaml:"user"`
	Project string `yaml:"project"`
}

type NamedProfile struct {
	Name    string
	User    string
	Project string
}
