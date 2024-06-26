package models

type TektonPipeline struct {
	APIVersion string             `yaml:"apiVersion"`
	Kind       string             `yaml:"kind"`
	Metadata   TektonMetadata     `yaml:"metadata"`
	Spec       TektonPipelineSpec `yaml:"spec"`
}

type TektonPipelineSpec struct {
	Tasks      []TektonPipelineTask `yaml:"tasks"`
	Workspaces []TektonWorkspace    `yaml:"workspaces,omitempty"`
	Params     []TektonParamSpec    `yaml:"params,omitempty"`
}

type TektonPipelineTask struct {
	Name       string                   `yaml:"name"`
	TaskRef    TektonTaskRef            `yaml:"taskRef"`
	Params     []TektonParams           `yaml:"params,omitempty"`
	RunAfter   []string                 `yaml:"runAfter,omitempty"`
	Workspaces []TektonWorkspaceBinding `yaml:"workspaces,omitempty"`
}
