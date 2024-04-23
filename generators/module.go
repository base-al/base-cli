package generators

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed templates/*
var templatesFS embed.FS

// Module creates a new module inside the app directory of the current directory.
func Module(moduleName string) {
	modulePath := filepath.Join("app", moduleName)
	if _, err := os.Stat(modulePath); !os.IsNotExist(err) {
		fmt.Printf("Directory %s already exists\n", modulePath)
		return
	}

	if err := os.MkdirAll(modulePath, 0755); err != nil {
		fmt.Printf("Failed to create module directory: %s\n", err)
		return
	}
	fmt.Println("Module created successfully.")

	entities := []string{"router", "service", "transport", "models"}
	for _, entity := range entities {
		if err := generateFileFromTemplate(modulePath, entity, moduleName); err != nil {
			fmt.Printf("Failed to generate %s file: %s\n", entity, err)
			return
		}
	}
}

// generateFileFromTemplate handles the creation of files from templates for a given entity type.
func generateFileFromTemplate(modulePath, entityType string, moduleName string) error {
	fileName := entityType + ".go"
	outFile := filepath.Join(modulePath, fileName)
	tplFile := "templates/" + entityType + ".go.tpl" // Corrected path for template file

	// Load and parse the template file
	tplData, err := templatesFS.ReadFile(tplFile)
	if err != nil {
		return fmt.Errorf("failed to read template file: %s, error: %v", tplFile, err)
	}

	tmpl, err := template.New(fileName).Parse(string(tplData))
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	file, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %s, error: %v", outFile, err)
	}
	defer file.Close()

	caser := cases.Title(language.English)
	data := struct {
		ModuleName  string
		EntityName  string
		EntityTitle string
	}{
		ModuleName:  moduleName,
		EntityName:  moduleName,
		EntityTitle: caser.String(moduleName),
	}

	// Execute the template and write to file
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template on file %s: %v", outFile, err)
	}

	fmt.Printf("%s file created successfully: %s\n", entityType, outFile)
	return nil
}
