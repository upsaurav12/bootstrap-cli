package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ProjectYAML struct {
	Name     string   `json:"name" yaml:"name"`
	Layers   []string `json:"layers" yaml:"layers"`
	Features []string `json:"features" yaml:"features"`
}

func ReadYAML(yamlPath string) (*ProjectYAML, error) {
	yamlByte, err := os.ReadFile(yamlPath)
	if err != nil {
		return &ProjectYAML{}, err
	}

	var yamlProject ProjectYAML

	err = yaml.Unmarshal(yamlByte, &yamlProject)
	if err != nil {
		fmt.Println("error: ", err)
		return &ProjectYAML{}, err
	}

	return &yamlProject, nil
}
