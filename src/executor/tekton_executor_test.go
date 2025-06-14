package executor

import (
	"testing"
)

func TestNewTektonExecutor_Defaults(t *testing.T) {
	ex := NewTektonExecutor()
	if ex.namespace != "default" {
		t.Errorf("expected default namespace, got %s", ex.namespace)
	}
}

func TestWithNamespace(t *testing.T) {
	ex := NewTektonExecutor().WithNamespace("custom")
	if ex.namespace != "custom" {
		t.Errorf("WithNamespace did not set namespace correctly, got %s", ex.namespace)
	}
}

func TestWithKubeconfig(t *testing.T) {
	path := "/fake/path/kubeconfig"
	ex := NewTektonExecutor().WithKubeconfig(path)
	if ex.kubeconfig != path {
		t.Errorf("WithKubeconfig did not set kubeconfig correctly, got %s", ex.kubeconfig)
	}
}

func TestName(t *testing.T) {
	ex := NewTektonExecutor()
	if ex.Name() != "Tekton" {
		t.Errorf("expected Name() to return Tekton, got %s", ex.Name())
	}
}

func TestErrors_DefaultEmpty(t *testing.T) {
	ex := NewTektonExecutor()
	if len(ex.Errors()) != 0 {
		t.Errorf("expected no errors, got %d", len(ex.Errors()))
	}
}
