package ctrl

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getCategoriesList() []string {
	return []string{"Cycling", "Push Ups"}
}

func GetCategories(c echo.Context) error {
	return c.JSON(http.StatusOK, getCategoriesList())
}
