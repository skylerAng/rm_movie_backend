package main

import (
	"movie_api/router"

	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	// Define echo object
	e := echo.New()
	router.API_V1(e)

	// Run ECHO 
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))
}