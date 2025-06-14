package generator

import (
	"strings"
	"testing"
)

func TestRenderTektonYAML_BasicPipeline(t *testing.T) {
	pipeline := &Pipeline{
		Name: "sample-pipeline",
		Steps: []Step{
			{
				Name:     "step-1",
				Commands: []string{"echo Hello from Step 1"},
				Metadata: map[string]any{"platform": "bitbucket"},
			},
			{
				Name:     "step-2",
				Commands: []string{"echo Hello from Step 2"},
				Metadata: map[string]any{"platform": "github"},
			},
		},
	}

	yamls, err := RenderTektonYAML(pipeline)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(yamls) != 4 { // 2 tasks + 1 pipeline + 1 pipelinerun
		t.Errorf("Expected 4 YAML docs (2 tasks + pipeline + pipelinerun), got %d", len(yamls))
	}

	// Check that task YAML contains expected names and commands
	for i, y := range yamls[:2] {
		if !strings.Contains(y, pipeline.Steps[i].Name) {
			t.Errorf("Task YAML %d missing step name: %s", i+1, pipeline.Steps[i].Name)
		}
		if !strings.Contains(y, pipeline.Steps[i].Commands[0]) {
			t.Errorf("Task YAML %d missing step command: %s", i+1, pipeline.Steps[i].Commands[0])
		}
	}

	// Check pipeline YAML includes task references
	if !strings.Contains(yamls[2], "Pipeline") {
		t.Errorf("Pipeline YAML missing kind: Pipeline")
	}
	if !strings.Contains(yamls[2], pipeline.Steps[0].Name) {
		t.Errorf("Pipeline YAML missing step reference: %s", pipeline.Steps[0].Name)
	}

	// Check PipelineRun references the pipeline
	if !strings.Contains(yamls[3], "PipelineRun") {
		t.Errorf("PipelineRun YAML missing kind: PipelineRun")
	}
	if !strings.Contains(yamls[3], pipeline.Name) {
		t.Errorf("PipelineRun YAML missing pipeline reference: %s", pipeline.Name)
	}
}
