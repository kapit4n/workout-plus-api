package ctrl

import (
	"net/http"
	"strconv"
	m "workoutplus/models"

	"github.com/labstack/echo/v4"
)

func getUsers() []m.User {
	u1 := m.User{ID: 1, FullName: "Luis Arce", Email: "luis.arce22@gmail.com", Picture: "", Username: "luarce", Password: "luarce"}
	u2 := m.User{ID: 2, FullName: "Camila Arce", Email: "camila.arce22@gmail.com", Picture: "", Username: "luarce", Password: "luarce"}
	u3 := m.User{ID: 3, FullName: "Daniel Arce", Email: "daniel.arce22@gmail.com", Picture: "", Username: "luarce", Password: "luarce"}
	users := []m.User{u1, u2, u3}
	return users
}

func UpdateUser(c echo.Context) error {
	u := &m.User{ID: 1, FullName: "Luis Arce"}
	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Deleted User")
}

func GetUser(c echo.Context) error {
	id, error := strconv.Atoi(c.Param("id"))
	users := getUsers()
	for _, u := range users {
		if u.ID == id {
			return c.JSON(http.StatusOK, u)
		}
	}
	if error != nil {
		return error
	}

	return c.String(http.StatusOK, "Not found user")
}

func ListUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, getUsers())
}

func CreateUser(c echo.Context) error {
	fullName := c.FormValue("fullName")
	email := c.FormValue("email")
	picture := c.FormValue("picture")
	username := c.FormValue("username")
	password := c.FormValue("password")
	u := &m.User{FullName: fullName, Email: email, Picture: picture, Username: username, Password: password}
	return c.JSON(http.StatusOK, u)
}
