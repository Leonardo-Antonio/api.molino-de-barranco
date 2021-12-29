package app

import (
	"fmt"
	"log"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/model"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	app  *echo.Echo
	port string
}

func newServer(_instance *echo.Echo, _port string) *server {
	return &server{_instance, _port}
}

func (s *server) Middlewares() {
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.CORS())
	s.app.Use(middleware.Recover())
}

func (s *server) Router(_db *mongo.Database) {
	router.NewProduct(s.app, model.NewProduct(_db))
	router.NewCategory(s.app, model.NewCategory(_db))
	router.NewSale(s.app, model.NewSale(_db))
}

func (s *server) Listeing() {
	if err := s.app.Start(fmt.Sprintf(":%s", s.port)); err != nil {
		log.Fatalln(err)
	}
}
