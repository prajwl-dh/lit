package cmd

import (
	"crypto/sha1"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	logger "github.com/prajwl-dh/lit/util"
)

func Add() error {
	ignoreList := make(map[string]bool)

	// Read from .litignore if it exists and add to ignoreList
	if data, err := os.ReadFile(".litignore"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" {
				ignoreList[trimmed] = true
			}
		}
	}

	indexPath := filepath.Join(".lit", "index")
	indexEntries := make([]string, 0)

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == "." {
			return nil
		}

		if path == ".lit" || strings.HasPrefix(path, ".lit/") || path == ".git" || strings.HasPrefix(path, ".git/") {
			return nil
		}

		if path == ".litignore" || path == ".gitignore" {
			return nil
		}

		if ignoreList[path] {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		hash := fmt.Sprintf("%x", sha1.Sum(content))

		objectPath := filepath.Join(".lit", "objects", hash)
		if _, err := os.Stat(objectPath); os.IsNotExist(err) {
			if err := os.WriteFile(objectPath, content, 0644); err != nil {
				logger.PrintError("Failed to write blob object : %v\n", err)
				return err
			}
		}

		indexEntries = append(indexEntries, fmt.Sprintf("%s %s", hash, path))

		return nil
	})

	if err != nil {
		logger.PrintError("Error walking directory: %v\n", err)
		return err
	}

	if err := os.WriteFile(indexPath, []byte(strings.Join(indexEntries, "\n")+"\n"), 0644); err != nil {
		return err
	}

	logger.PrintSuccess("Added files to index\n")
	fmt.Println()
	return nil
}
