package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	s.app.Use(middleware.Recover())
}

func (s *server) Router() {
	s.app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hola")
	})
}

func (s *server) Listeing() {
	if err := s.app.Start(fmt.Sprintf(":%s", s.port)); err != nil {
		log.Fatalln(err)
	}
}
