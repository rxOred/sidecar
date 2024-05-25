package main

import (
	"log"
	"os"

	"github.com/rxored/sidecar/generator"
	"github.com/rxored/sidecar/orchestration"
	"github.com/rxored/sidecar/parser"
)

func main() {
	// 0. specify whether github, gitlab or bb
	// 1. kubeconfig path provided and a k8s cluster is running with tekton installed
	// 2. kubeconfig path provided and k8s cluster is running but need to install tekton
	// 3. need to create a k8s cluster and install tekton (k3s or microk8s)

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
