package controller

import (
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/api/domain"
	"github.com/midnight-trigger/todo/logger"
)

type User struct {
	Base
}

func (c *User) PostSigninUser(ctx echo.Context) (response *Response) {
	defer func() {
		if e := recover(); e != nil {
			c.ServerErrorException(errors.New(""), fmt.Sprintf("%+v", e))
			logger.L.Error(c.ErrMessage)
			c.FormatResult(&c.Result, ctx)
		}
	}()

	body, err := definition.CreatePostSigninUserRequestBody(ctx)
	if err != nil {
		c.ValidationException(err, err.Error())
		return c.FormatResult(&c.Result, ctx)
	}

	service := domain.GetNewUserService()
	result := service.PostSigninUser(body)
	return c.FormatResult(&result, ctx)
}

func (c *User) PostUser(ctx echo.Context) (response *Response) {
	defer func() {
		if e := recover(); e != nil {
			c.ServerErrorException(errors.New(""), fmt.Sprintf("%+v", e))
			logger.L.Error(c.ErrMessage)
			c.FormatResult(&c.Result, ctx)
		}
	}()

	body, err := definition.CreatePostUserRequestBody(ctx)
	if err != nil {
		c.ValidationException(err, err.Error())
		return c.FormatResult(&c.Result, ctx)
	}

	service := domain.GetNewUserService()
	result := service.PostUser(body)
	return c.FormatResult(&result, ctx)
}
