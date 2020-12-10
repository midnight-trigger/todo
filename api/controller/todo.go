package controller

import (
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/api/domain"
	"github.com/midnight-trigger/todo/logger"
	"github.com/midnight-trigger/todo/third_party/jwt"
)

type Todo struct {
	Base
}

func (c *Todo) GetTodos(ctx echo.Context, claims *jwt.Claims) (response *Response) {
	defer func() {
		if e := recover(); e != nil {
			c.ServerErrorException(errors.New(""), fmt.Sprintf("%+v", e))
			logger.L.Error(c.ErrMessage)
			c.FormatResult(&c.Result, ctx)
		}
	}()

	params, err := definition.CreateGetTodosParam(ctx)
	if err != nil {
		c.ValidationException(err, err.Error())
		return c.FormatResult(&c.Result, ctx)
	}

	service := domain.GetNewTodoService()
	result := service.GetTodos(params, claims.UserId)
	return c.FormatResult(&result, ctx)
}

func (c *Todo) PostTodo(ctx echo.Context, claims *jwt.Claims) (response *Response) {
	defer func() {
		if e := recover(); e != nil {
			c.ServerErrorException(errors.New(""), fmt.Sprintf("%+v", e))
			logger.L.Error(c.ErrMessage)
			c.FormatResult(&c.Result, ctx)
		}
	}()

	body, err := definition.CreatePostTodoRequestBody(ctx)
	if err != nil {
		c.ValidationException(err, err.Error())
		return c.FormatResult(&c.Result, ctx)
	}

	service := domain.GetNewTodoService()
	result := service.PostTodo(body, claims.UserId)
	return c.FormatResult(&result, ctx)
}

func (c *Todo) PutTodo(ctx echo.Context, claims *jwt.Claims) (response *Response) {
	defer func() {
		if e := recover(); e != nil {
			c.ServerErrorException(errors.New(""), fmt.Sprintf("%+v", e))
			logger.L.Error(c.ErrMessage)
			c.FormatResult(&c.Result, ctx)
		}
	}()

	param, body, err := definition.CreatePutTodoRequestBodyAndParam(ctx)
	if err != nil {
		c.ValidationException(err, err.Error())
		return c.FormatResult(&c.Result, ctx)
	}

	service := domain.GetNewTodoService()
	result := service.PutTodo(param, body, claims.UserId)
	return c.FormatResult(&result, ctx)
}

func (c *Todo) PatchTodo(ctx echo.Context, claims *jwt.Claims) (response *Response) {
	defer func() {
		if e := recover(); e != nil {
			c.ServerErrorException(errors.New(""), fmt.Sprintf("%+v", e))
			logger.L.Error(c.ErrMessage)
			c.FormatResult(&c.Result, ctx)
		}
	}()

	param, body, err := definition.CreatePatchTodoRequestBodyAndParam(ctx)
	if err != nil {
		c.ValidationException(err, err.Error())
		return c.FormatResult(&c.Result, ctx)
	}

	service := domain.GetNewTodoService()
	result := service.PatchTodo(param, body, claims.UserId)
	return c.FormatResult(&result, ctx)
}

func (c *Todo) DeleteTodo(ctx echo.Context, claims *jwt.Claims) (response *Response) {
	defer func() {
		if e := recover(); e != nil {
			c.ServerErrorException(errors.New(""), fmt.Sprintf("%+v", e))
			logger.L.Error(c.ErrMessage)
			c.FormatResult(&c.Result, ctx)
		}
	}()

	param, err := definition.CreateDeleteTodoParam(ctx)
	if err != nil {
		c.ValidationException(err, err.Error())
		return c.FormatResult(&c.Result, ctx)
	}

	service := domain.GetNewTodoService()
	result := service.DeleteTodo(param, claims.UserId)
	return c.FormatResult(&result, ctx)
}
