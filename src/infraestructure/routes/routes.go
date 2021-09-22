package routes

import (
	"shop-go-ddd/src/api/check/delivery/http"
	"shop-go-ddd/src/infraestructure/dependencies"
)

func InitializeRoutes(container *dependencies.Container) {
	routerHandler := container.RouterHandler()

	checkHandler := http.NewCheckHandler(container)

	routerHandler.HandleFunc("/check", checkHandler.Check).Methods("GET")
}
