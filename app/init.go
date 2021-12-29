package app

import (
	"log"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/database"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/util/env"
	"github.com/Leonardo-Antonio/validmor"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Start() {
	validmor.Errors(validmor.ERR_ES)
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	env.Load()

	server := newServer(echo.New(), env.PORT)
	server.Middlewares()
	server.Router(database.Connect())
	server.Listeing()
}
