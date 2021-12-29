package handler

import (
	"errors"
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

type category struct {
	store model.Icategory
}

func NewCategory(_store model.Icategory) *category {
	return &category{_store}
}

func (p *category) Create(c echo.Context) error {
	body := new(entity.Category)
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.Err(err.Error(), nil))
	}

	errs := validmor.ValidateStruct(*body)
	if len(errs) != 0 {
		return c.JSON(http.StatusBadRequest, response.Err("los campos son obligatorios", util.ErrToString(errs)))
	}

	result, err := p.store.Create(body)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.JSON(http.StatusBadRequest, response.Err("los datos que ingreso ya están en uso", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, response.Success(fmt.Sprintf("la categorias %s, se creo con correctamente", body.Name), result))
}

func (p *category) Update(c echo.Context) error {
	body := new(entity.Category)
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.Err(err.Error(), nil))
	}

	if body.Id.IsZero() {
		return c.JSON(http.StatusBadRequest, response.Err("el id es obligatorio", nil))
	}

	if len(body.Name) == 0 {
		return c.JSON(http.StatusBadRequest, response.Err("el nombre de la categoria es obligatorio", nil))
	}
	body.Ean = ""

	result, err := p.store.Update(body)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.JSON(http.StatusBadRequest, response.Err("los datos que ingreso ya están en uso", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success(fmt.Sprintf("la categoria %s, se actualizo correctamente", body.Name), result))
}

func (p *category) GetAll(c echo.Context) error {
	categories, err := p.store.FindAll()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusNoContent, response.Err("ok", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success("ok", categories))
}

func (p *category) DeleteById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Err("el id ingresado no es valido", nil))
	}
	result, err := p.store.DeleteById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success("ok", result))
}
