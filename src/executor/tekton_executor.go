package executor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonclientset "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type TektonExecutor struct {
	kubeconfig   string
	namespace    string
	config       *rest.Config
	kubeClient   *kubernetes.Clientset
	tektonClient *tektonclientset.Clientset
	errors       []error
}

func NewTektonExecutor() *TektonExecutor {
	return &TektonExecutor{
		namespace: "default",
	}
}

func (t *TektonExecutor) WithKubeconfig(kubeconfig string) *TektonExecutor {
	t.kubeconfig = kubeconfig
	return t
}

func (t *TektonExecutor) WithNamespace(namespace string) *TektonExecutor {
	t.namespace = namespace
	return t
}

func (t *TektonExecutor) SetupCluster() error {
	fmt.Println("[Executor] Setting up cluster via client-go...")
	config, err := clientcmd.BuildConfigFromFlags("", t.kubeconfig)
	if err != nil {
		t.errors = append(t.errors, err)
		return err
	}
	t.config = config

	t.kubeClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		t.errors = append(t.errors, err)
		return err
	}

	t.tektonClient, err = tektonclientset.NewForConfig(config)
	if err != nil {
		t.errors = append(t.errors, err)
		return err
	}

	return nil
}

func (t *TektonExecutor) WithDependencies() error {
	fmt.Println("[Executor] Checking for dependencies...")
	// Helm logic or tekton installation can be added here
	return nil
}

func (t *TektonExecutor) ApplyPipeline(yamls []string) error {
	fmt.Println("[Executor] Applying Tekton YAMLs...")
	tektonScheme := runtime.NewScheme()
	_ = tektonv1.AddToScheme(tektonScheme)
	_ = tektonv1beta1.AddToScheme(tektonScheme)
	decoder := serializer.NewCodecFactory(tektonScheme).UniversalDeserializer()

	for i, y := range yamls {
		obj, _, err := decoder.Decode([]byte(y), nil, nil)
		if err != nil {
			return fmt.Errorf("failed to decode YAML %d: %w", i+1, err)
		}

		switch o := obj.(type) {
		case *tektonv1.Task:
			_, err := t.tektonClient.TektonV1().Tasks(t.namespace).Create(context.Background(), o, metav1.CreateOptions{})
			if err != nil {
				return fmt.Errorf("failed to apply Task %s: %w", o.Name, err)
			}
			log.Printf("Task %s applied", o.Name)
		case *tektonv1.Pipeline:
			_, err := t.tektonClient.TektonV1().Pipelines(t.namespace).Create(context.Background(), o, metav1.CreateOptions{})
			if err != nil {
				return fmt.Errorf("failed to apply Pipeline %s: %w", o.Name, err)
			}
			log.Printf("Pipeline %s applied", o.Name)
		case *tektonv1.PipelineRun:
			_, err := t.tektonClient.TektonV1().PipelineRuns(t.namespace).Create(context.Background(), o, metav1.CreateOptions{})
			if err != nil {
				return fmt.Errorf("failed to apply PipelineRun %s: %w", o.Name, err)
			}
			log.Printf("PipelineRun %s applied", o.Name)
		default:
			js, _ := json.MarshalIndent(obj, "", "  ")
			log.Printf("[WARN] Unknown Tekton object: %s", js)
		}
	}
	return nil
}

func (t *TektonExecutor) MonitorExecution() error {
	fmt.Println("[Executor] Monitoring execution (placeholder)...")
	// Optionally poll PipelineRun status
	return nil
}

func (t *TektonExecutor) Name() string {
	return "Tekton"
}

func (t *TektonExecutor) Errors() []error {
	return t.errors
}
