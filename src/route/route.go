package route

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	users = map[string]User{
		"1": User{
			Id:   1,
			Name: "TsukinoMito",
			Age:  16,
		},
		"2": User{
			Id:   2,
			Name: "YukiChihiro",
			Age:  10,
		},
	}
)

func Root(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func UserById(c echo.Context) error {
	return c.JSON(http.StatusOK, users[c.Param("id")])
}
