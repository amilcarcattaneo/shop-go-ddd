package database

import (
	"github.com/jinzhu/gorm"

	"shop-go-ddd/src/infraestructure/dependencies"
)

// CheckDatabaseRepository defines a repository struct
type CheckDatabaseRepository struct {
	database *gorm.DB
}

// NewRepository returns a new Check repository
func NewRepository(container *dependencies.Container) *CheckDatabaseRepository {
	return &CheckDatabaseRepository{
		database: container.Database(),
	}
}
