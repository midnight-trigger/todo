package definition

import (
	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/logger"
	"github.com/pkg/errors"
)

type PostUserRequestBody struct {
	Username string `json:"username" validate:"required,lte=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=255"`
}

type PostUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func CreatePostUserRequestBody(ctx echo.Context) (body *PostUserRequestBody, err error) {
	body = new(PostUserRequestBody)
	if err = ctx.Bind(body); err != nil {
		logger.L.Error(err)
		return
	}
	if message, ok := Validator(body); !ok {
		err = errors.New(message)
		logger.L.Error(message)
	}
	return
}
