#!/bin/bash

# This script initializes a Go project with Gin and GORM (for PostgreSQL).
# It requires a project name as an argument.

set -e # Exit immediately if a command exits with a non-zero status.

# Check if a project name was provided
if [ -z "$1" ]; then
  echo "Usage: $0 <project-name>"
  exit 1
fi

set -x # Print commands and their arguments as they are executed.

PROJECT_NAME=$1

# Removing files of main module if exists
if [ -f go.mod ]; then
  rm -rf go.mod
fi

if [ -f go.sum ]; then
  rm -rf go.sum
fi

# Initialize go module
go mod init "$PROJECT_NAME"

# Cleaning
go clean -modcache
rm -rf $(go env GOPATH)/pkg/mod
rm -rf $(go env GOPATH)/pkg/sumdb
go clean -cache

# Get necessary dependencies
go get -u github.com/gin-gonic/gin

go get github.com/golang-jwt/jwt

go get github.com/jackc/pgx/v5

go get gorm.io/gorm

go get github.com/jackc/pgx/pgtype

go get github.com/go-playground/validator/v10

go get gorm.io/driver/postgres

go get github.com/google/uuid

go get github.com/joho/godotenv

go get github.com/brianvoe/gofakeit

go get github.com/go-gorm/datatypes

go get gorm.io/datatypes

# Handle message of success
echo "Go project '$PROJECT_NAME' initialized with Gin and GORM (PostgreSQL) dependencies."
echo "You can now start building your application."