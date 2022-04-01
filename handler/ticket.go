package handler

import (
	"bytes"
	"html/template"

	"github.com/labstack/echo/v4"
)

func PrintTicket(c echo.Context) error {
	return nil
}

func ReadTemplate(data interface{}) (string, error) {
	tpl, err := template.ParseFiles("template/ticketPOS.html")
	if err != nil {
		return "", err
	}

	var bf bytes.Buffer
	if err := tpl.Execute(&bf, &data); err != nil {
		return "", err
	}

	return bf.String(), nil
}
