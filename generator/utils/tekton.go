package utils

import (
	"strings"

	"github.com/rxored/sidecar/generator/models"
	"github.com/rxored/sidecar/utils"
	"golang.org/x/text/cases"
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

func WriteResource(data interface{}) error {
	switch v := data.(type) {
	case *From:
		for _, k := range v {
			if err := utils.WriteYamlToFile(k.Metadata.Name+".yaml", k); if err != nil {
				return err
			}
		}
	}
}
