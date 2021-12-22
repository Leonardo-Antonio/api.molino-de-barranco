package app

import (
	"log"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/util/env"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Start() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	env.Load()

	server := newServer(echo.New(), env.PORT)
	server.Middlewares()
	server.Router()
	server.Listeing()
}
