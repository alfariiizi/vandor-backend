// Package utils provides utility functions for the command application.
package utils

import (
	"bytes"
	"os/exec"
)

func GetModuleName() string {
	cmd := exec.Command("go", "list", "-m")
	modNameBytes, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	moduleName := string(bytes.TrimSpace(modNameBytes))
	return moduleName
}
