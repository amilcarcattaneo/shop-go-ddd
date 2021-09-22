package http

import (
	"shop-go-ddd/src/api/check/usecase"
	"shop-go-ddd/src/infraestructure/dependencies"
)

type CheckHandler struct {
	usecase *usecase.UseCases
}

func NewCheckHandler(container *dependencies.Container) *CheckHandler {
	return &CheckHandler{
		usecase: usecase.NewUseCases(container),
	}
}
