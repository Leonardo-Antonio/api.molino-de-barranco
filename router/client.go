package router

import (
	"fmt"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/handler"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/util/enum"
	"github.com/labstack/echo/v4"
)

func NewClient(_app *echo.Echo) {
	group := _app.Group(fmt.Sprintf("%s/%s", enum.API_BASE, "ruc"))
	group.GET("/:ruc", handler.GetInfoByRuc)
}
