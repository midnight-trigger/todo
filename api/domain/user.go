package domain

import (
	"errors"

	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/infra/cognito"
	"github.com/midnight-trigger/todo/infra/mysql"
	"github.com/midnight-trigger/todo/logger"
)

type User struct {
	Base
	MUsers   mysql.IUsers
	MCognito cognito.ICognito
}

func GetNewUserService() *User {
	user := new(User)
	user.MUsers = mysql.GetNewUser()
	user.MCognito = cognito.GetNewCognito()
	return user
}

// ログイン
func (s *User) PostSigninUser(body *definition.PostSigninUserRequestBody) (r Result) {
	r.New()

	// Cognitoログイン
	cognito, err := s.MCognito.AdminInitiateAuth(body)
	if err != nil {
		r.CognitoErrorFoundException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	// レスポンス作成
	response := new(definition.PostSigninUserResponse)
	response.IdToken = *cognito.AuthenticationResult.IdToken

	r.Data = response
	return
}

// 会員登録
func (s *User) PostSignupUser(body *definition.PostSignupUserRequestBody) (r Result) {
	r.New()

	// Cognitoサインアップ
	cognito, err := s.MCognito.SignUp(body)
	if err != nil {
		r.CognitoErrorFoundException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	// DBインサート
	user := new(mysql.Users)
	s.SetStructOnSameField(body, user)
	user.Id = *cognito.UserSub

	createdUser, err := s.MUsers.Create(user)
	if err != nil {
		r.ServerErrorException(errors.New(""), err.Error())
		logger.L.Error(err)
		return
	}

	// レスポンス作成
	response := new(definition.PostSignupUserResponse)
	s.SetStructOnSameField(createdUser, response)

	r.Data = response
	return
}
