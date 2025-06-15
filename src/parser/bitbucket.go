package parser

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type BitbucketParser struct{}

type BitbucketYAML struct {
	Pipelines struct {
		Default  []PipelineEntry            `yaml:"default"`
		Branches map[string][]PipelineEntry `yaml:"branches"`
		Tags     map[string][]PipelineEntry `yaml:"tags"`
	} `yaml:"pipelines"`
}

type PipelineEntry struct {
	Step     Step   `yaml:"step"`
	Parallel []Step `yaml:"parallel"`
}

type Step struct {
	Name       string   `yaml:"name"`
	Script     []string `yaml:"script"`
	Artifacts  []string `yaml:"artifacts,omitempty"`
	Services   []string `yaml:"services,omitempty"`
	Caches     []string `yaml:"caches,omitempty"`
	Deployment string   `yaml:"deployment,omitempty"`
}

func (b *BitbucketParser) PlatformName() string {
	return "bitbucket"
}

func (b *BitbucketParser) ParseYAML(data []byte) (*PipelineDefinition, error) {
	var parsed BitbucketYAML
	err := yaml.Unmarshal(data, &parsed)
	if err != nil {
		return nil, fmt.Errorf("bitbucket yaml unmarshal failed: %w", err)
	}

	def := &PipelineDefinition{}
	collectSteps := func(entries []PipelineEntry) {
		for _, entry := range entries {
			if entry.Step.Name != "" || len(entry.Step.Script) > 0 {
				def.Steps = append(def.Steps, PipelineStep{
					Name:       entry.Step.Name,
					Script:     entry.Step.Script,
					Artifacts:  entry.Step.Artifacts,
					Services:   entry.Step.Services,
					Caches:     entry.Step.Caches,
					Deployment: entry.Step.Deployment,
					Platform:   "bitbucket",
					Metadata: map[string]any{
						"type": "step",
					},
				})
			}
			for _, pstep := range entry.Parallel {
				def.Steps = append(def.Steps, PipelineStep{
					Name:       pstep.Name,
					Script:     pstep.Script,
					Artifacts:  pstep.Artifacts,
					Services:   pstep.Services,
					Caches:     pstep.Caches,
					Deployment: pstep.Deployment,
					Platform:   "bitbucket",
					Metadata: map[string]any{
						"type": "parallel",
					},
				})
			}
		}
	}

	collectSteps(parsed.Pipelines.Default)
	for _, entries := range parsed.Pipelines.Branches {
		collectSteps(entries)
	}
	for _, entries := range parsed.Pipelines.Tags {
		collectSteps(entries)
	}

	return def, nil
}

