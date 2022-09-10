package structure

import (
	"errors"
	"github.com/rwcosta/folderz/pkg/config"
	"os"
)

// Function that execute all the folder struct creation instructions
func mountFolderStructure() error {
	app, err := readYAML()
	if err != nil {
		return err
	}

	return createFolders(app.FolderStructure, ".")
}

func readYAML() (config.Config, error) {
	return config.GetConfig(os.Args[1:]...)
}

func createFolders(app map[string]interface{}, initialPath string) error {
	for k, v := range app {
		if v != nil {
			newFolderStructure, ok := v.(map[string]interface{})
			if !ok {
				return errors.New("error reading " + k + " folder structure")
			}

			if err := createFolders(newFolderStructure, initialPath + "/" + k); err != nil {
				return err
			}
		} else {
			if err := os.MkdirAll(initialPath + "/" + k, os.ModePerm); err != nil {
				return err
			}

			continue
		}
	}

	return nil
}
