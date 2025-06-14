package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rxored/sidecar/executor"
	"github.com/rxored/sidecar/generator"
	"github.com/rxored/sidecar/parser"
)

func main() {
	platform := flag.String("platform", "bitbucket", "The CI platform to parse (e.g., bitbucket, github)")
	pipelinePath := flag.String("pipeline", "pipeline.yaml", "Path to the CI pipeline YAML file")
	executorBackend := flag.String("executor", "tekton", "Executor backend to use (e.g., tekton)")
	kubeconfig := flag.String("kubeconfig", os.Getenv("KUBECONFIG"), "Path to kubeconfig")
	namespace := flag.String("namespace", "default", "Kubernetes namespace to deploy into")

	flag.Parse()

	data, err := os.ReadFile(*pipelinePath)
	if err != nil {
		log.Fatalf("Failed to read pipeline YAML: %v", err)
	}

	p, err := parser.GetParser(*platform)
	if err != nil {
		log.Fatalf("Parser error: %v", err)
	}

	parsedDef, err := p.ParseYAML(data)
	if err != nil {
		log.Fatalf("YAML parsing error: %v", err)
	}

	gen := &generator.TektonTaskGenerator{}
	pipeline, err := gen.GeneratePipeline(parsedDef)
	if err != nil {
		log.Fatalf("Pipeline generation failed: %v", err)
	}

	yamls, err := generator.RenderTektonYAML(pipeline)
	if err != nil {
		log.Fatalf("YAML rendering failed: %v", err)
	}

	fmt.Println("--- Generated Tekton YAML ---")
	for i, yml := range yamls {
		fmt.Printf("\n--- YAML Document %d ---\n%s\n", i+1, yml)
	}

	exec, err := executor.GetExecutor(*executorBackend)
	if err != nil {
		log.Fatalf("Executor init failed: %v", err)
	}

	execTyped, ok := exec.(*executor.TektonExecutor)
	if ok {
		execTyped.WithKubeconfig(*kubeconfig).WithNamespace(*namespace)
	}

	if err := exec.SetupCluster(); err != nil {
		log.Fatalf("Cluster setup failed: %v", err)
	}

	if err := exec.WithDependencies(); err != nil {
		log.Fatalf("Dependency installation failed: %v", err)
	}

	if err := exec.ApplyPipeline(yamls); err != nil {
		log.Fatalf("Pipeline apply failed: %v", err)
	}

	if err := exec.MonitorExecution(); err != nil {
		log.Fatalf("Monitoring failed: %v", err)
	}

	fmt.Println("✅ Pipeline execution complete")
}
