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
	users := []User {
		{
			Id: 1,
			Name: "TsukinoMito",
			Age: 16,
		},
		{
			Id: 2,
			Name: "YukiChihiro",
			Age: 10,
		},
	}
	return c.JSON(http.StatusOK, users)
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
