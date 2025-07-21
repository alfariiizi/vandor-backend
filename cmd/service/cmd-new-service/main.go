package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/alfariiizi/vandor/cmd/service/utils"
	cmdutils "github.com/alfariiizi/vandor/cmd/utils"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <group> <ServiceName>")
	}

	group := os.Args[1]
	name := os.Args[2]
	receiver := strings.ToLower(name[:1]) + name[1:]
	groupDir := filepath.Join("internal/core/service", group)
	targetFile := filepath.Join(groupDir, fmt.Sprintf("%s.go", receiver))
	err := os.MkdirAll("internal/core/service/"+group, 0o755)
	if err != nil {
		log.Fatalf("failed to create service dir: %v", err)
	}

	// Create service file
	renderTemplate(
		targetFile,
		filepath.Join("cmd", "service", "cmd-new-service", "service.tmpl"),
		map[string]string{
			"ModuleName":  cmdutils.GetModuleName(),
			"ServiceName": group,
			"Name":        name,
			"Receiver":    receiver,
		},
	)

	// Create or update service.go for group
	err = utils.RegenerateGroupServiceGo()
	if err != nil {
		log.Fatalf("❌ failed to generate %s/service.go: %v", group, err)
	}

	// Regenerate top-level services.go
	err = utils.RegenerateServicesGo()
	if err != nil {
		log.Fatalf("❌ failed to regenerate services.go: %v", err)
	}

	fmt.Printf("✅ Service '%s' created under '%s', and all registries updated.\n", name, group)
}

func renderTemplate(targetPath string, templatePath string, data map[string]string) {
	tplContent, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("failed to read template file: %v", err)
	}

	tmpl := template.Must(template.New("tpl").Parse(string(tplContent)))

	f, err := os.Create(targetPath)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("failed to render template: %v", err)
	}
}
