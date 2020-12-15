package definition

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreateGetTodosParam_正常系(t *testing.T) {
	e := TestInit()

	q := make(url.Values)
	q.Set("limit", "10")
	q.Set("offset", "0")
	q.Set("sort", "DESC")
	q.Set("title", "Test")
	q.Set("body", "Test")
	q.Set("status", "todo")
	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	result, err := CreateGetTodosParam(c)
	assert.NoError(t, err)
	assert.IsType(t, &GetTodosParam{}, result)
}

func TestCreateGetTodosParam_lte(t *testing.T) {
	e := TestInit()

	elements := []string{
		// 256文字（異常系）
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		// 255文字（正常系）
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}

	for index, element := range elements {
		q := make(url.Values)
		q.Set("title", element)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/todos?"+q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		result, err := CreateGetTodosParam(c)
		if index != 0 {
			assert.NoError(t, err)
			assert.IsType(t, &GetTodosParam{}, result)
		} else {
			assert.EqualError(t, errors.New("Todoタイトルは255文字以内で入力して下さい"), err.Error())
		}
	}
}

func TestCreateGetTodosParam_oneof(t *testing.T) {
	e := TestInit()

	data := map[string][]string{
		"sort": []string{
			// 異常系
			"test",
			// 正常系
			"DESC",
			"ASC",
		},
		"status": []string{
			// 異常系
			"test",
			// 正常系
			"",
			"todo",
			"progress",
			"finished",
		},
	}

	for element, value := range data {
		q := make(url.Values)
		switch element {
		case "sort":
			for index, v := range value {
				q.Set("sort", v)
				req := httptest.NewRequest(http.MethodGet, "/api/v1/todos?"+q.Encode(), nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				result, err := CreateGetTodosParam(c)
				if index != 0 {
					assert.NoError(t, err)
					assert.IsType(t, &GetTodosParam{}, result)
				} else {
					assert.EqualError(t, errors.New("ソートはDESC, ASCのいずれかで入力して下さい"), err.Error())
				}
			}
		case "status":
			for index, v := range value {
				q.Set("status", v)
				req := httptest.NewRequest(http.MethodGet, "/api/v1/todos?"+q.Encode(), nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				result, err := CreateGetTodosParam(c)
				if index != 0 {
					assert.NoError(t, err)
					assert.IsType(t, &GetTodosParam{}, result)
				} else {
					assert.EqualError(t, errors.New("ステータスはtodo, progress, finishedのいずれかで入力して下さい"), err.Error())
				}
			}
		}
	}
}
