package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Leonardo-Antonio/api-molino-de-barranco/entity"
	"github.com/Leonardo-Antonio/api-molino-de-barranco/util/response"
	"github.com/labstack/echo/v4"
)

func GetInfoByRuc(c echo.Context) error {
	ruc := c.Param("ruc")
	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.apis.net.pe/v1/ruc?numero=%s", ruc),
		bytes.NewBuffer(nil),
	)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Err(err.Error(), nil))
	}

	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, response.Err(err.Error(), nil))
		}

		if resp.StatusCode == http.StatusUnprocessableEntity {
			return c.JSON(http.StatusBadRequest, response.Err("El RUC debe tener 11 digitos", nil))
		}

		return c.JSON(http.StatusInternalServerError, response.Err("error "+resp.Status, nil))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}
	resp.Body.Close()

	var infoRUC entity.RUC
	if err := json.Unmarshal(body, &infoRUC); err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Success("ok", infoRUC))
}
