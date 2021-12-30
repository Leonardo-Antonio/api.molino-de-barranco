package handler

import (
	"errors"
	"fmt"
	"log"
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

type product struct {
	store model.Iproduct
}

func NewProduct(_store model.Iproduct) *product {
	return &product{_store}
}

func (p *product) Create(c echo.Context) error {
	body := new(entity.Product)
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.Err(err.Error(), nil))
	}

	errs := validmor.ValidateStruct(*body)
	if len(errs) != 0 {
		return c.JSON(http.StatusBadRequest, response.Err("los campos son obligatorios", util.ErrToString(errs)))
	}

	if len(body.Categories) == 0 {
		return c.JSON(http.StatusBadRequest, response.Err("debe tener por lo menos una categoria", nil))
	}

	body.Amount = 1

	result, err := p.store.Create(body)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.JSON(http.StatusBadRequest, response.Err("los datos que ingreso ya están en uso", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, response.Success(fmt.Sprintf("el producto %s, se creo con correctamente", body.Name), result))
}

func (p *product) Update(c echo.Context) error {
	body := new(entity.Product)
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

	if len(body.Categories) == 0 {
		return c.JSON(http.StatusBadRequest, response.Err("debe tener por lo menos una categoria", nil))
	}

	result, err := p.store.Update(body)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.JSON(http.StatusBadRequest, response.Err("los datos que ingreso ya están en uso", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success(fmt.Sprintf("el producto %s, se actualizo correctamente", body.Name), result))
}

func (p *product) GetAll(c echo.Context) error {
	products, err := p.store.FindAll()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusNoContent, response.Err("ok", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success("ok", products))
}

func (p *product) GetByEan(c echo.Context) error {
	product, err := p.store.FindByEan(c.Param("ean"))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusBadRequest, response.Err("el ean que ingreso no existe", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success("ok", product))
}

func (p *product) DeleteById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	log.Println(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Err("el id ingresado no es valido", nil))
	}
	result, err := p.store.DeleteById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success("ok", result))
}
