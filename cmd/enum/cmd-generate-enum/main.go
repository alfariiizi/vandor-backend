package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/alfariiizi/vandor/cmd/utils"
	"gopkg.in/yaml.v3"
)

// EnumDefinition represents the structure of a YAML enum definition
type EnumDefinition struct {
	Type        string            `yaml:"type"`
	Values      map[string]string `yaml:"values"`
	Description string            `yaml:"description,omitempty"`
}

func generateEnums() {
	// Ensure output directory exists
	if err := os.MkdirAll("internal/enum", 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Find all YAML files in enum directory
	yamlFiles, err := findYAMLFiles("enum")
	if err != nil {
		fmt.Printf("Error finding YAML files: %v\n", err)
		os.Exit(1)
	}

	if len(yamlFiles) == 0 {
		fmt.Println("No YAML files found in enum directory")
		return
	}

	// Parse template
	templateFile := filepath.Join("cmd", "enum", "cmd-create-enum", "enum.tmpl")
	tmpl, err := template.New("enum").Funcs(template.FuncMap{
		"lcFirst": lcFirst,
	}).Parse(templateFile)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		os.Exit(1)
	}

	// Process each YAML file
	for _, yamlFile := range yamlFiles {
		fmt.Printf("Processing %s...\n", yamlFile)

		if err := processEnumFile(yamlFile, tmpl); err != nil {
			fmt.Printf("Error processing %s: %v\n", yamlFile, err)
			continue
		}
	}

	fmt.Printf("Successfully generated %d enum files\n", len(yamlFiles))
}

func findYAMLFiles(dir string) ([]string, error) {
	var yamlFiles []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && (strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml")) {
			yamlFiles = append(yamlFiles, path)
		}
		return nil
	})

	return yamlFiles, err
}

func processEnumFile(yamlFile string, tmpl *template.Template) error {
	// Read YAML file
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	// Parse YAML
	var enumDef EnumDefinition
	if err := yaml.Unmarshal(data, &enumDef); err != nil {
		return fmt.Errorf("parsing YAML: %w", err)
	}

	// Validate required fields
	if enumDef.Type == "" {
		return fmt.Errorf("missing required field 'type'")
	}
	if len(enumDef.Values) == 0 {
		return fmt.Errorf("missing required field 'values' or values is empty")
	}

	// Prepare template data
	templateData := TemplateData{
		Type:        enumDef.Type,
		Package:     "enum",
		Description: enumDef.Description,
	}

	// Process values
	var keys []string
	for key := range enumDef.Values {
		keys = append(keys, key)
	}
	sort.Strings(keys) // Ensure consistent ordering

	for _, key := range keys {
		description := enumDef.Values[key]
		if description == "" {
			description = key // Fallback to key if no description
		}

		templateData.Values = append(templateData.Values, EnumValue{
			Name:        key,
			ConstName:   fmt.Sprintf("%s%s", enumDef.Type, utils.ToPascalCase(key)),
			FieldName:   utils.ToPascalCase(key),
			Description: description,
		})
	}

	// Generate output file
	outputFile := fmt.Sprintf("internal/enum/%s.go", utils.ToSnakeCase(enumDef.Type))
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer file.Close()

	// Execute template
	if err := tmpl.Execute(file, templateData); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	fmt.Printf("Generated: %s\n", outputFile)
	return nil
}

func lcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
