package definition

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/logger"
	"github.com/pkg/errors"
)

type DeleteTodoParam struct {
	TodoId int64 `validate:"required"`
}

type DeleteTodoResponse struct {
	Id int64 `json:"id"`
}

func CreateDeleteTodoParam(ctx echo.Context) (param *DeleteTodoParam, err error) {
	todoId := ctx.Param("todoId")
	param = new(DeleteTodoParam)
	param.TodoId, _ = strconv.ParseInt(todoId, 10, 64)
	if message, ok := Validator(param); !ok {
		err = errors.New(message)
		logger.L.Error(message)
	}
	return
}
