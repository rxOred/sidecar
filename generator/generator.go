package generator

import (
	"fmt"
	"log"

	Generator "github.com/rxored/sidecar/generator/models"
	Utils "github.com/rxored/sidecar/generator/utils"
	Parser "github.com/rxored/sidecar/parser/models"
	"github.com/rxored/sidecar/utils"
)

func WriteResource(data interface{}) error {
	switch v := data.(type) {
	case *fromGithubWorkflow:
		// write all the tasks
		for _, k := range v.TektonTasks {
			err := utils.WriteYamlToFile(k.Metadata.Name+".yaml", k)
			if err != nil {
				return err
			}
		}

		for _, k := range v.TektonTaskRuns {
			err := utils.WriteYamlToFile(k.Metadata.Name+".yaml", k)
			if err != nil {
				return err
			}
		}
		// write all the pipelines

		// write all the pipeline runs
	default:
		log.Println("not a matching type ", v)
	}
	return nil
}

type TektonPipeline interface {
	GeneratePipelineRun()
	GeneratePipeline()
	GenerateTask()
	extractStep()
	WriteResources()
}

type TektonPipelineImpl struct {
	TektonTasks     []Generator.TektonTask
	TektonTaskRuns  []Generator.TektonTaskRuns
	TektonPipelines Generator.TektonPipeline
}

type fromGithubWorkflow struct {
	TektonPipelineImpl
	Workflow *Parser.GitHubActionsWorkflow
}

type fromGitlabPipeline struct {
	TektonPipelineImpl
}

type fromBitbucketPipeline struct {
	TektonPipelineImpl
}

func NewFromGithubActionsWorkflow(wf *Parser.GitHubActionsWorkflow) *fromGithubWorkflow {
	obj := &fromGithubWorkflow{Workflow: wf}
	return obj
}

func (fg *fromGithubWorkflow) WriteResources() error {
	return WriteResource(fg)
}

func (fg *fromGithubWorkflow) GeneratePipeline() {

}

func (fg *fromGithubWorkflow) extractStep(wfStep Parser.Step, tektonStep *Generator.TektonTaskStep) {
	tektonStep.Name = Utils.ToTektonTaskName(wfStep.Name)
	tektonStep.Image = "alpine" // set alpine as default for now
	tektonStep.WorkDir = "/workspace/shared-workspace"
	tektonStep.Workspaces = append(tektonStep.Workspaces, Generator.TektonWorkspace{Name: "shared-workspace"})

	// handle actions logic manually for now
	if wfStep.Uses != "" {
		switch wfStep.Uses {
		case "actions/checkout@v1", "actions/checkout@v2", "actions/checkout@v3":
		default:
			tektonStep.Script = fmt.Sprintf("echo 'Using %s is not yet supported'", wfStep.Uses)
		}
	} else if wfStep.Run != "" {
		tektonStep.Script = wfStep.Run
	}
}

func (fg *fromGithubWorkflow) GenerateTask() {
	for jobName, job := range fg.Workflow.Jobs {
		task := Generator.TektonTask{
			APIVersion: "tekton.dev/v1",
			Kind:       "Task",
			Metadata: Generator.TektonMetadata{
				Name: Utils.ToTektonTaskName(jobName),
			},
			Spec: Generator.TektonTaskSpec{},
		}
		for _, step := range job.Steps {
			var tektonStep Generator.TektonTaskStep
			fg.extractStep(step, &tektonStep)
			task.Spec.Steps = append(task.Spec.Steps, tektonStep)
		}
		task.Spec.Workspaces = append(task.Spec.Workspaces, Generator.TektonWorkspace{Name: "shared-workspace"})
		fg.TektonTasks = append(fg.TektonTasks, task)
	}
}

func (fg *fromGithubWorkflow) GenerateTaskRun() {
	for _, task := range fg.TektonTasks {
		run := Generator.TektonTaskRuns{
			APIVersion: task.APIVersion,
			Kind:       "TaskRun",
			Metadata: Generator.TektonMetadata{
				Name: task.Metadata.Name + "-taskrun",
			},
			Spec: Generator.TektonTaskRunSpec{},
		}
		run.Spec.TaskRef = Generator.TektonTaskRef{Name: task.Metadata.Name}
		// fill up params
		run.Spec.Workspaces = append(run.Spec.Workspaces, Generator.TektonWorkspaceWithPvc{
			Name: "shared-workspace",
			PersistentVolumeClaim: Generator.TekTonPvc{
				ClaimName: "pvc1",
			},
		})

		fg.TektonTaskRuns = append(fg.TektonTaskRuns, run)
	}
}
