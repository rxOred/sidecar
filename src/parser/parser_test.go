package parser

import (
	"testing"
)

// Unit test for interface compliance
func TestInterfaceCompliance(t *testing.T) {
	var _ PipelineParser = &BitbucketParser{}
	var _ PipelineParser = &GitHubParser{}
}

// Unit test for basic PipelineDefinition usage
func TestPipelineDefinition_Empty(t *testing.T) {
	def := &PipelineDefinition{}
	if def == nil {
		t.Error("Expected PipelineDefinition to be initialized")
	}
	if len(def.Steps) != 0 {
		t.Errorf("Expected 0 steps, got %d", len(def.Steps))
	}
}

func TestPipelineStep_MetadataAccess(t *testing.T) {
	step := PipelineStep{
		Name:     "Deploy",
		Metadata: map[string]any{"retry": 3, "timeout": "30s"},
	}

	retry, ok := step.Metadata["retry"].(int)
	if !ok || retry != 3 {
		t.Errorf("Expected retry=3, got %v", step.Metadata["retry"])
	}

	timeout, ok := step.Metadata["timeout"].(string)
	if !ok || timeout != "30s" {
		t.Errorf("Expected timeout=30s, got %v", step.Metadata["timeout"])
	}
}
