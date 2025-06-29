package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	args := os.Args[1:]
	name := args[0]
	receiver := strings.ToLower(name[:1]) + name[1:]
	targetFile := fmt.Sprintf("internal/core/usecase/%s.go", receiver)

	err := os.MkdirAll("internal/core/usecase", 0o755)
	if err != nil {
		log.Fatalf("failed to create usecase dir: %v", err)
	}

	tplPath := filepath.Join("cmd", "usecase-generator", "usecase.tmpl")
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
		"Name":     name,
		"Receiver": receiver,
	})
	if err != nil {
		log.Fatalf("failed to render template: %v", err)
	}

	fmt.Printf("✅ Usecase '%s' created at %s\n", name, targetFile)
}
