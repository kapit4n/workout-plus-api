package main

import (
	"code/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getUsers() []models.User {
	u1 := models.User{FullName: "Luis Arce", Email: "luis.arce22@gmail.com", Picture: "", Username: "luarce", Password: "luarce"}
	users := []models.User{u1}
	return users
}

func updateUser(c echo.Context) error {
	u := &models.User{ID: 1, FullName: "Luis Arce"}
	return c.JSON(http.StatusOK, u)
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Deleted User")
}

func getUser(c echo.Context) error {
	users := getUsers()
	return c.JSON(http.StatusOK, users[0])
}

func listUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, getUsers())
}

func createUser(c echo.Context) error {
	fullName := c.FormValue("fullName")
	email := c.FormValue("email")
	picture := c.FormValue("picture")
	username := c.FormValue("username")
	password := c.FormValue("password")
	u := &models.User{FullName: fullName, Email: email, Picture: picture, Username: username, Password: password}
	return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", createUser)
	e.GET("/users", listUsers)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
