package api

import (
	"fmt"
	"log"
	"net/http"
	"shop-go-ddd/src/infraestructure/dependencies"
	"shop-go-ddd/src/infraestructure/routes"
)

func main() {
	container, err := dependencies.NewContainer()
	if err != nil {
		fmt.Printf(err.Error())
	}

	routes.InitializeRoutes(container)

	log.Fatal(http.ListenAndServe(":8080", container.RouterHandler()))
}
