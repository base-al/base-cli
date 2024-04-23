#!/bin/bash

# Step 1: Check if Go is installed
if ! [ -x "$(command -v go)" ]; then
  echo "Error: Go is not installed." >&2
  echo "Please install Go and rerun this script." >&2
  exit 1
fi

# Step 2: Define the repository URL and installation directory
REPO_URL="https://github.com/base-al/base-cli.git"
INSTALL_DIR="$HOME/.base"
BIN_PATH="/usr/local/bin"

# Step 3: Clone the repository
echo "Cloning the repository..."
git clone $REPO_URL $INSTALL_DIR || { echo "Failed to clone repository."; exit 1; }

# Step 4: Build the tool
echo "Building the tool..."
cd $INSTALL_DIR
go build -o base || { echo "Failed to build the tool."; exit 1; }

# Step 5: Install the binary
echo "Installing the tool..."
sudo mv base $BIN_PATH/base || { echo "Failed to install the tool."; exit 1; }

echo "Installation completed successfully."
echo "You can now use 'base' from anywhere in your terminal."
