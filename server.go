package main

import (
	"github.com/labstack/echo"
	"2434api-go/src/route"
)

func main() {
	e := echo.New()
	api := e.Group("/api/v1")

	api.GET("/", route.Root)
	api.GET("/users/:id", route.UserById)

	e.Logger.Fatal(e.Start(":1323"))
}
