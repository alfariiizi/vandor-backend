package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/alfariiizi/vandor/cmd/utils"
	"github.com/alfariiizi/vandor/internal/config"
	"github.com/hedwigz/entviz"
	"github.com/lrstanley/entrest"
	"github.com/ogen-go/ogen"
)

func main() {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	schemaPath := filepath.Join(projectRoot, "database", "schema")
	outputPath := filepath.Join(projectRoot, "internal", "infrastructure", "db")

	module := utils.GetModuleName()

	// Run Ent codegen with entrest extension
	err = entc.Generate(
		schemaPath,
		&gen.Config{
			Target:  outputPath,
			Package: fmt.Sprintf("%s/internal/infrastructure/db", module),
			Schema:  fmt.Sprintf("%s/database/schema", module),
			Features: []gen.Feature{
				gen.FeatureUpsert,
			},
		},
		entc.Extensions(
			getEntrestExt(module),
			getVizGraphExt(),
		),
	)
	if err != nil {
		log.Fatalf("ent codegen failed: %v", err)
	}

	// Patch imports in generated entrest files
	if err := patchEntrestImports(projectRoot, module); err != nil {
		log.Fatalf("failed to patch entrest imports: %v", err)
	}

	fmt.Println("✅ Ent + Entrest generation completed successfully")
}

func getEntrestExt(module string) *entrest.Extension {
	cfg := config.GetConfig()
	ex, err := entrest.NewExtension(&entrest.Config{
		Handler:      entrest.HandlerStdlib,
		WithTesting:  true,
		StrictMutate: true,
		Spec: &ogen.Spec{
			Info: ogen.Info{
				Title:   fmt.Sprintf("Admin API for %s", cfg.App.Name),
				Version: fmt.Sprintf("v%s", cfg.App.Version),
			},
			Components: &ogen.Components{
				SecuritySchemes: map[string]*ogen.SecurityScheme{
					"bearerAuth": {
						Type:         "http",
						Scheme:       "bearer",
						BearerFormat: "JWT",
						Description:  "JWT Bearer token in Authorization header",
					},
				},
			},
			Security: []ogen.SecurityRequirement{
				{
					"bearerAuth": {},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("creating entrest extension: %v", err)
	}
	return ex
}

func patchEntrestImports(projectRoot, module string) error {
	restPath := filepath.Join(projectRoot, "internal", "infrastructure", "db", "rest")
	dbImport := fmt.Sprintf(`ent "%s/internal/infrastructure/db"`, module)

	err := filepath.Walk(restPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		updated := src

		// Replace old ent import with db import
		updated = bytes.ReplaceAll(
			updated,
			[]byte(fmt.Sprintf(`"%s/ent"`, module)),
			[]byte(dbImport),
		)
		updated = bytes.ReplaceAll(
			updated,
			[]byte(fmt.Sprintf(`ent "%s/internal/infrastructure/db"`, module)),
			[]byte(dbImport),
		)

		// Replace ent. with db.
		// updated = bytes.ReplaceAll(updated, []byte("ent."), []byte("db."))

		// If file does not contain "db." usage, remove db import entirely
		if !bytes.Contains(updated, []byte("ent.")) {
			reImport := regexp.MustCompile(`(?m)^\s*ent\s+"[^"]+"\s*\n`)
			updated = reImport.ReplaceAll(updated, nil)
		} else {
			// Ensure db import exists
			if !bytes.Contains(updated, []byte(dbImport)) {
				lines := strings.Split(string(updated), "\n")
				for i, line := range lines {
					if strings.HasPrefix(line, "package ") {
						// insert import block right after the package line
						lines = append(lines[:i+1],
							append([]string{fmt.Sprintf("import (\n\t%s\n)", dbImport)}, lines[i+1:]...)...)
						break
					}
				}
				updated = []byte(strings.Join(lines, "\n"))
			}
		}

		// Format Go code
		formatted, err := format.Source(updated)
		if err != nil {
			log.Printf("warning: gofmt failed on %s: %v", path, err)
			formatted = updated
		}

		if err := os.WriteFile(path, formatted, 0644); err != nil {
			return err
		}

		// Log success
		// log.Printf("patched imports in %s", path)

		return nil
	})
	return err
}

func getVizGraphExt() entviz.Extension {
	return entviz.Extension{}
}
