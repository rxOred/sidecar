package utils

import (
	"strings"

	"github.com/rxored/sidecar/generator/models"
)

func MapEnv(env map[string]string) []models.TektonEnvVar {
	var vars []models.TektonEnvVar
	for key, value := range env {
		vars = append(vars, models.TektonEnvVar{
			Name:  key,
			Value: value,
		})
	}
	return vars
}

func ToTektonTaskName(taskName string) string {
	return strings.ReplaceAll(strings.ToLower(taskName), " ", "-")
}
