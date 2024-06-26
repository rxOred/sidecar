package models

import "fmt"

type Needs []string

func (n *Needs) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var need string
	if err := unmarshal(&need); err == nil {
		*n = Needs{need}
		return nil
	}

	var needs []string
	if err := unmarshal(&needs); err == nil {
		*n = needs
		return nil
	}

	return fmt.Errorf("failed to unmarshal Needs")
}

type On struct {
	Events       []string
	EventMapping map[string]interface{}
}

func (o *On) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var on string
	if err := unmarshal(&on); err == nil {
		o.Events = []string{on}
		return nil
	}

	var ons []string
	if err := unmarshal(&ons); err == nil {
		o.Events = ons
		return nil
	}

	var mapping map[string]interface{}
	if err := unmarshal(&mapping); err == nil {
		o.EventMapping = mapping
		return nil
	}

	return fmt.Errorf("failed to unmarshal Ons")
}

type Step struct {
	ID               string            `yaml:"id,omitempty"`
	Name             string            `yaml:"name,omitempty"`
	Uses             string            `yaml:"uses,omitempty"`
	Run              string            `yaml:"run,omitempty"`
	WorkingDirectory string            `yaml:"working-directory,omitempty"`
	Shell            string            `yaml:"shell,omitempty"`
	With             map[string]string `yaml:"with,omitempty"`
	Env              map[string]string `yaml:"env,omitempty"`
	If               string            `yaml:"if,omitempty"`
	ContinueOnError  bool              `yaml:"continue-on-error,omitempty"`
	TimeoutMinutes   int               `yaml:"timeout-minutes,omitempty"`
}

type Job struct {
	RunsOn string `yaml:"runs-on"`
	Steps  []Step `yaml:"steps"`
	Needs  Needs  `yaml:"needs,omitempty"`
}

type GitHubActionsWorkflow struct {
	Name string         `yaml:"name"`
	On   On             `yaml:"on"`
	Jobs map[string]Job `yaml:"jobs"`
}
