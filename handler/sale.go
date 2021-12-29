package handler

import (
	"fmt"
	"net/http"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/entity"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/model"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/util"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/util/response"
	"github.com/Leonardo-Antonio/validmor"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type sale struct {
	store model.Isale
}

func NewSale(_store model.Isale) *sale {
	return &sale{_store}
}

func (s *sale) Create(c echo.Context) error {
	body := new(entity.Sale)
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.Err(err.Error(), nil))
	}

	errs := validmor.ValidateStruct(*body)
	if len(errs) != 0 {
		return c.JSON(http.StatusBadRequest, response.Err("los campos son obligatorios", util.ErrToString(errs)))
	}

	if len(body.Products) == 0 {
		return c.JSON(http.StatusBadRequest, response.Err("debe tener por lo menos un producto para realizar una venta", nil))
	}

	result, err := s.store.Create(body)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.JSON(http.StatusBadRequest, response.Err("los datos que ingreso ya están en uso", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, response.Success(fmt.Sprintf("la orden de %s, se creo con correctamente", body.Nick), result))
}

func (s *sale) Update(c echo.Context) error {
	body := new(entity.Sale)
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.Err(err.Error(), nil))
	}

	if body.Id.IsZero() {
		return c.JSON(http.StatusBadRequest, response.Err("el id es obligatorio", nil))
	}

	errs := validmor.ValidateStruct(*body)
	if len(errs) != 0 {
		return c.JSON(http.StatusBadRequest, response.Err("los campos son obligatorios", util.ErrToString(errs)))
	}

	if len(body.Products) == 0 {
		return c.JSON(http.StatusBadRequest, response.Err("debe tener por lo menos una categoria", nil))
	}

	result, err := s.store.Update(body)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.JSON(http.StatusBadRequest, response.Err("los datos que ingreso ya están en uso", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success(fmt.Sprintf("la orden de %s, se actualizo correctamente", body.Nick), result))
}

func (s *sale) GetAll(c echo.Context) error {
	products, err := s.store.FindAll(true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success("ok", products))
}

func (s *sale) DeleteById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Err("el id ingresado no es valido", nil))
	}
	result, err := s.store.DeleteById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success("ok", result))
}
