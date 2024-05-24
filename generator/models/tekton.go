package models

type TektonMetadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace,omitempty"`
}

type TektonEnvVar struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type TektonWorkspace struct {
	Name string `yaml:"name"`
}

type TektonWorkspaceBinding struct {
	Name      string `yaml:"name"`
	Workspace string `yaml:"workspace"`
}

type TektonParamSpec struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Description string `yaml:"description,omitempty"`
	Default     string `yaml:"default,omitempty"`
}

type TektonTaskResult struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
}
