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

func TestCreatePutTodoRequestBodyAndParam_正常系(t *testing.T) {
	e := TestInit()
	var data = `{
		"title": "Test Title",
		"body": "Test Body"
	}`

	req := httptest.NewRequest(http.MethodPut, "/api/v1/todos/:todoId", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_, result, err := CreatePutTodoRequestBodyAndParam(c)
	assert.NoError(t, err)
	assert.IsType(t, &PutTodoRequestBody{}, result)
}

func TestCreatePutTodoRequestBodyAndParam_required(t *testing.T) {
	e := TestInit()
	var data = map[string][]string{
		"title": []string{
			`{
				"title": "",
				"body": "Test Body"
			}`,
			`{
				"title": null,
				"body": "Test Body"
			}`,
			`{
				"body": "Test Body"
			}`,
		},
		"body": []string{
			`{
				"title": "Test Title",
				"body": ""
			}`,
			`{
				"title": "Test Title",
				"body": null
			}`,
			`{
				"title": "Test Title"
			}`,
		},
	}

	for element, value := range data {
		switch element {
		case "title":
			for _, v := range value {
				req := httptest.NewRequest(http.MethodPut, "/api/v1//todos/:todoId", strings.NewReader(v))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				_, _, err := CreatePutTodoRequestBodyAndParam(c)
				assert.EqualError(t, errors.New("Todoタイトルは必須です"), err.Error())
			}
		case "body":
			for _, v := range value {
				req := httptest.NewRequest(http.MethodPut, "/api/v1//todos/:todoId", strings.NewReader(v))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				_, _, err := CreatePutTodoRequestBodyAndParam(c)
				assert.EqualError(t, errors.New("Todo詳細は必須です"), err.Error())
			}
		}
	}
}

func TestCreatePutTodoRequestBody_lte(t *testing.T) {
	e := TestInit()
	var data = []string{
		// 256文字（異常系）
		`{
			"title": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"body": "Test Body"
		}`,
		// 255文字（正常系）
		`{
			"title": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"body": "Test Body"
		}`,
	}

	for index, value := range data {
		if index != 0 {
			req := httptest.NewRequest(http.MethodPut, "/api/v1/todos/:todoId", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			_, result, err := CreatePutTodoRequestBodyAndParam(c)
			assert.NoError(t, err)
			assert.IsType(t, &PutTodoRequestBody{}, result)
		} else {
			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			_, _, err := CreatePutTodoRequestBodyAndParam(c)
			assert.EqualError(t, errors.New("Todoタイトルは255文字以内で入力して下さい"), err.Error())
		}
	}
}
