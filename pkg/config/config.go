package config

// Struct for serialize the YAML parametrization file
type Config struct {
    Directories     []string               `yaml:"directories"`
    FolderStructure map[string]interface{} `yaml:"folder-structure"`
}
