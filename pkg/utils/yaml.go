package utils

import "strings"

// Function to check if the file passed as parameter is a YAML
func IsYAML(file string) bool {
    if len(file) < 5 {
        return false
    }

    lastChars := file[len(file)-5:]
    return strings.Contains(lastChars, ".yml") || strings.Contains(lastChars, ".yaml")
}

// Function to make the program read the file as a YAML, it changes the file name if necessary
func MakeItYAML(file string) string {
    if IsYAML(file) {
        return file
    }

    return file + ".yml"
}
