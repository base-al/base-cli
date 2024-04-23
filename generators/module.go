package generators

import (
	"fmt"

	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Module creates a new module inside the app directory of the current directory.
func Module(moduleName string) {
	modulePath := filepath.Join("app", moduleName)
	if _, err := os.Stat(modulePath); !os.IsNotExist(err) {
		fmt.Printf("Directory %s already exists\n", modulePath)
		return
	}

	// Create directory for module with module name
	if err := os.MkdirAll(modulePath, 0755); err != nil {
		fmt.Printf("Failed to create module directory: %s\n", err)
		return
	}
	fmt.Println("Module created successfully.")

	// Generate files for module
	generateFiles(modulePath, moduleName)
}

// generateFiles creates the necessary files for a module.
func generateFiles(modulePath, moduleName string) {
	// Paths for files

	// Create model.go
	GenerateModelsFile(modulePath, moduleName, moduleName)

	// Use GenerateRouterFile to create router.go
	GenerateRouterFile(modulePath, moduleName, moduleName)

	// Use GenerateServiceFile to create service.go
	GenerateServiceFile(modulePath, moduleName, moduleName)

	// Use GenerateTransportFile to create transport.go
	GenerateTransportFile(modulePath, moduleName, moduleName)
}

// GenerateRouterFile generates a router file from a template for the specified module.
func GenerateRouterFile(modulePath, moduleName, entityName string) error {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd) // This will show you the current working directory

	tplPath := filepath.Join("generators", "templates", "router.go.tpl") // Adjust this path as necessary
	outFile := filepath.Join(modulePath, "router.go")

	fmt.Println("Loading template from:", tplPath) // Debug print

	// Prepare the file where the output will be written
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Failed to create output file:", err) // Debug print
		return err
	}
	defer file.Close()

	fmt.Println("Created output file at:", outFile) // Debug print

	// Parse the template file
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		fmt.Println("Error parsing template file:", err) // Debug print
		return err
	}

	fmt.Println("Template parsed successfully") // Debug print
	caser := cases.Title(language.English)
	// Define the data for the template
	data := struct {
		ModuleName  string
		EntityName  string
		EntityTitle string
	}{
		ModuleName:  moduleName,
		EntityName:  entityName,
		EntityTitle: caser.String(entityName),
	}

	// Execute the template and write to file
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err) // Debug print
		return err
	}

	fmt.Printf("Router file created successfully: %s\n", outFile) // Debug print
	return nil
}

func GenerateServiceFile(modulePath, moduleName, entityName string) error {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd) // This will show you the current working directory

	tplPath := filepath.Join("generators", "templates", "service.go.tpl") // Adjust this path as necessary
	outFile := filepath.Join(modulePath, "service.go")

	fmt.Println("Loading template from:", tplPath) // Debug print

	// Prepare the file where the output will be written
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Failed to create output file:", err) // Debug print
		return err
	}
	defer file.Close()

	fmt.Println("Created output file at:", outFile) // Debug print

	// Parse the template file
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		fmt.Println("Error parsing template file:", err) // Debug print
		return err
	}

	fmt.Println("Template parsed successfully") // Debug print
	caser := cases.Title(language.English)
	// Define the data for the template
	data := struct {
		ModuleName  string
		EntityName  string
		EntityTitle string
	}{
		ModuleName:  moduleName,
		EntityName:  entityName,
		EntityTitle: caser.String(entityName),
	}

	// Execute the template and write to file
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err) // Debug print
		return err
	}

	fmt.Printf("Service file created successfully: %s\n", outFile) // Debug print
	return nil
}

func GenerateTransportFile(modulePath, moduleName, entityName string) error {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd) // This will show you the current working directory

	tplPath := filepath.Join("generators", "templates", "transport.go.tpl") // Adjust this path as necessary
	outFile := filepath.Join(modulePath, "transport.go")

	fmt.Println("Loading template from:", tplPath) // Debug print

	// Prepare the file where the output will be written
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Failed to create output file:", err) // Debug print
		return err
	}
	defer file.Close()

	fmt.Println("Created output file at:", outFile) // Debug print

	// Parse the template file
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		fmt.Println("Error parsing template file:", err) // Debug print
		return err
	}

	fmt.Println("Template parsed successfully") // Debug print
	caser := cases.Title(language.English)
	// Define the data for the template
	data := struct {
		ModuleName  string
		EntityName  string
		EntityTitle string
	}{
		ModuleName:  moduleName,
		EntityName:  entityName,
		EntityTitle: caser.String(entityName),
	}

	// Execute the template and write to file
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err) // Debug print
		return err
	}

	fmt.Printf("Transport file created successfully: %s\n", outFile) // Debug print
	return nil
}

// GenerateModelsFile creates a models file for the specified module.
func GenerateModelsFile(modulePath, moduleName, entityName string) error {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd) // This will show you the current working directory

	tplPath := filepath.Join("generators", "templates", "models.go.tpl") // Adjust this path as necessary
	outFile := filepath.Join(modulePath, "models.go")

	fmt.Println("Loading template from:", tplPath) // Debug print

	// Prepare the file where the output will be written
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Failed to create output file:", err) // Debug print
		return err
	}
	defer file.Close()

	fmt.Println("Created output file at:", outFile) // Debug print
	caser := cases.Title(language.English)
	// Parse the template file
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		fmt.Println("Error parsing template file:", err) // Debug print
		return err
	}

	fmt.Println("Template parsed successfully") // Debug print

	// Define the data for the template
	data := struct {
		ModuleName  string
		EntityName  string
		EntityTitle string
	}{
		ModuleName:  moduleName,
		EntityName:  entityName,
		EntityTitle: caser.String(entityName),
	}

	// Execute the template and write to file
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err) // Debug print
		return err
	}

	fmt.Printf("Models file created successfully: %s\n", outFile) // Debug print
	return nil
}
