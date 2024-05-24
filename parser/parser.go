package parser

import (
	"github.com/rxored/sidecar/parser/models"
	"gopkg.in/yaml.v2"
)

func ParseGithubActionsWorkflow(workflowFile []byte) (models.GitHubActionsWorkflow, error) {
	var workflowObj models.GitHubActionsWorkflow
	err := yaml.Unmarshal(workflowFile, &workflowObj)
	return workflowObj, err
}
