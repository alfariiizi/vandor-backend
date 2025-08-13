package main

import (
	"fmt"
	"log"
	"os"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"gopkg.in/yaml.v3"
)

type FieldInfo struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type RelationInfo struct {
	Type   string `yaml:"type"`   // hasMany, hasOne, manyToMany
	Target string `yaml:"target"` // related table
}

type TableInfo struct {
	Name      string         `yaml:"name"`
	Fields    []FieldInfo    `yaml:"fields"`
	Relations []RelationInfo `yaml:"relations"`
}

func main() {
	schemaPath := "./database/schema"
	graph, err := entc.LoadGraph(schemaPath, &gen.Config{})
	if err != nil {
		log.Fatalf("failed to load schema graph: %v", err)
	}

	var tables []TableInfo
	for _, node := range graph.Nodes {
		ti := TableInfo{Name: node.Name}
		for _, f := range node.Fields {
			ti.Fields = append(ti.Fields, FieldInfo{Name: f.Name, Type: f.Type.String()})
		}
		for _, e := range node.Edges {
			rType := "hasMany"
			if e.M2M() {
				rType = "manyToMany"
			} else if e.Unique {
				rType = "hasOne"
			}
			ti.Relations = append(ti.Relations, RelationInfo{Type: rType, Target: e.Type.Name})
		}
		tables = append(tables, ti)
	}

	out, err := yaml.Marshal(tables)
	if err != nil {
		log.Fatalf("failed to marshal yaml: %v", err)
	}
	if err := os.WriteFile("ent_schema_summary.yaml", out, 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
	fmt.Println("✅ Generated ent_schema_summary.yaml")
}
