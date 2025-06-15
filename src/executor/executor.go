package executor

import (
	"errors"
	"fmt"
)

type PipelineExecutor interface {
	SetupCluster() error
	WithDependencies() error
	ApplyPipeline(yamls []string) error
	MonitorExecution() error
	Name() string
}

func GetExecutor(name string) (PipelineExecutor, error) {
	switch name {
	case "tekton":
		return &TektonExecutor{}, nil
	case "custom":
		// return &CustomExecutor{}, nil
		return nil, errors.New("custom executor not yet implemented")
	default:
		return nil, fmt.Errorf("unsupported executor backend: %s", name)
	}
}
