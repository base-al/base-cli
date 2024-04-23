# Base CLI Tool

The Base CLI Tool is a command-line interface designed to help generate content for https://github.com/base-al/base-core

## Features

- **Module Generation**: Automatically generates module structure including routers, services, transports, and models based on predefined templates.
- **Scalability**: Easily extend the CLI tool to include more commands and functionalities.
- **Embedded Templates**: Utilizes Go's `embed` package to include all necessary templates directly within the binary for standalone operation.

## Prerequisites

Before you install and start using the CLI tool, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or higher)

## Installation

To install the Base CLI Tool, you can use our installation script which will download the binary and set it up on your system:

```bash
curl -sSL https://raw.githubusercontent.com/base-al/base-cli/main/install.sh | sh

```

## Usage

The Base CLI is designed to streamline operations and management tasks within your project. Below are detailed instructions on how to use the various commands available.

### Generate new project

```
base new [projectname]
```

cd to [projectname] and start use base.

### Generating Modules

To generate a new module with all necessary components including models, routers, services, and transport logic, use the following command:

```bash
base generate module <module_name>
```


To write a comprehensive Usage section for the README of your CLI tool, given the directory structure and components you've described, we'll need to cover how to utilize the tool effectively. This involves detailing the functionality provided by your CLI, particularly focusing on the module generation features and other utilities embedded within your application.

Here's a breakdown of how to use the CLI tool effectively based on the directory structure and scripts you've listed:

Usage Section for README.md
markdown
Copy code
## Usage

The Base CLI is designed to streamline operations and management tasks within your project. Below are detailed instructions on how to use the various commands available.

### Generating Modules

To generate a new module with all necessary components including models, routers, services, and transport logic, use the following command:

```bash
base generate module <module_name>
```
This command will create a new directory under app/ with the specified module name and populate it with boilerplate files based on the templates located in generators/templates/.

Example:
To generate a module named user, you would run:

```bash
base generate module user
```
This will create the following structure:

```
app/
└── user/
    ├── models.go
    ├── router.go
    ├── service.go
    └── transport.go
```
### Database Seeding
If you need to populate your database with initial data, you can utilize the seeding script provided:

```
base seed
```
This script uses predefined data structures and inserts them into your database, which is useful for development and testing environments.

### Installation
Ensure you have installed the CLI tool by following the instructions provided in the Installation section.

### Additional Commands
You can extend the CLI tool with additional commands by modifying the base.go file. Each new command should be registered within this file to ensure it is recognized by the CLI.

### Advanced Configuration
For advanced users, you can modify the templates used for module generation by editing the .tpl files located in generators/templates/. This allows for customization of the boilerplate code that is generated for new modules.

### More command soon!
