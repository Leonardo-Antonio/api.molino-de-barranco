package router

import (
	"fmt"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/handler"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/model"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/util/enum"
	"github.com/labstack/echo/v4"
)

func NewCategory(_app *echo.Echo, _store model.Icategory) {
	controller := handler.NewCategory(_store)
	group := _app.Group(fmt.Sprintf("%s/%s", enum.API_BASE, "categories"))
	group.POST("", controller.Create)
	group.PUT("", controller.Update)
	group.GET("", controller.GetAll)
	group.DELETE("/:id", controller.DeleteById)
}
