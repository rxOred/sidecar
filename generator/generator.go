package generator

import (
	"fmt"
	"os"

	Generator "github.com/rxored/sidecar/generator/models"
	Parser "github.com/rxored/sidecar/parser/models"
	"gopkg.in/yaml.v2"
)

type TektonResourcesFromGithub struct{}

func mapEnv(env map[string]string) []Generator.TektonEnvVar {
	var vars []Generator.TektonEnvVar
	for key, value := range env {
		vars = append(vars, Generator.TektonEnvVar{
			Name:  key,
			Value: value,
		})
	}
	return vars
}

func writeYamlToFile(filename string, data interface{}) error {
	file, err := os.Create("./resources/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	return encoder.Encode(data)
}

func GenerateTektonTasks(jobs map[string]Parser.Job) {
	for jobName, job := range jobs {
		task := Generator.TektonTask{
			APIVersion: "tekton.dev/v1beta",
			Kind:       "Task",
			Metadata: Generator.TektonMetadata{
				Name: jobName,
			},
			Spec: Generator.TektonTaskSpec{},
		}

		for _, step := range job.Steps {
			var tektonStep Generator.TektonTaskStep

			if step.Uses != "" {
				switch step.Uses {
				case "actions/checkout@v1", "actions/checkout@v2", "actions/checkout@v3":
					tektonStep = Generator.TektonTaskStep{
						Name:  "checkout-and-build",
						Image: "alpine:latest",
						Script: `apk add --no-cache git make
git clone https://github.com/your-repo/your-project.git /workspace/shared-workspace
cd /workspace/shared-workspace
git checkout $GITHUB_REF
make build`,
						WorkDir: "/workspace/shared-workspace",
					}
				case "actions/upload-artifact@v2":
					tektonStep = Generator.TektonTaskStep{
						Name:  "upload-artifact",
						Image: "bash:latest",
						Script: `echo "Uploading artifact..."
# Your artifact upload script here`,
					}
				default:
					tektonStep = Generator.TektonTaskStep{
						Name:    step.Name,
						Image:   "alpine",
						Script:  fmt.Sprintf("echo 'Using %s is not yet supported'", step.Uses),
						WorkDir: "/workspace/shared-workspace",
					}
				}
			} else if step.Run != "" {
				tektonStep = Generator.TektonTaskStep{
					Name:    step.Name,
					Image:   "alpine",
					Script:  step.Run,
					Env:     mapEnv(step.Env),
					WorkDir: "/workspace/shared-workspace",
				}
			}
			task.Spec.Steps = append(task.Spec.Steps, tektonStep)
		}
		task.Spec.Workspaces = append(task.Spec.Workspaces, Generator.TektonWorkspace{Name: "shared-workspace"})
		writeYamlToFile(jobName+".yaml", task)
	}
}

/*
func (tgw *TektonResourcesFromGithub) GenerateResourcesGithub(models.GitHubActionsWorkflow) {

}*/
