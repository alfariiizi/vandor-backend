package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	cmdutils "github.com/alfariiizi/vandor/cmd/utils"
)

type DomainData struct {
	ModuleName  string // e.g., "github.com/alfariiizi/vandor"
	Name        string // e.g., "User"
	LowerName   string // e.g., "user"
	PackageName string // e.g., "domain"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/domain/cmd-new-domain/main.go <domain-name>")
		fmt.Println("Example: go run cmd/domain/cmd-new-domain/main.go Product")
		os.Exit(1)
	}

	domainName := strings.TrimSpace(os.Args[1])
	if domainName == "" {
		fmt.Println("Domain name cannot be empty")
		os.Exit(1)
	}

	// Capitalize first letter
	domainName = strings.ToUpper(domainName[:1]) + domainName[1:]

	data := DomainData{
		ModuleName:  cmdutils.GetModuleName(),
		Name:        domainName,
		LowerName:   strings.ToLower(domainName),
		PackageName: "domain",
	}

	// Check if domain already exists
	domainPath := filepath.Join("internal", "core", "domain", "model", strings.ToLower(domainName)+".go")
	if _, err := os.Stat(domainPath); err == nil {
		fmt.Printf("Domain %s already exists at %s\n", domainName, domainPath)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to overwrite it? (y/N): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response != "y" && response != "yes" {
			fmt.Println("Operation cancelled")
			return
		}
	}

	// Create domain file from template
	if err := createDomainFile(data, domainPath); err != nil {
		fmt.Printf("Error creating domain file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Domain %s created successfully at %s\n", domainName, domainPath)
	// fmt.Println("Run 'go run cmd/generate-domains/main.go' to update domain registry")
}

func createDomainFile(data DomainData, outputPath string) error {
	// Ensure directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Parse template
	templateFile := filepath.Join("cmd", "domain", "cmd-new-domain", "domain.tmpl")
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create output file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Execute template
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
