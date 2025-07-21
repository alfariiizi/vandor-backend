package main

import (
	"fmt"
	"os"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage:")
	// 	fmt.Println("  enum-generator add <enum-name>    - Create a new YAML enum file")
	// 	fmt.Println("  enum-generator generate           - Generate Go enums from all YAML files")
	// 	os.Exit(1)
	// }
	//
	// command := os.Args[1]
	// switch command {
	// case "add":
	// 	if len(os.Args) < 3 {
	// 		fmt.Println("Usage: enum-generator add <enum-name>")
	// 		os.Exit(1)
	// 	}
	// 	addEnum(os.Args[2])
	// case "generate":
	// 	generateEnums()
	// default:
	// 	fmt.Printf("Unknown command: %s\n", command)
	// 	os.Exit(1)
	// }
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	addEnum(os.Args[1])
}

func addEnum(enumName string) {
	// Ensure enum directory exists
	if err := os.MkdirAll("enum", 0755); err != nil {
		fmt.Printf("Error creating enum directory: %v\n", err)
		os.Exit(1)
	}

	// Convert enum name to proper format
	typeName := toPascalCase(enumName)
	fileName := fmt.Sprintf("enum/%s.yaml", toSnakeCase(enumName))

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("Enum file %s already exists\n", fileName)
		return
	}

	// Create template YAML content
	template := fmt.Sprintf(`type: %s
description: Description for %s enum
values:
  EXAMPLE_VALUE_1: Description for example value 1
  EXAMPLE_VALUE_2: Description for example value 2
`, typeName, typeName)

	// Write file
	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		fmt.Printf("Error creating enum file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created enum file: %s\n", fileName)
	fmt.Println("Please edit the file and add your enum values, then run 'enum-generator generate'")
}
