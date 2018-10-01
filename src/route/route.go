package route

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func UserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, getUser(id))
}

func getUser(id int) *User {
	u := &User{
		Id:   id,
		Name: "Sample Name",
		Age:  99,
	}
	return u
}
