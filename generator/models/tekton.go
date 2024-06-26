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

type TektonParams struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type TekTonPvc struct {
	ClaimName string `yaml:"claimName"`
}

type TektonWorkspaceWithPvc struct {
	Name                  string    `yaml:"name"`
	PersistentVolumeClaim TekTonPvc `yaml:"persistentVolumeClaim"`
}
