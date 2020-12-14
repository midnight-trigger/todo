package domain

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/api/error_handling"

	"github.com/midnight-trigger/todo/infra/mysql"
	"github.com/midnight-trigger/todo/infra/mysql/mock_mysql"

	"github.com/stretchr/testify/assert"
)

func TestGetTodos_サーバで問題が起きた場合サーバエラーを返すか検証(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	param := new(definition.GetTodosParam)
	param.Offset = 0
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	status := []string{"todo", "progress", "finished"}
	var todos []*mysql.Todos
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 100; i++ {
		todo := new(mysql.Todos)
		todo.Id = int64(i)
		todo.UserId = userId
		todo.Title = "Title" + string(i)
		todo.Body = "Body" + string(i)
		todo.Status = status[rand.Intn(len(status))]
		todos = append(todos, todo)
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
			mockedTodos.EXPECT().FindByQuery(param, userId).Return([]mysql.Todos{}, errors.New("not implemented")).Times(2)
		case 1:
			gomock.InOrder(
				mockedTodos.EXPECT().FindByQuery(param, userId).Return([]mysql.Todos{}, nil),
				mockedTodos.EXPECT().GetTotalCount(param, userId).Return(0, errors.New("not implemented")).Times(1),
			)
		}
		domain := new(Todo)
		domain.MTodos = mockedTodos

		result := domain.GetTodos(param, userId)
		assert.Equal(t, *expect, result.ErrorHandling)
	}
}
