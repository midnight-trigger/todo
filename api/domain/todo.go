package domain

import (
	"github.com/midnight-trigger/todo/infra/mysql"
)

type Todo struct {
	Base
	MTodos mysql.ITodos
}

func GetNewTodoService() *Todo {
	todo := new(User)
	todo.MTodos = mysql.GetNewTodo()
	return todo
}
