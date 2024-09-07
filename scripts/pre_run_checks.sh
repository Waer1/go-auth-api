#!/bin/bash

# Stop execution if any command fails
set -e

echo "Running Go Vet..."
go vet ./...

echo "Running Linters..."
# Example: Using golangci-lint
golangci-lint run

# Any additional checks can be added here
