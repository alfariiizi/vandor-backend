package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/alfariiizi/vandor/cmd/utils"
)

type JobTemplateData struct {
	ModuleName string
	StructName string
	VarName    string
	TaskName   string
	JobKey     string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vandor job:new <name>")
		return
	}

	jobName := os.Args[1] // e.g. sendEmail
	structName := toStructName(jobName)
	varName := toVarName(jobName)
	taskName := toTaskName(jobName)

	data := JobTemplateData{
		ModuleName: utils.GetModuleName(),
		StructName: structName,
		VarName:    varName,
		TaskName:   taskName,
		JobKey:     fmt.Sprintf("job:%s", strings.ToLower(taskName)),
	}

	tmplPath := "cmd/job/cmd-new-job/templates/job.tmpl"
	tmplBytes, _ := os.ReadFile(tmplPath)

	tmpl := template.Must(template.New("job").Parse(string(tmplBytes)))

	outputDir := "internal/core/job"
	os.MkdirAll(outputDir, os.ModePerm)

	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_job.go", strings.ToLower(jobName)))

	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	os.WriteFile(outputFile, buf.Bytes(), 0644)

	fmt.Println("Created job:", outputFile)
}

func toStructName(name string) string {
	return strings.Title(name)
}

func toVarName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func toTaskName(name string) string {
	// sendEmail -> send_email
	var result []rune
	for i, r := range name {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_', r+32)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
