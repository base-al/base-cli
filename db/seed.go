package db

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Seeds all seed files in the app/[module]/seed.go  and core/[module]/seed.go

func SeedAll() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory")
		return
	}

	// Get the app directory
	appDir := filepath.Join(cwd, "app")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		fmt.Println("App directory does not exist")
		return
	}

	// Get the core directory
	coreDir := filepath.Join(cwd, "core")
	if _, err := os.Stat(coreDir); os.IsNotExist(err) {
		fmt.Println("Core directory does not exist")
		return
	}

	// Get all the modules in the app directory
	appModules, err := getModules(appDir)
	if err != nil {
		fmt.Println("Failed to get app modules")
		return
	}

	// Get all the modules in the core directory
	coreModules, err := getModules(coreDir)
	if err != nil {
		fmt.Println("Failed to get core modules")
		return
	}

	// Seed all the modules in the app directory
	for _, module := range appModules {
		seedFile := filepath.Join(appDir, module, "seed.go")
		if _, err := os.Stat(seedFile); os.IsNotExist(err) {
			fmt.Printf("Seed file for module %s does not exist\n", module)
			continue
		}
		fmt.Printf("Seeding module %s\n", module)
		err := seedModule(seedFile)
		if err != nil {
			fmt.Printf("Failed to seed module %s\n", module)
		}
	}

	// Seed all the modules in the core directory
	for _, module := range coreModules {
		seedFile := filepath.Join(coreDir, module, "seed.go")
		if _, err := os.Stat(seedFile); os.IsNotExist(err) {
			fmt.Printf("Seed file for module %s does not exist\n", module)
			continue
		}
		fmt.Printf("Seeding module %s\n", module)
		err := seedModule(seedFile)
		if err != nil {
			fmt.Printf("Failed to seed module %s\n", module)
		}
	}
}

// getModules returns all the modules in the given directory

func getModules(dir string) ([]string, error) {
	var modules []string
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			modules = append(modules, file.Name())
		}
	}
	return modules, nil
}

// seedModule seeds the given module
// seedModule seeds the given module
func seedModule(seedFile string) error {
	// Change to the directory of the seed file
	seedDir := filepath.Dir(seedFile)
	err := os.Chdir(seedDir)
	if err != nil {
		return err
	}

	// Get the package name
	packageName := getPackageName(seedFile)
	if packageName == "" {
		return fmt.Errorf("failed to get package name")
	}

	// Import the seed file
	imports := fmt.Sprintf("\t\"%s\"\n", packageName)

	// Open the seed file
	file, err := os.Open(seedFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the contents of the seed file
	var content string
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		content += string(buf[:n])
	}

	// Check if the seed function already exists
	if strings.Contains(content, "func Seed(db *gorm.DB)") {
		return fmt.Errorf("seed function already exists")
	}

	// Add the seed function
	seedFunction := "\nfunc Seed(db *gorm.DB) {\n\n}\n"
	content += seedFunction

	// Write the contents back to the seed file
	err = os.WriteFile(seedFile, []byte(content), 0644)
	if err != nil {
		return err
	}

	// Add the import to the seed file
	content = strings.Replace(content, "import (", "import (\n"+imports, 1)
	err = os.WriteFile(seedFile, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

// getPackageName returns the package name of the given file

func getPackageName(file string) string {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer f.Close()

	// Read the contents of the file
	var content string
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return ""
		}
		content += string(buf[:n])
	}

	// Get the package name
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "package ") {
			return strings.TrimPrefix(line, "package ")
		}
	}

	return ""
}
