package ctrl

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getIntervals() []int {
	return []int{10, 20, 30, 50, 100}
}

func GetIntervals(c echo.Context) error {
	return c.JSON(http.StatusOK, getIntervals())
}
