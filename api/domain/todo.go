package domain

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/infra/mysql"
	"github.com/midnight-trigger/todo/logger"
)

type Todo struct {
	Base
	MUsers mysql.IUsers
	MTodos mysql.ITodos
}

func GetNewTodoService() *Todo {
	todo := new(Todo)
	todo.MUsers = mysql.GetNewUser()
	todo.MTodos = mysql.GetNewTodo()
	return todo
}

func (s *Todo) PostTodo(body *definition.PostTodoRequestBody, userId string) (r Result) {
	r.New()

	// jwtの持ち主のDB存在チェック
	_, err := s.MUsers.FindById(userId)
	if gorm.IsRecordNotFoundError(err) {
		r.UserNotFoundException(errors.New(""))
		logger.L.Error(r.ErrMessage)
		return
	}
	if err != nil {
		r.ServerErrorException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	// DBインサート
	todo := new(mysql.Todos)
	s.SetStructOnSameField(body, todo)
	todo.UserId = userId

	insertedTodo, err := s.MTodos.Create(todo)
	if err != nil {
		r.ServerErrorException(err, err.Error())
		logger.L.Error(err)
		return
	}

	response := new(definition.PostTodoResponse)
	s.SetStructOnSameField(insertedTodo, response)
	r.Data = response
	return
}
