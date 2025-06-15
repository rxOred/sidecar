package generator

import (
	"fmt"
	"sigs.k8s.io/yaml"
	"strings"

	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
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
	var tasks []*tektonv1.Task

	// Convert steps into individual Tasks
	for _, step := range pipeline.Steps {
		task := &tektonv1.Task{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Task",
				APIVersion: "tekton.dev/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: strings.ToLower(step.Name),
			},
			Spec: tektonv1.TaskSpec{
				Steps: []tektonv1.Step{{
					Name:   strings.ToLower(step.Name),
					Image:  "alpine",
					Script: fmt.Sprintf("#!/bin/sh\n%s", strings.Join(step.Commands, "\n")),
					SecurityContext: &corev1.SecurityContext{
						AllowPrivilegeEscalation: pointer.Bool(true),
						Capabilities: &corev1.Capabilities{
							Drop: []corev1.Capability{"ALL"},
						},
						SeccompProfile: &corev1.SeccompProfile{
							Type: corev1.SeccompProfileTypeRuntimeDefault,
						},
					},
				}},
			},
		}
		tasks = append(tasks, task)
	}

	// Create a Pipeline with references to the Tasks
	var pipelineTasks []tektonv1.PipelineTask
	for _, task := range tasks {
		pipelineTasks = append(pipelineTasks, tektonv1.PipelineTask{
			Name: task.Name,
			TaskRef: &tektonv1.TaskRef{
				Name: task.Name,
				Kind: tektonv1.NamespacedTaskKind,
			},
		})
	}

	pipe := &tektonv1.Pipeline{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pipeline",
			APIVersion: "tekton.dev/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: strings.ToLower(pipeline.Name),
		},
		Spec: tektonv1.PipelineSpec{
			Tasks: pipelineTasks,
		},
	}

	run := &tektonv1.PipelineRun{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PipelineRun",
			APIVersion: "tekton.dev/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: strings.ToLower(pipeline.Name + "-run"),
		},
		Spec: tektonv1.PipelineRunSpec{
			PipelineRef: &tektonv1.PipelineRef{
				Name: strings.ToLower(pipeline.Name),
			},
		},
	}

	// Marshal all resources to YAML
	allObjects := []any{}
	for _, task := range tasks {
		allObjects = append(allObjects, task)
	}
	allObjects = append(allObjects, pipe, run)

	for _, obj := range allObjects {
		out, err := yaml.Marshal(obj)
		if err != nil {
			return nil, err
		}
		yamls = append(yamls, string(out))
	}

	return yamls, nil
}
