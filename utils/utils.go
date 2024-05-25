package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

func WriteYamlToFile(filename string, data interface{}) error {
	file, err := os.Create("./resources/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	return encoder.Encode(data)
}
