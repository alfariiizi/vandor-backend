package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// LoadConfigFile tries to load a config file (e.g., agent.yaml, task.yaml)
// It first checks the dev path (/config/llm), then the docker path (/app/llm).
func LoadLLMConfigFile(name, filename string) (string, error) {
	// Define possible base paths
	paths := []string{
		"config/llm", // local dev
		"app/llm",    // docker
	}

	for _, base := range paths {
		fullPath := filepath.Join(base, name, filename)
		if stat, err := os.Stat(fullPath); err == nil {
			log.Printf("Found file: %s (size: %d)\n", fullPath, stat.Size())
			return fullPath, nil
		} else {
			log.Printf("Not found at %s: %v", fullPath, err)
		}
	}

	return "", fmt.Errorf("config file %s/%s not found in any base path", name, filename)
}
