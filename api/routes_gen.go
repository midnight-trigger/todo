// This file was auto-generated.
// DO NOT EDIT MANUALLY!!!
package api

import (
	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/api/controller"
	"github.com/midnight-trigger/todo/third_party/jwt"
)

func RegisterRoutes(e *echo.Echo) {
	PostSigninUser(e, &controller.User{})
	PostUser(e, &controller.User{})
}
func RegisterAuthRoutes(e *echo.Group) {
	GetTodos(e, &controller.Todo{})
	PostTodo(e, &controller.Todo{})
	PutTodo(e, &controller.Todo{})
	PatchTodo(e, &controller.Todo{})
	DeleteTodo(e, &controller.Todo{})
}
func PostSigninUser(
	e *echo.Echo,
	inter *controller.User,
) {
	e.POST("api/v1/users/signin", func(c echo.Context) error {
		res := inter.PostSigninUser(c)
		return c.JSON(res.Meta.Code, res)
	})
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
func GetTodos(
	e *echo.Group,
	inter *controller.Todo,
) {
	e.GET("api/v1/todos", func(c echo.Context) error {
		claims, r := jwt.GetJWTClaims(c)
		if claims == nil {
			return c.JSON(r.Code, r)
		}
		res := inter.GetTodos(c, claims)
		return c.JSON(res.Meta.Code, res)
	})
}
func PostTodo(
	e *echo.Group,
	inter *controller.Todo,
) {
	e.POST("api/v1/todos", func(c echo.Context) error {
		claims, r := jwt.GetJWTClaims(c)
		if claims == nil {
			return c.JSON(r.Code, r)
		}
		res := inter.PostTodo(c, claims)
		return c.JSON(res.Meta.Code, res)
	})
}
func PutTodo(
	e *echo.Group,
	inter *controller.Todo,
) {
	e.PUT("api/v1/todos/:todoId", func(c echo.Context) error {
		claims, r := jwt.GetJWTClaims(c)
		if claims == nil {
			return c.JSON(r.Code, r)
		}
		res := inter.PutTodo(c, claims)
		return c.JSON(res.Meta.Code, res)
	})
}
func PatchTodo(
	e *echo.Group,
	inter *controller.Todo,
) {
	e.PATCH("api/v1/todos/:todoId", func(c echo.Context) error {
		claims, r := jwt.GetJWTClaims(c)
		if claims == nil {
			return c.JSON(r.Code, r)
		}
		res := inter.PatchTodo(c, claims)
		return c.JSON(res.Meta.Code, res)
	})
}
func DeleteTodo(
	e *echo.Group,
	inter *controller.Todo,
) {
	e.DELETE("api/v1/todos/:todoId", func(c echo.Context) error {
		claims, r := jwt.GetJWTClaims(c)
		if claims == nil {
			return c.JSON(r.Code, r)
		}
		res := inter.DeleteTodo(c, claims)
		return c.JSON(res.Meta.Code, res)
	})
}
