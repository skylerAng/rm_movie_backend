package main

import (
	"movie_api/router"

	"movie_api/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.API_V1(e)
	handler.Startup()
	e.Logger.Fatal(e.Start(":1323"))
}

// func doAdd(num1 int, num2 int) (int, string) {
// 	return num1 + num2, "ladad"
// }
