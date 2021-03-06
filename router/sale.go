package router

import (
	"fmt"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/handler"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/model"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/util/enum"
	"github.com/labstack/echo/v4"
)

func NewSale(_app *echo.Echo, _store model.Isale) {
	controller := handler.NewSale(_store)
	group := _app.Group(fmt.Sprintf("%s/%s", enum.API_BASE, "sales"))
	group.POST("", controller.Create)
	group.PUT("", controller.Update)
	group.GET("", controller.GetAll)
	group.DELETE("/:id", controller.DeleteById)
	group.GET("/search/:id", controller.GetById)
	group.GET("/search/ticket/:id", controller.GetByIdTicket)
	group.GET("/date/register/money/:date", controller.GetAllByDate)
}
