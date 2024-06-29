package parser

import (
	"github.com/rxored/sidecar/parser/models/bitbucket"
	"github.com/rxored/sidecar/parser/models/github"
	"gopkg.in/yaml.v2"
)

func ParseGithubActionsWorkflow(workflowFile []byte) (github.GitHubActionsWorkflow, error) {
	var workflowObj github.GitHubActionsWorkflow
	err := yaml.Unmarshal(workflowFile, &workflowObj)
	return workflowObj, err
}

func ParseGitlabPipeline(pipeline []byte) {

}

func BitbucketPipeline(pipelineFile []byte) (bitbucket.BitbucketPipeline, error) {
	var pipelineObj bitbucket.BitbucketPipeline
	err := yaml.Unmarshal(pipelineFile, &pipelineObj)
	return pipelineObj, err
}
