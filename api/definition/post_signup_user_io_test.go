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

func TestCreatePostSignupUserRequestBody_正常系(t *testing.T) {
	e := TestInit()
	var data = `{
		"username": "test-user",
		"email": "test@test.com",
		"password": "Testtest="
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	result, err := CreatePostSignupUserRequestBody(c)
	assert.NoError(t, err)
	assert.IsType(t, &PostSignupUserRequestBody{}, result)
}

func TestCreatePostSignupUserRequestBody_required(t *testing.T) {
	e := TestInit()
	var data = map[string][]string{
		"username": []string{
			`{
				"username": "",
				"email": "test@test.com",
				"password": "Testtest="
			}`,
			`{
				"username": null,
				"email": "test@test.com",
				"password": "Testtest="
			}`,
			`{
				"email": "test@test.com",
				"password": "Testtest="
			}`,
		},
		"email": []string{
			`{
				"username": "test-user",
				"email": "",
				"password": "Testtest="
			}`,
			`{
				"username": "test-user",
				"email": null,
				"password": "Testtest="
			}`,
			`{
				"username": "test-user",
				"password": "Testtest="
			}`,
		},
		"password": []string{
			`{
				"username": "test-user",
				"email": "test@test.com",
				"password": ""
			}`,
			`{
				"username": "test-user",
				"email": "test@test.com",
				"password": null
			}`,
			`{
				"username": "test-user",
				"email": "test@test.com"
			}`,
		},
	}

	for element, value := range data {
		switch element {
		case "username":
			for _, v := range value {
				req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(v))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				_, err := CreatePostSignupUserRequestBody(c)
				assert.EqualError(t, errors.New("ユーザ名は必須です"), err.Error())
			}
		case "email":
			for _, v := range value {
				req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(v))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				_, err := CreatePostSignupUserRequestBody(c)
				assert.EqualError(t, errors.New("メールアドレスは必須です"), err.Error())
			}
		case "password":
			for _, v := range value {
				req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(v))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				_, err := CreatePostSignupUserRequestBody(c)
				assert.EqualError(t, errors.New("パスワードは必須です"), err.Error())
			}
		}
	}
}

func TestCreatePostSignupUserRequestBody_email(t *testing.T) {
	e := TestInit()
	var data = []string{
		// 異常系
		`{
			"username": "test-user",
			"email": "test",
			"password": "Testtest="
		}`,
		`{
			"username": "test-user",
			"email": "test@",
			"password": "Testtest="
		}`,
		`{
			"username": "test-user",
			"email": "test@test",
			"password": "Testtest="
		}`,
		`{
			"username": "test-user",
			"email": "test@test.",
			"password": "Testtest="
		}`,
		// 正常系
		`{
			"username": "test-user",
			"email": "test@test.com",
			"password": "Testtest="
		}`,
	}

	for index, value := range data {
		if index == len(data)-1 {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			result, err := CreatePostSignupUserRequestBody(c)
			assert.NoError(t, err)
			assert.IsType(t, &PostSignupUserRequestBody{}, result)
		} else {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			_, err := CreatePostSignupUserRequestBody(c)
			assert.EqualError(t, errors.New("メールアドレスの形式が正しくありません"), err.Error())
		}
	}
}

func TestCreatePostSignupUserRequestBody_gte(t *testing.T) {
	e := TestInit()
	var data = []string{
		// 8文字未満（異常系）
		`{
			"username": "test-user",
			"email": "test@test.com",
			"password": "T"
		}`,
		`{
			"username": "test-user",
			"email": "test@test.com",
			"password": "Testtes"
		}`,
		// 8文字以上（正常系）
		`{
			"username": "test-user",
			"email": "test@test.com",
			"password": "Testtest"
		}`,
	}

	for index, value := range data {
		if index == len(data)-1 {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			result, err := CreatePostSignupUserRequestBody(c)
			assert.NoError(t, err)
			assert.IsType(t, &PostSignupUserRequestBody{}, result)
		} else {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(value))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			_, err := CreatePostSignupUserRequestBody(c)
			assert.EqualError(t, errors.New("パスワードは8文字以上で入力して下さい"), err.Error())
		}
	}
}

func TestCreatePostSignupUserRequestBody_lte(t *testing.T) {
	e := TestInit()
	var data = map[string][]string{
		"username": []string{
			// 31文字（異常系）
			`{
				"username": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				"email": "test@test.com",
				"password": "Testtest="
			}`,
			// 30文字（正常系）
			`{
				"username": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				"email": "test@test.com",
				"password": "Testtest="
			}`,
		},
		"password": []string{
			// 256文字（異常系）
			`{
				"username": "test-user",
				"email": "test@test.com",
				"password": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			}`,
			// 255文字（正常系）
			`{
				"username": "test-user",
				"email": "test@test.com",
				"password": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			}`,
		},
	}

	for element, value := range data {
		switch element {
		case "username":
			for index, v := range value {
				if index%2 != 0 {
					req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(v))
					req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
					rec := httptest.NewRecorder()
					c := e.NewContext(req, rec)

					result, err := CreatePostSignupUserRequestBody(c)
					assert.NoError(t, err)
					assert.IsType(t, &PostSignupUserRequestBody{}, result)
				} else {
					req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(v))
					req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
					rec := httptest.NewRecorder()
					c := e.NewContext(req, rec)
					_, err := CreatePostSignupUserRequestBody(c)
					assert.EqualError(t, errors.New("ユーザ名は30文字以内で入力して下さい"), err.Error())
				}
			}
		case "password":
			for index, v := range value {
				if index%2 != 0 {
					req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(v))
					req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
					rec := httptest.NewRecorder()
					c := e.NewContext(req, rec)

					result, err := CreatePostSignupUserRequestBody(c)
					assert.NoError(t, err)
					assert.IsType(t, &PostSignupUserRequestBody{}, result)
				} else {
					req := httptest.NewRequest(http.MethodPost, "/api/v1/users/signup", strings.NewReader(v))
					req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
					rec := httptest.NewRecorder()
					c := e.NewContext(req, rec)
					_, err := CreatePostSignupUserRequestBody(c)
					assert.EqualError(t, errors.New("パスワードは255文字以内で入力して下さい"), err.Error())
				}
			}
		}
	}
}
