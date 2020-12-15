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

// ログイン
func (c *User) PostSigninUser(ctx echo.Context) (response *Response) {
	defer func() {
		// panicエラー発生時のハンドリング
		if e := recover(); e != nil {
			c.ServerErrorException(errors.New(""), fmt.Sprintf("%+v", e))
			logger.L.Error(c.ErrMessage)
			c.FormatResult(&c.Result, ctx)
		}
	}()

	// リクエスト取得
	body, err := definition.CreatePostSigninUserRequestBody(ctx)
	if err != nil {
		c.ValidationException(err, err.Error())
		return c.FormatResult(&c.Result, ctx)
	}

	// ビジネスロジック&レスポンス返却
	service := domain.GetNewUserService()
	result := service.PostSigninUser(body)
	return c.FormatResult(&result, ctx)
}

// 会員登録
func (c *User) PostSignupUser(ctx echo.Context) (response *Response) {
	defer func() {
		// panicエラー発生時のハンドリング
		if e := recover(); e != nil {
			c.ServerErrorException(errors.New(""), fmt.Sprintf("%+v", e))
			logger.L.Error(c.ErrMessage)
			c.FormatResult(&c.Result, ctx)
		}
	}()

	// リクエスト取得
	body, err := definition.CreatePostSignupUserRequestBody(ctx)
	if err != nil {
		c.ValidationException(err, err.Error())
		return c.FormatResult(&c.Result, ctx)
	}

	// ビジネスロジック&レスポンス返却
	service := domain.GetNewUserService()
	result := service.PostSignupUser(body)
	return c.FormatResult(&result, ctx)
}
