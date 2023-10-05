package handler

import (
	"net/http"
	"movie_api/models"	

	"github.com/labstack/echo/v4"
)

func GetBase(c echo.Context) error {
	return models.HandleResponseSimple(c, http.StatusOK, "Movie API Backend")
}
