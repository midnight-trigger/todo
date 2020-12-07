package domain

import (
	"errors"
	"fmt"

	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/infra/cognito"
	"github.com/midnight-trigger/todo/infra/mysql"
	"github.com/midnight-trigger/todo/logger"
)

type User struct {
	Base
	MUsers mysql.IUsers
}

func GetNewUserService() *User {
	user := new(User)
	user.MUsers = mysql.GetNewUser()
	return user
}

func (s *User) PostSigninUser(body *definition.PostSigninUserRequestBody) (r Result) {
	r.New()

	// Cognitoサインイン
	cognito, err := cognito.AdminInitiateAuth(body)
	if err != nil {
		r.CognitoErrorFoundException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	fmt.Println(cognito)

	response := new(definition.PostSigninUserResponse)
	response.IdToken = *cognito.AuthenticationResult.IdToken

	r.Data = response
	return
}

func (s *User) PostUser(body *definition.PostUserRequestBody) (r Result) {
	r.New()

	// Cognitoサインアップ
	cognito, err := cognito.SignUp(body)
	if err != nil {
		r.CognitoErrorFoundException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	// DBインサート
	user := new(mysql.Users)
	s.SetStructOnSameField(body, user)
	user.Id = *cognito.UserSub

	err = s.MUsers.Create(user)
	if err != nil {
		r.ServerErrorException(err, err.Error())
		logger.L.Error(err)
		return
	}

	response := new(definition.PostUserResponse)
	response.Id = *cognito.UserSub
	response.Username = user.Username

	r.Data = response
	return
}
