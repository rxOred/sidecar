package generator

import (
	"fmt"
	"github.com/rxored/sidecar/parser"
)

type Step struct {
	Name     string
	Commands []string
	Metadata map[string]any
}

type Pipeline struct {
	Name  string
	Steps []Step
}

type TaskGenerator interface {
	GeneratePipeline(def *parser.PipelineDefinition) (*Pipeline, error)
}

type TektonTaskGenerator struct{}

func (g *TektonTaskGenerator) GeneratePipeline(def *parser.PipelineDefinition) (*Pipeline, error) {
	var steps []Step

	for _, step := range def.Steps {
		s := Step{
			Name:     step.Name,
			Commands: step.Script,
			Metadata: map[string]any{
				"platform": step.Platform,
			},
		}

		switch step.Platform {
		case "bitbucket":
			s.Metadata["artifacts"] = step.Artifacts
			s.Metadata["services"] = step.Services
			s.Metadata["caches"] = step.Caches
		case "github":
			if jobID, ok := step.Metadata["job_id"]; ok {
				s.Metadata["job_id"] = jobID
			}
		}

		steps = append(steps, s)
	}

	if len(steps) == 0 {
		return nil, fmt.Errorf("no steps generated: input pipeline had 0 steps")
	}

	pipeline := &Pipeline{
		Name:  "default", // can be extended based on parser input
		Steps: steps,
	}

	return pipeline, nil
}
