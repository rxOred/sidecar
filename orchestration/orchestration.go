package orchestration

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	tektonclientset "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
)

type cluster struct {
	kubeconfig   string
	config       *rest.Config
	namespace    string
	clientClient *kubernetes.Clientset
	tektonClient *tektonclientset.Clientset
	err          []error
}

func FromKubeConfig(kubeconfig string) *cluster {
	c := &cluster{}
	c.namespace = "default"
	c.kubeconfig = kubeconfig
	return c
}

func (c *cluster) WithNamespace(namespace string) *cluster {
	c.namespace = namespace
	return c
}

func (c *cluster) SetupCluster() *cluster {
	config, err := clientcmd.BuildConfigFromFlags("", c.kubeconfig)
	if err != nil {
		c.err = append(c.err, err)
		return c
	}

	c.config = config
	c.clientClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		c.err = append(c.err, err)
		return c
	}

	return c
}

func (c *cluster) WithTekton() *cluster {
	// check if tekton has installed
	if c.err == nil {
		tc, err := tektonclientset.NewForConfig(c.config)
		if err != nil {
			c.err = append(c.err, err)
		}
		c.tektonClient = tc
	}
	return c
}

func (c *cluster) GetErrors() []error {
	return c.err
}

func (c *cluster) ApplyTektonTaskRun(taskYaml []byte) error {
	tektonScheme := runtime.NewScheme()

	decoder := serializer.NewCodecFactory(tektonScheme).UniversalDeserializer()
	obj, _, err := decoder.Decode(taskYaml, nil, nil)
	if err != nil {
		return fmt.Errorf("Error decoding Yaml")
	}

	task, ok := obj.(*tektonv1.TaskRun)
	if !ok {
		return fmt.Errorf("Error object is not a tekton taskrun")
	}

	taskJSON, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal task to JSON: %w", err)
	}

	// Print the formatted JSON string
	log.Printf("Parsed Task object:\n%s\n", string(taskJSON))

	if task.ObjectMeta.Name == "" {
		return fmt.Errorf("task metadata name is empty")
	}

	namespace := c.namespace
	if task.ObjectMeta.Namespace != "" {
		namespace = task.ObjectMeta.Namespace
	}

	createdTask, err := c.tektonClient.TektonV1().TaskRuns(namespace).Create(context.Background(), task, v1.CreateOptions{})
	if err != nil {
		return err
	}

	// created task details
	createdTaskJSON, err := json.MarshalIndent(createdTask, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal created task to JSON: %w", err)
	}
	log.Printf("Created Task object:\n%s\n", string(createdTaskJSON))

	return nil
	// return err

}

func (c *cluster) ApplyTektonTask(taskYaml []byte) error {
	tektonScheme := runtime.NewScheme()
	if err := tektonv1.AddToScheme(tektonScheme); err != nil {
		return fmt.Errorf("Error adding Tekton scheme: %v\n", err)
	}

	decoder := serializer.NewCodecFactory(tektonScheme).UniversalDeserializer()
	obj, _, err := decoder.Decode(taskYaml, nil, nil)
	if err != nil {
		return fmt.Errorf("Error decoding YAML: %v\n", err)
	}

	task, ok := obj.(*tektonv1.Task)
	if !ok {
		return fmt.Errorf("Error: object is not a Tekton Task")
	}

	taskJSON, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal task to JSON: %w", err)
	}

	// Print the formatted JSON string
	log.Printf("Parsed Task object:\n%s\n", string(taskJSON))

	if task.ObjectMeta.Name == "" {
		return fmt.Errorf("task metadata name is empty")
	}

	namespace := c.namespace
	if task.ObjectMeta.Namespace != "" {
		namespace = task.ObjectMeta.Namespace
	}

	createdTask, err := c.tektonClient.TektonV1().Tasks(namespace).Create(context.Background(), task, v1.CreateOptions{})
	if err != nil {
		return err
	}

	// created task details
	createdTaskJSON, err := json.MarshalIndent(createdTask, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal created task to JSON: %w", err)
	}
	log.Printf("Created Task object:\n%s\n", string(createdTaskJSON))

	return nil
	// return err
}
