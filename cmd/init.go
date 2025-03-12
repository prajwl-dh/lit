package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	logger "github.com/prajwl-dh/lit/util"
)

// Initialize a lit repository
// Command : ./lit init
func Init() error {
	// Create lit repository structure
	litDir := ".lit"

	subDirs := []string{
		"objects",
		"refs/heads",
	}

	files := map[string]string{
		"HEAD":            "ref: refs/heads/main\n",
		"config":          "[core]\n\trepositoryformatversion = 0\n",
		"index":           "",
		"refs/heads/main": "",
	}

	// Create main `.lit` directory
	if _, err := os.Stat(litDir); os.IsNotExist(err) {
		// Directory does not exist, so create it
		if err := os.Mkdir(litDir, 0755); err != nil {
			logger.PrintError("Error creating .lit directory\n")
			fmt.Println("Error Message: ", err)
			fmt.Println("Could not initialize lit repository")
			fmt.Println()
			os.Exit(1)
		}
	} else {
		// Directory already exists
		logger.PrintWarning("Lit repository is already initialized. Please remove the '.lit' folder in order to re-initialize the repository.\n\n")
		os.Exit(0)
	}

	// Create sub directories
	for _, dir := range subDirs {
		path := filepath.Join(litDir, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			logger.PrintError("Error creating '%v' sub directory\n", path)
			fmt.Println("Error Message: ", err)
			fmt.Println("Could not initialize lit repository")
			fmt.Println()
			os.Exit(1)
		}
	}

	// Create files
	for file, content := range files {
		path := filepath.Join(litDir, file)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			logger.PrintError("Error creating the file: %v\n", path)
			fmt.Println("Error Message: ", err)
			fmt.Println("Could not initialize lit repository")
			fmt.Println()
			os.Exit(1)
		}
	}

	logger.PrintSuccess("Initialized empty Lit repository in: %v\n\n", litDir)

	return nil
}
