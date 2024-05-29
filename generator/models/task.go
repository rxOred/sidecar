package models

type TektonTask struct {
	APIVersion string         `yaml:"apiVersion"`
	Kind       string         `yaml:"kind"`
	Metadata   TektonMetadata `yaml:"metadata"`
	Spec       TektonTaskSpec `yaml:"spec"`
}

type TektonTaskSpec struct {
	Steps      []TektonTaskStep   `yaml:"steps"`
	Workspaces []TektonWorkspace  `yaml:"workspaces,omitempty"`
	Params     []TektonParamSpec  `yaml:"params,omitempty"`
	Results    []TektonTaskResult `yaml:"results,omitempty"`
}

type TektonTaskResult struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
}

type TektonTaskStep struct {
	Name       string            `yaml:"name"`
	Image      string            `yaml:"image"`
	Command    []string          `yaml:"command,omitempty"`
	Args       []string          `yaml:"args,omitempty"`
	Env        []TektonEnvVar    `yaml:"env,omitempty"`
	WorkDir    string            `yaml:"workingDir,omitempty"`
	Script     string            `yaml:"script,omitempty"`
	Workspaces []TektonWorkspace `yaml:"workspaces,omitempty"`
}
