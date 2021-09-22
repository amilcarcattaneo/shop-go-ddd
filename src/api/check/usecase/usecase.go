package usecase

import (
	"shop-go-ddd/src/api/check/repository"
	"shop-go-ddd/src/infraestructure/dependencies"
)

// UseCases defines a usecases struct
type UseCases struct {
	checkRepository repository.CheckRepository
}

// NewUseCases returns a new usecases
func NewUseCases(container *dependencies.Container) *UseCases {
	return &UseCases{
		checkRepository: repository.NewCheckRepository(container),
	}
}
