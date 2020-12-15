package definition

import (
	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/logger"
	"github.com/pkg/errors"
)

type PostSignupUserRequestBody struct {
	Username string `json:"username" validate:"required,lte=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=255"`
}

type PostSignupUserResponse struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreatePostSignupUserRequestBody(ctx echo.Context) (body *PostSignupUserRequestBody, err error) {
	body = new(PostSignupUserRequestBody)
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
