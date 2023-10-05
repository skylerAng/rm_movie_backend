package router

import (
	"movie_api/handler"

	"github.com/labstack/echo/v4"
)

func API_V1(e *echo.Echo) {
	// Movie group
	group := e.Group("movie")

	group.GET("/", handler.GetBase)
	group.GET("/crawl", handler.CrawlMovies)
	group.GET("/list", handler.ListMovies)
	group.GET("/search", handler.SearchMovie)

	user_group := e.Group("user")
	user_group.GET("/", handler.GetBase)
	user_group.GET("/login", handler.Login)
}
