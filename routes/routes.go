package routes

import (
	"crud/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello this is echo")
	})

	e.GET("/user", controllers.FetchAllUser)
	e.POST("/user", controllers.StoreUser)
	e.PUT("/user", controllers.UpdateUser)
	e.DELETE("/user", controllers.DeleteUser)

	return e
}
