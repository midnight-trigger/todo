// This file was auto-generated.
// DO NOT EDIT MANUALLY!!!
package api

import (
	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/api/controller"
)

func RegisterRoutes(e *echo.Echo) {
	PostUser(e, &controller.User{})
}
func RegisterAuthRoutes(e *echo.Group) {
}
func PostUser(
	e *echo.Echo,
	inter *controller.User,
) {
	e.POST("api/v1/users", func(c echo.Context) error {
		res := inter.PostUser(c)
		return c.JSON(res.Meta.Code, res)
	})
}
