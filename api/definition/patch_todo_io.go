package definition

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/logger"
	"github.com/pkg/errors"
)

type PatchTodoParam struct {
	TodoId int64 `validate:"required"`
}

type PatchTodoRequestBody struct {
	Status string `json:"status" validate:"required,oneof=todo progress finished"`
}

type PatchTodoResponse struct {
	Id        int64  `json:"id"`
	UserId    string `json:"user_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreatePatchTodoRequestBodyAndParam(ctx echo.Context) (param *PatchTodoParam, body *PatchTodoRequestBody, err error) {
	todoId := ctx.Param("todoId")
	param = new(PatchTodoParam)
	param.TodoId, _ = strconv.ParseInt(todoId, 10, 64)
	if message, ok := Validator(param); !ok {
		err = errors.New(message)
		logger.L.Error(message)
		return
	}

	body = new(PatchTodoRequestBody)
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
