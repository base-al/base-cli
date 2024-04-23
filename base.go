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
	configureCommands(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func configureCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(newAppCommand())
	rootCmd.AddCommand(generateCommand())
}

func newAppCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "new [APPNAME]",
		Short: "Create a new application",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			appName := args[0]
			fmt.Printf("Creating new app: %s\n", appName)
			cloneAndSetup(appName)
		},
	}
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
	if err := cloneCmd.Run(); err != nil {
		fmt.Printf("Failed to create project: %s\n", stderr.String())
		return
	}
	fmt.Println("Project created successfully.")
	os.Chdir(appName)
	// Change to the project directory with safety check
	if err := os.Chdir(appName); err != nil {
		fmt.Printf("Failed to change directory: %s\n", err)
		return
	}
	fmt.Println("Changed to project directory to " + appName)
	// rename base.go to appName.go
	if err := os.Rename("base.go", appName+".go"); err != nil {
		fmt.Printf("Failed to rename base.go: %s\n", err)
		return
	}

	fmt.Println("Renamed base.go to " + appName + ".go")

	fmt.Println("Setup completed successfully.")
}

func generateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate various components",
	}
	cmd.AddCommand(moduleCommand())
	return cmd
}

func moduleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "module [MODULENAME]",
		Short: "Generate a new module",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			moduleName := args[0]
			fmt.Printf("Creating new module: %s\n", moduleName)
			generators.Module(moduleName)
		},
	}
}
