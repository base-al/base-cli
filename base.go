package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/base-al/base-cli/generators"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "base"}
	rootCmd.AddCommand(newCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var newCmd = &cobra.Command{
	Use:   "new [APPNAME]",
	Short: "Create a new application",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		fmt.Printf("Creating new app: %s\n", appName)
		// Here you can add the code to clone the repo and set up the environment
		cloneAndSetup(appName)
	},
}

func cloneAndSetup(appName string) {

	if _, err := os.Stat(appName); !os.IsNotExist(err) {
		fmt.Printf("Directory %s already exists\n", appName)
		return
	}

	// Clone the repository
	var stdout, stderr bytes.Buffer
	cloneCmd := exec.Command("git", "clone", "https://github.com/base-al/base-core.git", appName)
	cloneCmd.Stdout = &stdout
	cloneCmd.Stderr = &stderr
	err := cloneCmd.Run()

	if err != nil {
		fmt.Printf("Failed to create project: %s\n", stderr.String())
		return
	}
	fmt.Println("Project created successfully.")

	// Change to the project directory
	os.Chdir(appName)

	fmt.Println("Setup completed successfully.")
}

// base generate module [MODULENAME] or base g module [MODULENAME]
// module.go is in the generators directory

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate various components",
	// Other properties and potential subcommands can be added here
}

var generateModuleCmd = &cobra.Command{
	Use:   "module [MODULENAME]",
	Short: "Generate a new module",
	Args:  cobra.ExactArgs(1), // Ensures exactly one argument is passed
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]
		fmt.Printf("Creating new module: %s\n", moduleName)
		generators.Module(moduleName) // Call the function from your module.go
	},
}
