package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"github.com/rwcosta/folderz/pkg/utils"
)

// Function to read YAML file
func GetConfig(path ...string) (Config, error) {
	if len(path) <= 0 {
		return Config{}, errors.New("config YAML file not specified")
	}

	yamlFile, err := os.ReadFile(utils.MakeItYAML(path[0]))
	if err != nil {
		return Config{}, err
	}

	config := Config{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
