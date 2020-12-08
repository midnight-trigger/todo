package definition

import (
	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/logger"
	"github.com/pkg/errors"
)

type PostTodoRequestBody struct {
	Title string `json:"title" validate:"required,lte=255"`
	Body  string `json:"body" validate:"required"`
}

type PostTodoResponse struct {
	Id        int64  `json:"id"`
	UserId    string `json:"user_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreatePostTodoRequestBody(ctx echo.Context) (body *PostTodoRequestBody, err error) {
	body = new(PostTodoRequestBody)
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
