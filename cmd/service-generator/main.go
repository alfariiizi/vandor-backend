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
	serviceName := strings.ToLower(args[0])
	name := args[1]
	receiver := strings.ToLower(name[:1]) + name[1:]

	targetDir := fmt.Sprintf("internal/core/service/%s", serviceName)
	targetFile := fmt.Sprintf("%s/%s.go", targetDir, receiver)

	err := os.MkdirAll(targetDir, 0755)
	if err != nil {
		log.Fatalf("failed to create service dir: %v", err)
	}

	tplPath := filepath.Join("cmd", "service-generator", "service.tmpl")
	tplContent, err := os.ReadFile(tplPath)
	if err != nil {
		log.Fatalf("failed to read template: %v", err)
	}

	tmpl := template.Must(template.New("service").Parse(string(tplContent)))
	f, err := os.Create(targetFile)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, map[string]string{
		"ServiceName": serviceName,
		"Name":        name,
		"Receiver":    receiver,
	})
	if err != nil {
		log.Fatalf("failed to render template: %v", err)
	}

	fmt.Printf("✅ service '%s' created at %s\n", name, targetFile)
}
