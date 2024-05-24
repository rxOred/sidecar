package main

import (
	"log"
	"os"

	"github.com/rxored/sidecar/generator"
	"github.com/rxored/sidecar/orchestration"
	"github.com/rxored/sidecar/parser"
)

func main() {
	pipelineFile, err := os.ReadFile("samples/github.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	workflowObj, err := parser.ParseGithubActionsWorkflow(pipelineFile)
	if err != nil {
		log.Fatalf(err.Error())
	}

	generator.GenerateTektonTasks(workflowObj.Jobs)
	orchestration.SetupK8sCluser()
}
