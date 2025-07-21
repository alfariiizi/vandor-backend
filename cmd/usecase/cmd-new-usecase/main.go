package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/alfariiizi/vandor/cmd/usecase/utils"
	cmdutils "github.com/alfariiizi/vandor/cmd/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("❌ Missing usecase name argument")
	}
	name := args[0]
	if name == "" {
		log.Fatal("❌ Usecase name cannot be empty")
	}
	if lowerName := strings.ToLower(name); lowerName == "usecases" || lowerName == "usecase" {
		log.Fatal("❌ Usecase name cannot be 'usecase' or 'usecases'")
	}
	receiver := strings.ToLower(name[:1]) + name[1:]
	targetFile := fmt.Sprintf("internal/core/usecase/%s.go", receiver)

	// Step 1: Create the new usecase file
	createUsecaseFile(name, receiver, targetFile)

	// Step 2: Regenerate usecases.go
	err := utils.RegenerateUsecasesGo()
	if err != nil {
		log.Fatalf("❌ Failed to update usecases.go: %v", err)
	}

	fmt.Printf("✅ Usecase '%s' created and usecases.go updated.\n", name)
}

func createUsecaseFile(name, receiver, targetFile string) {
	err := os.MkdirAll("internal/core/usecase", 0o755)
	if err != nil {
		log.Fatalf("failed to create usecase dir: %v", err)
	}

	tplPath := filepath.Join("cmd", "usecase", "cmd-new-usecase", "usecase.tmpl")
	tplContent, err := os.ReadFile(tplPath)
	if err != nil {
		log.Fatalf("failed to read template: %v", err)
	}

	tmpl := template.Must(template.New("usecase").Parse(string(tplContent)))
	f, err := os.Create(targetFile)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, map[string]string{
		"ModuleName": cmdutils.GetModuleName(),
		"Name":       name,
		"Receiver":   receiver,
	})
	if err != nil {
		log.Fatalf("failed to render template: %v", err)
	}
}
