package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alfariiizi/vandor/cmd/http/utils"
	cmdutils "github.com/alfariiizi/vandor/cmd/utils"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: go run main.go <group> <HandlerName> <METHOD>")
	}

	group := os.Args[1]                   // system
	name := os.Args[2]                    // Health
	method := strings.ToUpper(os.Args[3]) // GET / POST / PUT etc.
	if method != "GET" && method != "POST" && method != "PUT" && method != "PATCH" && method != "DELETE" {
		log.Fatalf("Unsupported HTTP method: %s", method)
	}

	receiver := strings.ToLower(name[:1]) + name[1:]
	path := strings.ToLower(name)      // for route path
	groupTitle := strings.Title(group) // for tag/title
	targetDir := filepath.Join("internal/delivery/http/route", group)
	targetFile := filepath.Join(targetDir, receiver+".go")

	err := os.MkdirAll(targetDir, 0o755)
	if err != nil {
		log.Fatalf("failed to create handler dir: %v", err)
	}

	tplPath := filepath.Join("cmd", "http", "cmd-new-handler", "handler.tmpl")
	tplContent, err := os.ReadFile(tplPath)
	if err != nil {
		log.Fatalf("failed to read template: %v", err)
	}

	tmpl := template.Must(template.New("handler").Parse(string(tplContent)))
	f, err := os.Create(targetFile)
	if err != nil {
		log.Fatalf("failed to create handler file: %v", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, map[string]string{
		"ModuleName": cmdutils.GetModuleName(),
		"Group":      group,
		"GroupTitle": groupTitle,
		"Name":       name,
		"Receiver":   receiver,
		"Path":       path,
		"Method":     method,
	})
	if err != nil {
		log.Fatalf("failed to render handler template: %v", err)
	}

	fmt.Println("✅ Handler generated!")

	// Regenerate routes.go
	if err := utils.RegenerateRoutesGo(); err != nil {
		log.Fatalf("failed to regenerate routes.go: %v", err)
	}
}
