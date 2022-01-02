package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var dirList []string

func ModuleList(moduleDir string) {
	fileList, err := filepath.Glob(moduleDir)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, modules := range fileList {
			if strings.Contains(modules, "node_modules") {
				ModuleList(modules)
			}

			file, err := os.Open(modules)
			if err != nil {
				fmt.Printf("File error: %s\n", err)
				return
			}

			fileInfo, err := file.Stat()
			if err != nil {
				fmt.Printf("FileInfo error: %s\n", err)
			}
			if fileInfo.IsDir() {
				dirList = append(dirList, modules)
			}
		}		
	}
}

func ManipulatePath() {
	// dirs := []string{"home", "nano", "go-projects", "..", "test"}
	// path := filepath.Join(dirs...)
	// fmt.Printf("The path after join: %s\n", path)
}

func main() {
    fmt.Println("Hello, World!")
	// ManipulatePath()
	ModuleList("./**/node_modules")
}