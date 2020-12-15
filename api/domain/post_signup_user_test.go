package domain

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/golang/mock/gomock"
	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/api/error_handling"
	"github.com/stretchr/testify/assert"

	"github.com/midnight-trigger/todo/infra/cognito/mock_cognito"
	"github.com/midnight-trigger/todo/infra/mysql"
	"github.com/midnight-trigger/todo/infra/mysql/mock_mysql"
)

func TestPostSignupUser_正常系(t *testing.T) {
	s := GetNewUserService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PostSignupUserRequestBody)
	payload.Username = "test-user"
	payload.Email = "test@test.com"
	payload.Password = "Testtest="
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	cognito := new(cognitoidentityprovider.SignUpOutput)
	cognito.UserSub = &userId

	user := new(mysql.Users)
	user.Id = *cognito.UserSub
	user.Username = "test-user"

	createdUser := new(mysql.Users)
	createdUser.Id = *cognito.UserSub
	createdUser.Username = "test-user"

	response := new(definition.PostSignupUserResponse)
	s.SetStructOnSameField(createdUser, response)

	mockedCognito := mock_cognito.NewMockICognito(ctrl)
	mockedUsers := mock_mysql.NewMockIUsers(ctrl)
	gomock.InOrder(
		mockedCognito.EXPECT().SignUp(payload).Return(cognito, nil),
		mockedUsers.EXPECT().Create(user).Return(createdUser, nil),
	)

	domain := new(User)
	domain.MCognito = mockedCognito
	domain.MUsers = mockedUsers

	result := domain.PostSignupUser(payload)
	assert.Equal(t, 200, result.Code)
	assert.Equal(t, response, result.Data.(*definition.PostSignupUserResponse))
}

func TestPostSignupUser_Cognitoで問題が起きた場合エラーを返すか検証(t *testing.T) {
	s := GetNewUserService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PostSignupUserRequestBody)
	payload.Username = "test-user"
	payload.Email = "test@test.com"
	payload.Password = "Testtest="

	mockedCognito := mock_cognito.NewMockICognito(ctrl)
	mockedCognito.EXPECT().SignUp(payload).Return(&cognitoidentityprovider.SignUpOutput{}, &cognitoidentityprovider.InternalErrorException{})

	expect := new(error_handling.ErrorHandling)
	expect.Code = 400
	expect.ErrMessage = "InternalErrorException: "
	expect.ErrStack = errors.New("")
	domain := new(User)
	domain.MCognito = mockedCognito

	result := domain.PostSignupUser(payload)
	assert.Equal(t, *expect, result.ErrorHandling)
}

func TestPostSignupUser_サーバで問題が起きた場合エラーを返すか検証(t *testing.T) {
	s := GetNewUserService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PostSignupUserRequestBody)
	payload.Username = "test-user"
	payload.Email = "test@test.com"
	payload.Password = "Testtest="
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	cognito := new(cognitoidentityprovider.SignUpOutput)
	cognito.UserSub = &userId

	user := new(mysql.Users)
	user.Id = *cognito.UserSub
	user.Username = "test-user"

	expect := new(error_handling.ErrorHandling)
	expect.Code = 500
	expect.ErrMessage = "サーバーエラー"
	expect.ErrStack = errors.New("")
	expect.RawErrMessage = "not implemented"

	mockedCognito := mock_cognito.NewMockICognito(ctrl)
	mockedUsers := mock_mysql.NewMockIUsers(ctrl)
	gomock.InOrder(
		mockedCognito.EXPECT().SignUp(payload).Return(cognito, nil),
		mockedUsers.EXPECT().Create(user).Return(&mysql.Users{}, errors.New("not implemented")),
	)

	domain := new(User)
	domain.MCognito = mockedCognito
	domain.MUsers = mockedUsers

	result := domain.PostSignupUser(payload)
	assert.Equal(t, *expect, result.ErrorHandling)
}
