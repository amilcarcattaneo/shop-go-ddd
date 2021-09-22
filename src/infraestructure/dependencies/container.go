package dependencies

import (
	"github.com/go-resty/resty"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Container defines a container for dependencies
type Container struct {
	db            *gorm.DB
	client        *resty.Client
	routerHandler *mux.Router
}

// NewContainer returns a container with the dependencies
func NewContainer() (*Container, error) {
	db, err := gorm.Open(
		"sqlite3",
		"challenge.db",
	)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	routerHandler := mux.NewRouter()
	client := resty.New()

	return &Container{
		db:            db,
		client:        client,
		routerHandler: routerHandler,
	}, nil
}

// Database returns database
func (container *Container) Database() *gorm.DB {
	return container.db
}

// RouterHandler returns router handler
func (container *Container) RouterHandler() *mux.Router {
	return container.routerHandler
}

// HTTPClient returns http client
func (container *Container) HTTPClient() *resty.Client {
	return container.client
}
