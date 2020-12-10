package definition

import (
	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type GetTodosParam struct {
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
	Sort   string `query:"sort" validate:"oneof=DESC ASC"`
	Title  string `query:"title" validate:"lte=255"`
	Body   string `query:"body"`
	Status string `query:"status" validate:"oneof='' todo progress finished"`
}

type GetTodosResponse struct {
	Id        int64  `json:"id"`
	UserId    string `json:"user_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreateGetTodosParam(ctx echo.Context) (params *GetTodosParam, err error) {
	params = new(GetTodosParam)
	if err = ctx.Bind(params); err != nil {
		logger.L.Error(err)
		return
	}
	if params.Limit == 0 {
		params.Limit = viper.GetInt("todos.listLimit")
	}
	if params.Sort == "" {
		params.Sort = "DESC"
	}
	if message, ok := Validator(params); !ok {
		err = errors.New(message)
		logger.L.Error(message)
	}
	return
}
