package definition

import (
	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/logger"
	"github.com/pkg/errors"
)

type PostSigninUserRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=255"`
}

type PostSigninUserResponse struct {
	IdToken string `json:"id_token"`
}

func CreatePostSigninUserRequestBody(ctx echo.Context) (body *PostSigninUserRequestBody, err error) {
	body = new(PostSigninUserRequestBody)
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
