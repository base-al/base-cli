package generators

import (
	"fmt"
	"os"
)

//   creates a new module inside app directory  of current directory

func Module(moduleName string) {
	if _, err := os.Stat(moduleName); !os.IsNotExist(err) {
		fmt.Printf("Directory %s already exists\n", moduleName)
		return
	}
	os.Chdir("app")
	// create directory for module with module name
	err := os.Mkdir(moduleName, 0755)
	if err != nil {
		fmt.Println("Failed to create module directory")
		return
	}
	fmt.Println("Module created successfully.")

	// Change to the module directory
	os.Chdir(moduleName)

	// Create files for module
	// model.go
	// service.go
	// router.go
	// schema.go
	// transport.go
	// helper.go

	// Create model.go
	modelFile, err := os.Create("model.go")
	if err != nil {
		fmt.Println("Failed to create model.go")
		return
	}
	defer modelFile.Close()
	modelFile.WriteString("package " + moduleName + "\n\n")
	modelFile.WriteString("type " + moduleName + " struct {\n\n}\n")

	// Create service.go
	serviceFile, err := os.Create("service.go")
	if err != nil {
		fmt.Println("Failed to create service.go")
		return
	}
	defer serviceFile.Close()
	serviceFile.WriteString("package " + moduleName + "\n\n")
	serviceFile.WriteString("type Service struct {\n\n}\n")

	// Create router.go
	routerFile, err := os.Create("router.go")
	if err != nil {
		fmt.Println("Failed to create router.go")
		return
	}
	defer routerFile.Close()
	routerFile.WriteString("package " + moduleName + "\n\n")
	routerFile.WriteString("import (\n\t\"github.com/gorilla/mux\"\n)\n\n")
	routerFile.WriteString("func RegisterRoutes(router *mux.Router) {\n\n}\n")

	// Create schema.go
	schemaFile, err := os.Create("schema.go")
	if err != nil {
		fmt.Println("Failed to create schema.go")
		return
	}
	defer schemaFile.Close()
	schemaFile.WriteString("package " + moduleName + "\n\n")
	schemaFile.WriteString("type " + moduleName + "Schema struct {\n\n}\n")

	// Create transport.go
	transportFile, err := os.Create("transport.go")
	if err != nil {
		fmt.Println("Failed to create transport.go")
		return
	}
	defer transportFile.Close()
	transportFile.WriteString("package " + moduleName + "\n\n")
	transportFile.WriteString("type Transport struct {\n\n}\n")

	// Create helper.go
	helperFile, err := os.Create("helper.go")
	if err != nil {
		fmt.Println("Failed to create helper.go")
		return
	}
	defer helperFile.Close()
	helperFile.WriteString("package " + moduleName + "\n\n")

	// Change to the project directory
	os.Chdir("..")
}
