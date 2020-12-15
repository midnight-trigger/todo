package error_handling

import (
	"fmt"
	"net/http"
)

type ErrorHandling struct {
	Code          int
	ErrMessage    string
	RawErrMessage string
	ErrStack      error
}

func (e *ErrorHandling) CognitoErrorFoundException(stackErr error, errMessage string) {
	e.Code = http.StatusBadRequest
	e.ErrMessage = errMessage
	e.setSlackErrorInfo(stackErr, "")
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

func (e *ErrorHandling) EmailIsExistedException(stackErr error) {
	e.Code = http.StatusBadRequest
	e.ErrMessage = "入力頂いたメールアドレスは既に登録されています"
	e.setSlackErrorInfo(stackErr, "")
}

func (e *ErrorHandling) FailureHashedPasswordException(stackErr error) {
	e.Code = http.StatusInternalServerError
	e.ErrMessage = "パスワードの暗号化に失敗しました"
	e.setSlackErrorInfo(stackErr, "")
}

func (e *ErrorHandling) UserNotFoundException(stackErr error) {
	e.Code = http.StatusNotFound
	e.ErrMessage = "対象ユーザーが見つかりません"
	e.setSlackErrorInfo(stackErr, "")
}

func (e *ErrorHandling) TodoNotFoundException(stackErr error) {
	e.Code = http.StatusNotFound
	e.ErrMessage = "対象Todoが見つかりません"
	e.setSlackErrorInfo(stackErr, "")
}

func (e *ErrorHandling) UserIsNotOwnerException(stackErr error) {
	e.Code = http.StatusBadRequest
	e.ErrMessage = "必要な権限がありません"
	e.setSlackErrorInfo(stackErr, "")
}

func GetValidationErrorMessage(field string, tag string, params string) (message string) {
	switch tag {
	case "required":
		switch field {
		case "Username":
			message = "ユーザ名は必須です"
		case "Email":
			message = "メールアドレスは必須です"
		case "Password":
			message = "パスワードは必須です"
		case "Title":
			message = "Todoタイトルは必須です"
		case "Body":
			message = "Todo詳細は必須です"
		case "Status":
			message = "ステータスは必須です"
		case "TodoId":
			message = "TodoIDは必須です"
		}
	case "email":
		switch field {
		case "Email":
			message = "メールアドレスの形式が正しくありません"
		}
	case "gte":
		switch field {
		case "Password":
			message = "パスワードは8文字以上で入力して下さい"
		}
	case "lte":
		switch field {
		case "Username":
			message = fmt.Sprintf("ユーザ名は%s文字以内で入力して下さい", params)
		case "Password":
			message = fmt.Sprintf("パスワードは%s文字以内で入力して下さい", params)
		case "Title":
			message = fmt.Sprintf("Todoタイトルは%s文字以内で入力して下さい", params)
		}
	case "oneof":
		switch field {
		case "Sort":
			message = "ソートはDESC, ASCのいずれかで入力して下さい"
		case "Status":
			message = "ステータスはtodo, progress, finishedのいずれかで入力して下さい"
		}
	}
	return
}
