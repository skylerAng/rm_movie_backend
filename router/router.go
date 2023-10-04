package router

import (
	"movie_api/handler"

	"github.com/labstack/echo/v4"
)

func API_V1(e *echo.Echo) {
	e.GET("/", handler.GetBase)

}
