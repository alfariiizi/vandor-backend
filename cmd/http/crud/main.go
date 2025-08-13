package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/alfariiizi/vandor/cmd/utils"
)

var templates = []string{
	"list_handler.tmpl",
	"detail_handler.tmpl",
	"create_handler.tmpl",
	"update_handler.tmpl",
	"delete_handler.tmpl",
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run ./cmd/http/crud <ModelName>")
	}
	model := os.Args[1]
	module := utils.GetModuleName()

	// Derive common naming formats
	ctx := map[string]string{
		"Model":            model,
		"ModelLower":       strings.ToLower(model),
		"ModelPlural":      model + "s",
		"ModelPluralLower": strings.ToLower(model) + "s",
		"SnakeCase":        strings.ToLower(model),
		"PluralKebab":      toKebabCase(model) + "s",
		"Module":           module,
	}

	outputDir := filepath.Join("internal", "delivery", "http", "route", ctx["SnakeCase"])
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("failed to create handler folder: %v", err)
	}

	for _, tmpl := range templates {
		outFile := strings.Replace(tmpl, "_handler.tmpl", ".go", 1)
		outPath := filepath.Join(outputDir, outFile)
		if err := renderTemplate(tmpl, outPath, ctx); err != nil {
			log.Fatalf("failed to render %s: %v", tmpl, err)
		}
		log.Printf("Generated %s", outPath)
	}
}

func renderTemplate(tmplPath, outputPath string, data any) error {
	tmpl, err := template.ParseFiles(filepath.Join("cmd", "http", "crud", "templates", tmplPath))
	if err != nil {
		return err
	}
	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()
	return tmpl.Execute(out, data)
}

func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_', r+('a'-'A'))
		} else {
			result = append(result, r)
		}
	}
	return strings.ToLower(string(result))
}

func toKebabCase(s string) string {
	return strings.ReplaceAll(toSnakeCase(s), "_", "-")
}
