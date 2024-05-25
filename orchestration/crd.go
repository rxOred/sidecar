package orchestration

/*

package orchestration

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-github/v39/github"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func DownloadTektonReleaseAssets(kubeconfig string) {
	ctx := context.Background()
	client := github.NewClient(nil)

	// Get latest Tekton release
	release, _, err := client.Repositories.GetLatestRelease(ctx, "tektoncd", "pipeline")
	if err != nil {
		log.Fatalf("Error getting latest Tekton release: %v", err)
	}

	for _, asset := range release.Assets {
		if filepath.Ext(asset.GetName()) == ".yaml" {
			resp, err := http.Get(asset.GetBrowserDownloadURL())
			if err != nil {
				log.Fatalf("Error downloading file %s: %v", asset.GetName(), err)
			}
			defer resp.Body.Close()

			filePath := filepath.Join(".", asset.GetName())
			file, err := os.Create(filePath)
			if err != nil {
				log.Fatalf("Error creating file %s: %v", filePath, err)
			}
			defer file.Close()

			_, err = io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Error reading response body: %v", err)
			}

			// Apply CRD to Kubernetes cluster
			err = applyCRD(filePath, kubeconfig)
			if err != nil {
				log.Fatalf("Error applying CRD %s: %v", filePath, err)
			}
			log.Printf("CRD %s applied successfully\n", filePath)
		}
	}

}

// applyCRD applies a CRD YAML file to the Kubernetes cluster
func applyCRD(filePath string, kubeconfig string) error {
	// Create Kubernetes client config
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return fmt.Errorf("error creating Kubernetes client config: %v", err)
	}

	// Create Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating Kubernetes clientset: %v", err)
	}

	// Read CRD YAML file
	crdYAML, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading CRD YAML file: %v", err)
	}

	// Apply CRD to Kubernetes cluster
	_, err = clientset.Discovery().ServerResourcesForGroupVersion("apiextensions.k8s.io/v1").List(context.Background())
	if err != nil {
		return fmt.Errorf("error listing API resources: %v", err)
	}

	return nil
} */
