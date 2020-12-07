package definition

import (
	"github.com/midnight-trigger/todo/api/error_handling"

	"github.com/go-playground/validator/v10"
)

func Validator(inputParams interface{}) (errMessage string, ok bool) {
	ok = true
	validate := validator.New()
	err := validate.Struct(inputParams)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			param := e.Param()
			message := error_handling.GetValidationErrorMessage(field, tag, param)
			if len(message) == 0 {
				message = err.Error()
			}

			errMessage = message
			ok = false
			return
		}
	}

	return
}
