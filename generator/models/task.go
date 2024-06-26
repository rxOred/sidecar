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

type TektonTaskRef struct {
	Name string `yaml:"name"`
}

type TektonTaskRuns struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   TektonMetadata    `yaml:"metadata"`
	Spec       TektonTaskRunSpec `yaml:"spec"`
}

type TektonTaskRunSpec struct {
	TaskRef    TektonTaskRef            `yaml:"taskRef"`
	Params     []TektonParams           `yaml:"params,omitempty"`
	Workspaces []TektonWorkspaceWithPvc `yaml:"workspaces"`
}
