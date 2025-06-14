package parser

import (
	"reflect"
	"testing"
)

func TestGitHubParser(t *testing.T) {
	yaml := `
    jobs:
      build:
        name: "Build Job"
        steps:
          - name: "Setup"
            run: echo "Setting up"
          - run: echo "Building..."
    `
	parser := &GitHubParser{}
	parsed, err := parser.ParseYAML([]byte(yaml))
	if err != nil {
		t.Fatalf("Failed to parse GitHub YAML: %v", err)
	}

	expected := PipelineDefinition{
		Steps: []PipelineStep{
			{
				Name:     "Setup",
				Script:   []string{"echo \"Setting up\""},
				Platform: "github",
				Metadata: map[string]any{"job_id": "build"},
			},
			{
				Name:     "build_step",
				Script:   []string{"echo \"Building...\""},
				Platform: "github",
				Metadata: map[string]any{"job_id": "build"},
			},
		},
	}

	if !reflect.DeepEqual(parsed.Steps, expected.Steps) {
		t.Errorf("Parsed GitHub pipeline does not match expected.\nGot: %+v\nExpected: %+v", parsed.Steps, expected.Steps)
	}
}
