package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBase(c echo.Context) error {
	return c.String(http.StatusOK, "Failed Context")
}
