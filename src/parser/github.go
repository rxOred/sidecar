package parser

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type GitHubParser struct{}

type GitHubYAML struct {
	Jobs map[string]struct {
		Name  string `yaml:"name"`
		Steps []struct {
			Name string `yaml:"name"`
			Run  string `yaml:"run"`
		} `yaml:"steps"`
	} `yaml:"jobs"`
}

func (g *GitHubParser) PlatformName() string {
	return "github"
}

func (g *GitHubParser) ParseYAML(data []byte) (*PipelineDefinition, error) {
	var parsed GitHubYAML
	err := yaml.Unmarshal(data, &parsed)
	if err != nil {
		return nil, fmt.Errorf("github yaml unmarshal failed: %w", err)
	}

	def := &PipelineDefinition{}
	for jobID, job := range parsed.Jobs {
		for _, step := range job.Steps {
			stepName := step.Name
			if stepName == "" {
				stepName = fmt.Sprintf("%s_step", jobID)
			}
			def.Steps = append(def.Steps, PipelineStep{
				Name:     stepName,
				Script:   []string{step.Run},
				Platform: "github",
				Metadata: map[string]any{
					"job_id": jobID,
				},
			})
		}
	}

	return def, nil
}
