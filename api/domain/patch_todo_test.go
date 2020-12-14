package domain

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/api/error_handling"

	"github.com/midnight-trigger/todo/infra/mysql"
	"github.com/midnight-trigger/todo/infra/mysql/mock_mysql"

	"github.com/stretchr/testify/assert"
)

func TestPatchTodo_正常系(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PatchTodoRequestBody)
	payload.Status = "progress"
	param := new(definition.PatchTodoParam)
	param.TodoId = 1
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	oldTodo := new(mysql.Todos)
	oldTodo.Id = param.TodoId
	oldTodo.UserId = userId
	oldTodo.Title = "Old Title"
	oldTodo.Body = "Old Body"
	oldTodo.Status = "todo"

	updatedTodo := map[string]interface{}{
		"Status": payload.Status,
	}

	response := new(definition.PatchTodoResponse)
	s.SetStructOnSameField(oldTodo, response)
	s.SetStructOnSameField(payload, response)

	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	gomock.InOrder(
		mockedTodos.EXPECT().FindById(param.TodoId).Return(*oldTodo, nil),
		mockedTodos.EXPECT().Update(*oldTodo, updatedTodo).Return(nil),
	)

	domain := new(Todo)
	domain.MTodos = mockedTodos

	result := domain.PatchTodo(param, payload, userId)
	assert.Equal(t, 200, result.Code)
	assert.Equal(t, response, result.Data.(*definition.PatchTodoResponse))
}

func TestPatchTodo_パスパラメータのTodoIdに紐づくレコードがDB上に存在しない場合エラーを返すか検証(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PatchTodoRequestBody)
	payload.Status = "progress"
	param := new(definition.PatchTodoParam)
	param.TodoId = 1
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	mockedTodos.EXPECT().FindById(param.TodoId).Return(mysql.Todos{}, gorm.ErrRecordNotFound)

	expect := new(error_handling.ErrorHandling)
	expect.Code = 404
	expect.ErrMessage = "対象Todoが見つかりません"
	expect.ErrStack = errors.New("")

	domain := new(Todo)
	domain.MTodos = mockedTodos

	result := domain.PatchTodo(param, payload, userId)
	assert.Equal(t, *expect, result.ErrorHandling)
}

func TestPatchTodo_ログインユーザにDB更新権限が無い場合エラーを返すか検証(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PatchTodoRequestBody)
	payload.Status = "progress"
	param := new(definition.PatchTodoParam)
	param.TodoId = 1
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	todo := new(mysql.Todos)
	todo.Id = param.TodoId
	todo.UserId = "1802f638-53f2-4848-9859-a54a2bdf5160"
	todo.Title = "Title"
	todo.Body = "Body"
	todo.Status = "todo"

	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	mockedTodos.EXPECT().FindById(param.TodoId).Return(*todo, nil)

	expect := new(error_handling.ErrorHandling)
	expect.Code = 400
	expect.ErrMessage = "必要な権限がありません"
	expect.ErrStack = errors.New("")

	domain := new(Todo)
	domain.MTodos = mockedTodos

	result := domain.PatchTodo(param, payload, userId)
	assert.Equal(t, *expect, result.ErrorHandling)
}

func TestPatchTodo_サーバで問題が起きた場合サーバエラーを返すか検証(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PatchTodoRequestBody)
	payload.Status = "progress"
	param := new(definition.PatchTodoParam)
	param.TodoId = 1
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	oldTodo := new(mysql.Todos)
	oldTodo.Id = param.TodoId
	oldTodo.UserId = userId
	oldTodo.Title = "Old Title"
	oldTodo.Body = "Old Body"
	oldTodo.Status = "todo"

	updatedTodo := map[string]interface{}{
		"Status": payload.Status,
	}

	expect := new(error_handling.ErrorHandling)
	expect.Code = 500
	expect.ErrMessage = "サーバーエラー"
	expect.ErrStack = errors.New("")
	expect.RawErrMessage = "not implemented"

	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	for i := 0; i < 2; i++ {
		switch i {
		case 0:
			mockedTodos.EXPECT().FindById(param.TodoId).Return(*oldTodo, errors.New("not implemented")).Times(2)
		case 1:
			gomock.InOrder(
				mockedTodos.EXPECT().FindById(param.TodoId).Return(*oldTodo, nil),
				mockedTodos.EXPECT().Update(*oldTodo, updatedTodo).Return(errors.New("not implemented")).Times(1),
			)
		}
		domain := new(Todo)
		domain.MTodos = mockedTodos

		result := domain.PatchTodo(param, payload, userId)
		assert.Equal(t, *expect, result.ErrorHandling)
	}
}
