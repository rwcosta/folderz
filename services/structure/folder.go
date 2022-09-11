package structure

import (
    "errors"
    "os"

    "github.com/rwcosta/folderz/pkg/config"
)

// Function that execute all the folder struct creation instructions
func mountFolderStructure() error {
    app, err := readYAML()
    if err != nil {
        return err
    }

    return byMainDirectories(app)
}

func readYAML() (config.Config, error) {
    return config.GetConfig()
}

func byMainDirectories(app config.Config) error {
    for _, d := range app.Directories {
        if err := createFolders(app.FolderStructure, os.Args[1]+"/"+d); err != nil {
            return err
        }
    }

    return nil
}

func createFolders(app map[string]interface{}, initialPath string) error {
    for k, v := range app {
        if v != nil {
            // TypeAssertion
            newFolderStructure, ok := v.(map[string]interface{})
            if !ok {
                return errors.New("error reading " + k + " folder structure")
            }

            if err := createFolders(newFolderStructure, initialPath+"/"+k); err != nil {
                return err
            }
        } else {
            if err := os.MkdirAll(initialPath+"/"+k, os.ModePerm); err != nil {
                return err
            }

            continue
        }
    }

    return nil
}
