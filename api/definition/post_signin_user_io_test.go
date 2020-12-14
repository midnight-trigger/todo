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

func TestCreatePostSigninUserRequestBody_正常系(t *testing.T) {
	e := TestInit()
	var data = `{
		"email": "test@test.com",
		"password": "Testtest="
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	result, err := CreatePostSigninUserRequestBody(c)
	assert.NoError(t, err)
	assert.IsType(t, &PostSigninUserRequestBody{}, result)
}

func TestCreatePostSigninUserRequestBody_required(t *testing.T) {
	e := TestInit()
	var data = map[string][]string{
		"email": []string{
			`{
				"email": "",
				"password": "Testtest="
			}`,
			`{
				"email": null,
				"password": "Testtest="
			}`,
			`{
				"password": "Testtest="
			}`,
		},
		"password": []string{
			`{
				"email": "test@test.com",
				"password": ""
			}`,
			`{
				"email": "test@test.com",
				"password": null
			}`,
			`{
				"email": "test@test.com"
			}`,
		},
	}

	for element, value := range data {
		switch element {
		case "email":
			for _, v := range value {
				req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(v))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				_, err := CreatePostSigninUserRequestBody(c)
				assert.EqualError(t, errors.New("メールアドレスは必須です"), err.Error())
			}
		case "password":
			for _, v := range value {
				req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(v))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				_, err := CreatePostSigninUserRequestBody(c)
				assert.EqualError(t, errors.New("パスワードは必須です"), err.Error())
			}
		}
	}
}

func TestCreatePostSigninUserRequestBody_email(t *testing.T) {
	e := TestInit()
	var data = []string{
		// 異常系
		`{
			"email": "test",
			"password": "Testtest="
		}`,
		`{
			"email": "test@",
			"password": "Testtest="
		}`,
		`{
			"email": "test@test",
			"password": "Testtest="
		}`,
		`{
			"email": "test@test.",
			"password": "Testtest="
		}`,
		// 正常系
		`{
			"email": "test@test.com",
			"password": "Testtest="
		}`,
	}

	for index, value := range data {
		if index == len(data)-1 {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			result, err := CreatePostSigninUserRequestBody(c)
			assert.NoError(t, err)
			assert.IsType(t, &PostSigninUserRequestBody{}, result)
		} else {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			_, err := CreatePostSigninUserRequestBody(c)
			assert.EqualError(t, errors.New("メールアドレスの形式が正しくありません"), err.Error())
		}
	}
}

func TestCreatePostSigninUserRequestBody_gte(t *testing.T) {
	e := TestInit()
	var data = []string{
		// 8文字未満（異常系）
		`{
			"email": "test@test.com",
			"password": "T"
		}`,
		`{
			"email": "test@test.com",
			"password": "Testtes"
		}`,
		// 8文字異常（正常系）
		`{
			"email": "test@test.com",
			"password": "Testtest"
		}`,
	}

	for index, value := range data {
		if index == len(data)-1 {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			result, err := CreatePostSigninUserRequestBody(c)
			assert.NoError(t, err)
			assert.IsType(t, &PostSigninUserRequestBody{}, result)
		} else {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			_, err := CreatePostSigninUserRequestBody(c)
			assert.EqualError(t, errors.New("パスワードは8文字以上で入力して下さい"), err.Error())
		}
	}
}

func TestCreatePostSigninUserRequestBody_lte(t *testing.T) {
	e := TestInit()
	var data = []string{
		// 256文字（異常系）
		`{
			"email": "test@test.com",
			"password": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}`,
		// 255文字（正常系）
		`{
			"email": "test@test.com",
			"password": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}`,
	}

	for index, value := range data {
		if index == len(data)-1 {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			result, err := CreatePostSigninUserRequestBody(c)
			assert.NoError(t, err)
			assert.IsType(t, &PostSigninUserRequestBody{}, result)
		} else {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signin", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			_, err := CreatePostSigninUserRequestBody(c)
			assert.EqualError(t, errors.New("パスワードは255文字以内で入力して下さい"), err.Error())
		}
	}
}
