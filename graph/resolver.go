package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver : checks schema and model consistancy
type Resolver struct {
	DB *gorm.DB
}
