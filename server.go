package main

import (
	"github.com/labstack/echo"
	"2434api-go/src/route"
)

func main() {
	e := echo.New()

	e.GET("/", route.Root)
	e.GET("/users/:id", route.UserById)

	e.Logger.Fatal(e.Start(":1323"))
}
