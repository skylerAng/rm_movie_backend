package models

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type SimpleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HandleResponseSimple(c echo.Context, status int, message string) error {
	return c.JSON(http.StatusOK, &SimpleResponse{
		Status: status,
		Message: message,
	})
}