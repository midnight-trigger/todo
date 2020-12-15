package definition

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeleteTodoParam_正常系(t *testing.T) {
	e := TestInit()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/todos/:todoId")
	c.SetParamNames("todoId")
	c.SetParamValues("1")

	result, err := CreateDeleteTodoParam(c)
	assert.NoError(t, err)
	assert.IsType(t, &DeleteTodoParam{TodoId: 1}, result)
}

func TestCreateDeleteTodoParam_required(t *testing.T) {
	e := TestInit()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/todos/:todoId")

	_, err := CreateDeleteTodoParam(c)
	assert.EqualError(t, errors.New("TodoIDは必須です"), err.Error())

}
