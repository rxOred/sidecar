package parser

import (
	"reflect"
	"testing"
)

func TestBitbucketParser(t *testing.T) {
	yaml := `
    pipelines:
      default:
        - step:
            name: "Build"
            script:
              - echo "Building..."
            artifacts:
              - dist/**
            services:
              - docker
            caches:
              - node
    `
	parser := &BitbucketParser{}
	parsed, err := parser.ParseYAML([]byte(yaml))
	if err != nil {
		t.Fatalf("Failed to parse Bitbucket YAML: %v", err)
	}

	expected := PipelineDefinition{
		Steps: []PipelineStep{
			{
				Name:      "Build",
				Script:    []string{"echo \"Building...\""},
				Artifacts: []string{"dist/**"},
				Services:  []string{"docker"},
				Caches:    []string{"node"},
				Platform:  "bitbucket",
				Metadata: map[string]any{
					"type": "step",
				},
			},
		},
	}

	if !reflect.DeepEqual(parsed.Steps, expected.Steps) {
		t.Errorf("Parsed Bitbucket pipeline does not match expected.\nGot: %+v\nExpected: %+v", parsed.Steps, expected.Steps)
	}
}
