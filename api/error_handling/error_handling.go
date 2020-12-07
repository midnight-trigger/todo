package error_handling

import (
	"net/http"
)

type ErrorHandling struct {
	Code          int
	ErrMessage    string
	RawErrMessage string
	ErrStack      error
}

func (e *ErrorHandling) setSlackErrorInfo(stackErr error, rawErrMessage string) {
	e.ErrStack = stackErr
	e.RawErrMessage = rawErrMessage
}

func (e *ErrorHandling) ServerErrorException(stackErr error, rawErrMessage string) {
	e.Code = http.StatusInternalServerError
	e.ErrMessage = "サーバーエラー"
	e.setSlackErrorInfo(stackErr, rawErrMessage)
}

func (e *ErrorHandling) ValidationException(stackErr error, errMessage string) {
	e.Code = http.StatusBadRequest
	e.ErrMessage = errMessage
	e.setSlackErrorInfo(stackErr, "")
}

func GetValidationErrorMessage(field string, tag string, params string) (message string) {
	switch tag {
	case "required":
		switch field {
		case "UserId":
			message = "ユーザーIDは必須です"
		}
	}
	return
}
