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

func (s *Todo) GetTodos(params *definition.GetTodosParam, userId string) (r Result) {
	r.New()

	// 検索条件をもとにTodo一覧を取得
	todos, err := s.MTodos.FindByQuery(params, userId)
	if err != nil {
		r.ServerErrorException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	var responses []*definition.GetTodosResponse
	for _, todo := range todos {
		response := new(definition.GetTodosResponse)
		s.SetStructOnSameField(todo, response)
		responses = append(responses, response)
	}
	r.Data = responses

	pagination := new(Pagination)
	s.SetStructOnSameField(params, pagination)

	pagination.Total, err = s.MTodos.GetTotalCount(params, userId)
	if err != nil {
		r.ServerErrorException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}
	r.Pagination = pagination
	return
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

func (s *Todo) PutTodo(param *definition.PutTodoParam, body *definition.PutTodoRequestBody, userId string) (r Result) {
	r.New()

	// Todo存在チェック
	oldParams, err := s.MTodos.FindById(param.TodoId)
	if gorm.IsRecordNotFoundError(err) {
		r.TodoNotFoundException(errors.New(""))
		logger.L.Error(r.ErrMessage)
		return
	}
	if err != nil {
		r.ServerErrorException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	// ログイン中ユーザのDB更新権限チェック
	if oldParams.UserId != userId {
		r.UserIsNotOwnerException(errors.New(""))
		logger.L.Error(r.ErrMessage)
		return
	}

	// DB更新
	updateParams := map[string]interface{}{
		"Title": body.Title,
		"Body":  body.Body,
	}
	err = s.MTodos.Update(oldParams, updateParams)
	if err != nil {
		r.ServerErrorException(err, err.Error())
		logger.L.Error(err)
		return
	}

	response := new(definition.PutTodoResponse)
	s.SetStructOnSameField(oldParams, response)
	s.SetStructOnSameField(body, response)
	r.Data = response
	return
}

func (s *Todo) PatchTodo(param *definition.PatchTodoParam, body *definition.PatchTodoRequestBody, userId string) (r Result) {
	r.New()

	// Todo存在チェック
	oldParams, err := s.MTodos.FindById(param.TodoId)
	if gorm.IsRecordNotFoundError(err) {
		r.TodoNotFoundException(errors.New(""))
		logger.L.Error(r.ErrMessage)
		return
	}
	if err != nil {
		r.ServerErrorException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	// ログイン中ユーザのDB更新権限チェック
	if oldParams.UserId != userId {
		r.UserIsNotOwnerException(errors.New(""))
		logger.L.Error(r.ErrMessage)
		return
	}

	// DB更新
	updateParam := map[string]interface{}{
		"Status": body.Status,
	}
	err = s.MTodos.Update(oldParams, updateParam)
	if err != nil {
		r.ServerErrorException(err, err.Error())
		logger.L.Error(err)
		return
	}

	response := new(definition.PatchTodoResponse)
	s.SetStructOnSameField(oldParams, response)
	s.SetStructOnSameField(body, response)
	r.Data = response
	return
}

func (s *Todo) DeleteTodo(param *definition.DeleteTodoParam, userId string) (r Result) {
	r.New()

	// Todo存在チェック
	todo, err := s.MTodos.FindById(param.TodoId)
	if gorm.IsRecordNotFoundError(err) {
		r.TodoNotFoundException(errors.New(""))
		logger.L.Error(r.ErrMessage)
		return
	}
	if err != nil {
		r.ServerErrorException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	// ログイン中ユーザのDB削除権限チェック
	if todo.UserId != userId {
		r.UserIsNotOwnerException(errors.New(""))
		logger.L.Error(r.ErrMessage)
		return
	}

	// レコード削除
	err = s.MTodos.Delete(&todo)
	if err != nil {
		r.ServerErrorException(err, err.Error())
		logger.L.Error(err)
		return
	}

	response := new(definition.DeleteTodoResponse)
	response.Id = param.TodoId
	r.Data = response
	return
}
