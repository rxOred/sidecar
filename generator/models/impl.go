package models

import (
	Parser "github.com/rxored/sidecar/parser/models"
)

type TektonPipeline interface {
	GeneratePipelineRun()
	GeneratePipeline()
	GenerateTask()
	extractStep()
	WriteResources()
}

type fromGithubWorkflow struct {
	TektonTasks []TektonTask
	Workflow    *Parser.GitHubActionsWorkflow
}
