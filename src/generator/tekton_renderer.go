package generator

import (
	"fmt"
	"sigs.k8s.io/yaml"
)

type TektonTask struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       struct {
		Steps []struct {
			Name   string `yaml:"name"`
			Image  string `yaml:"image"`
			Script string `yaml:"script"`
		} `yaml:"steps"`
	} `yaml:"spec"`
}

type TektonPipeline struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       struct {
		Tasks []struct {
			Name    string `yaml:"name"`
			TaskRef struct {
				Name string `yaml:"name"`
			} `yaml:"taskRef"`
		} `yaml:"tasks"`
	} `yaml:"spec"`
}

type TektonPipelineRun struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       struct {
		PipelineRef struct {
			Name string `yaml:"name"`
		} `yaml:"pipelineRef"`
	} `yaml:"spec"`
}

func RenderTektonYAML(pipeline *Pipeline) ([]string, error) {
	var yamls []string

	// Tasks
	for _, step := range pipeline.Steps {
		task := TektonTask{
			APIVersion: "tekton.dev/v1",
			Kind:       "Task",
			Metadata:   map[string]string{"name": step.Name},
		}
		taskStep := struct {
			Name   string `yaml:"name"`
			Image  string `yaml:"image"`
			Script string `yaml:"script"`
		}{
			Name:   step.Name,
			Image:  "alpine",
			Script: fmt.Sprintf("#!/bin/sh\n%s", step.Commands[0]), // Simplified
		}
		task.Spec.Steps = append(task.Spec.Steps, taskStep)

		out, err := yaml.Marshal(task)
		if err != nil {
			return nil, err
		}
		yamls = append(yamls, string(out))
	}

	// Pipeline
	tp := TektonPipeline{
		APIVersion: "tekton.dev/v1",
		Kind:       "Pipeline",
		Metadata:   map[string]string{"name": pipeline.Name},
	}
	for _, step := range pipeline.Steps {
		tpStep := struct {
			Name    string `yaml:"name"`
			TaskRef struct {
				Name string `yaml:"name"`
			} `yaml:"taskRef"`
		}{
			Name: step.Name,
		}
		tpStep.TaskRef.Name = step.Name
		tp.Spec.Tasks = append(tp.Spec.Tasks, tpStep)
	}

	pipelineOut, err := yaml.Marshal(tp)
	if err != nil {
		return nil, err
	}
	yamls = append(yamls, string(pipelineOut))

	// PipelineRun
	run := TektonPipelineRun{
		APIVersion: "tekton.dev/v1",
		Kind:       "PipelineRun",
		Metadata:   map[string]string{"name": pipeline.Name + "-run"},
	}
	run.Spec.PipelineRef.Name = pipeline.Name

	runOut, err := yaml.Marshal(run)
	if err != nil {
		return nil, err
	}
	yamls = append(yamls, string(runOut))

	return yamls, nil
}
