package config

import (
    "os"
    "path/filepath"

    "gopkg.in/yaml.v3"
)

// Function to read YAML file
func GetConfig() (Config, error) {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }

    yamlFile, err := os.ReadFile(filepath.Dir(ex) + "/structure.yml")
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
