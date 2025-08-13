package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/alfariiizi/vandor/cmd/utils"
)

type TemplateData struct {
	ModuleName string
	PascalName string
	SnakeName  string
}

// Convert PascalCase or camelCase to snake_case
func toSnakeCase(s string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}

// Normalize first letter uppercase for PascalCase
func toPascalCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/vandor/cron_new.go <JobName>")
		os.Exit(1)
	}

	input := os.Args[1]
	pascalName := toPascalCase(input)
	snakeName := toSnakeCase(input)
	fileName := filepath.Join("internal/cron/scheduler", snakeName+".go")

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		fmt.Println("[ERROR] File already exists:", fileName)
		os.Exit(1)
	}

	// Create directory if not exists
	_ = os.MkdirAll(filepath.Dir(fileName), os.ModePerm)

	// Create file from template
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tmplPath := "cmd/scheduler/cmd-new-scheduler/templates/scheduler.tmpl"
	tmplBytes, _ := os.ReadFile(tmplPath)
	tmpl, err := template.New("scheduler").Parse(string(tmplBytes))
	if err != nil {
		panic(err)
	}

	data := TemplateData{
		ModuleName: utils.GetModuleName(),
		PascalName: pascalName,
		SnakeName:  snakeName,
	}

	if err := tmpl.Execute(f, data); err != nil {
		panic(err)
	}

	fmt.Println("[OK] Created scheduler job:", fileName)
}
