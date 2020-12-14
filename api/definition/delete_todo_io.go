package definition

import (
	"strconv"

	"github.com/labstack/echo"
)

type DeleteTodoParam struct {
	TodoId int64
}

type DeleteTodoResponse struct {
	Id int64 `json:"id"`
}

func CreateDeleteTodoParam(ctx echo.Context) (param *DeleteTodoParam, err error) {
	todoId := ctx.Param("todoId")
	param = new(DeleteTodoParam)
	param.TodoId, _ = strconv.ParseInt(todoId, 10, 64)
	return
}
