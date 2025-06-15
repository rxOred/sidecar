package parser

import "fmt"

type PipelineStep struct {
	Name       string
	Script     []string
	Artifacts  []string
	Services   []string
	Caches     []string
	Deployment string
	Platform   string
	Metadata   map[string]any
}

type PipelineDefinition struct {
	Steps []PipelineStep
	Env   map[string]string
}

type PipelineParser interface {
	ParseYAML([]byte) (*PipelineDefinition, error)
	PlatformName() string
}

func GetParser(platform string) (PipelineParser, error) {
	switch platform {
	case "bitbucket":
		return &BitbucketParser{}, nil
	case "github":
		return &GitHubParser{}, nil
	default:
		return nil, fmt.Errorf("unsupported platform: %s", platform)
	}
}
