package orchestration

import (
	"context"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonclientset "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
)

func SetupK8sCluser() {
	var kubeconfig string
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf(err.Error())
	}

	kubeconfig = cwd + "/resources/kubeconfig"
	log.Println("loading kubeconfig from " + kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		config, err = rest.InClusterConfig()
		log.Fatalf(err.Error())
	}

	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	tektonClient, err := tektonclientset.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Tekton client: %v", err)
	}

	taskYaml, err := os.ReadFile("./resources/build.yaml")
	if err != nil {
		log.Fatalf("Error reading task YAML file: %v", err)
	}

	var task tektonv1beta1.Task
	if err := yaml.Unmarshal(taskYaml, &task); err != nil {
		log.Fatalf("Error parsing task YAML: %v", err)
	}

	// Apply the Task to the cluster
	namespace := "default"
	if task.ObjectMeta.Namespace != "" {
		namespace = task.ObjectMeta.Namespace
	}

	_, err = tektonClient.TektonV1beta1().Tasks(namespace).Create(context.Background(), &task, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error applying Task to the cluster: %v", err)
	}

	fmt.Println("Tekton Task applied successfully")
}
