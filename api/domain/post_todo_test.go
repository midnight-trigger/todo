package domain

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/api/error_handling"

	"github.com/midnight-trigger/todo/infra/mysql"
	"github.com/midnight-trigger/todo/infra/mysql/mock_mysql"

	"github.com/stretchr/testify/assert"
)

func TestPostTodo_正常系(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PostTodoRequestBody)
	payload.Title = "Test Title"
	payload.Body = "Test Body"

	user := new(mysql.Users)
	user.Id = "1802f638-53f2-4848-9859-a54a2bdf5163"
	user.Username = "test-user"

	todo := new(mysql.Todos)
	s.SetStructOnSameField(payload, todo)
	todo.UserId = "1802f638-53f2-4848-9859-a54a2bdf5163"

	createdTodo := new(mysql.Todos)
	createdTodo.Id = 1
	createdTodo.Title = todo.Title
	createdTodo.Body = todo.Body
	createdTodo.UserId = todo.UserId
	createdTodo.CreatedAt = time.Now()
	createdTodo.UpdatedAt = time.Now()

	response := new(definition.PostTodoResponse)
	s.SetStructOnSameField(createdTodo, response)

	mockedUsers := mock_mysql.NewMockIUsers(ctrl)
	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	gomock.InOrder(
		mockedUsers.EXPECT().FindById("1802f638-53f2-4848-9859-a54a2bdf5163").Return(*user, nil),
		mockedTodos.EXPECT().Create(todo).Return(createdTodo, nil),
	)

	domain := new(Todo)
	domain.MUsers = mockedUsers
	domain.MTodos = mockedTodos

	result := domain.PostTodo(payload, todo.UserId)
	assert.Equal(t, 200, result.Code)
	assert.Equal(t, response, result.Data.(*definition.PostTodoResponse))
}

func TestPostTodo_jwtから解析したUserIdがDB上に存在しない場合エラーを返すか検証(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PostTodoRequestBody)
	payload.Title = "Test Title"
	payload.Body = "Test Body"
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	mockedUsers := mock_mysql.NewMockIUsers(ctrl)
	mockedUsers.EXPECT().FindById("1802f638-53f2-4848-9859-a54a2bdf5163").Return(mysql.Users{}, gorm.ErrRecordNotFound)

	expect := new(error_handling.ErrorHandling)
	expect.Code = 404
	expect.ErrMessage = "対象ユーザーが見つかりません"
	expect.ErrStack = errors.New("")

	domain := new(Todo)
	domain.MUsers = mockedUsers

	result := domain.PostTodo(payload, userId)
	assert.Equal(t, *expect, result.ErrorHandling)
}

func TestPostTodo_サーバで問題が起きた場合エラーを返すか検証(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PostTodoRequestBody)
	payload.Title = "Test Title"
	payload.Body = "Test Body"

	user := new(mysql.Users)
	user.Id = "1802f638-53f2-4848-9859-a54a2bdf5163"
	user.Username = "test-user"

	todo := new(mysql.Todos)
	s.SetStructOnSameField(payload, todo)
	todo.UserId = "1802f638-53f2-4848-9859-a54a2bdf5163"

	expect := new(error_handling.ErrorHandling)
	expect.Code = 500
	expect.ErrMessage = "サーバーエラー"
	expect.ErrStack = errors.New("")
	expect.RawErrMessage = "not implemented"

	mockedUsers := mock_mysql.NewMockIUsers(ctrl)
	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	for i := 0; i < 2; i++ {
		switch i {
		case 0:
			mockedUsers.EXPECT().FindById(user.Id).Return(*user, errors.New("not implemented")).Times(2)
		case 1:
			gomock.InOrder(
				mockedUsers.EXPECT().FindById("1802f638-53f2-4848-9859-a54a2bdf5163").Return(*user, nil),
				mockedTodos.EXPECT().Create(todo).Return(&mysql.Todos{}, errors.New("not implemented")).Times(1),
			)
		}
		domain := new(Todo)
		domain.MUsers = mockedUsers

		result := domain.PostTodo(payload, user.Id)
		assert.Equal(t, *expect, result.ErrorHandling)
	}
}
