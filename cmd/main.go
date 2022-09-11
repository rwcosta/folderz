package main

import "github.com/rwcosta/folderz/cmd/components"

func main() {
    for _, s := range components.Services {
        if err := s.Execute(); err != nil {
            panic(err)
        }
    }
}
