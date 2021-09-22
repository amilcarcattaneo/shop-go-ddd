package repository

import (
	"shop-go-ddd/src/api/check/repository/database"
	"shop-go-ddd/src/api/errors/entities"
	"shop-go-ddd/src/infraestructure/dependencies"
)

// CheckRepository defines an interface
type CheckRepository interface {
	Check() (bool, *entities.ErrorType)
}

// Repository defines a repository struct
type Repository struct {
	*database.CheckDatabaseRepository
}

// NewCheckRepository returns a new Check repository
func NewCheckRepository(container *dependencies.Container) CheckRepository {
	return &Repository{
		database.NewRepository(container),
	}
}
