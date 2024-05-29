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

	pipeline := generator.NewFromGithubActionsWorkflow(&workflowObj)
	pipeline.GenerateTask()
	if err := pipeline.WriteResources(); err != nil {
		log.Fatalf(err.Error())
	}

	cluster := orchestration.FromKubeConfig("./resources/kubeconfig").SetupCluster().WithTekton()
	if len(cluster.GetErrors()) != 0 {
		log.Fatalf(cluster.GetErrors()[len(cluster.GetErrors())-1].Error())
	}

	file, err := os.ReadFile("./resources/test.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	cluster.ApplyTektonTask(file)
}
