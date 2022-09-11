package config

// Struct for serialize the YAML parametrization file
type Config struct {
    FolderStructure map[string]interface{} `yaml:"folder-structure"`
}
