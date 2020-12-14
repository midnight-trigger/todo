package definition

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatchTodoRequestBodyAndParam_正常系(t *testing.T) {
	e := TestInit()
	var data = `{
		"status": "progress"
	}`

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/todos/:todoId", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_, result, err := CreatePatchTodoRequestBodyAndParam(c)
	assert.NoError(t, err)
	assert.IsType(t, &PatchTodoRequestBody{}, result)
}

func TestCreatePatchTodoRequestBodyAndParam_required(t *testing.T) {
	e := TestInit()
	var data = []string{
		// 異常系
		`{
			"status": ""
		}`,
		`{
			"status": null
		}`,
		`{
		}`,
		// 正常系
		`{
			"status": "progress"
		}`,
	}

	for index, value := range data {
		if index == len(data)-1 {
			req := httptest.NewRequest(http.MethodPatch, "/api/v1/todos/:todoId", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			_, result, err := CreatePatchTodoRequestBodyAndParam(c)
			assert.NoError(t, err)
			assert.IsType(t, &PatchTodoRequestBody{}, result)
		} else {
			req := httptest.NewRequest(http.MethodPatch, "/api/v1/todos/:todoId", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			_, _, err := CreatePatchTodoRequestBodyAndParam(c)
			assert.EqualError(t, errors.New("ステータスは必須です"), err.Error())
		}
	}
}

func TestCreatePatchTodoRequestBodyAndParam_oneof(t *testing.T) {
	e := TestInit()
	var data = []string{
		// 異常系
		`{
			"status": "test"
		}`,
		// 正常系
		`{
			"status": "todo"
		}`,
		`{
			"status": "progress"
		}`,
		`{
			"status": "finished"
		}`,
	}

	for index, value := range data {
		if index != 0 {
			req := httptest.NewRequest(http.MethodPatch, "/api/v1/todos/:todoId", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			_, result, err := CreatePatchTodoRequestBodyAndParam(c)
			assert.NoError(t, err)
			assert.IsType(t, &PatchTodoRequestBody{}, result)
		} else {
			req := httptest.NewRequest(http.MethodPatch, "/api/v1/todos/:todoId", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			_, _, err := CreatePatchTodoRequestBodyAndParam(c)
			assert.EqualError(t, errors.New("ステータスはtodo, progress, finishedのいずれかで入力して下さい"), err.Error())
		}
	}
}
