package domain

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/api/error_handling"
	"github.com/stretchr/testify/assert"

	"github.com/midnight-trigger/todo/infra/cognito/mock_cognito"
)

func TestPostSigninUser_正常系(t *testing.T) {
	s := GetNewUserService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PostSigninUserRequestBody)
	payload.Email = "test@test.com"
	payload.Password = "Testtest="
	token := "eyJraWQiOiJabWFmTlIzb1lvaWtXQVU4dzJoSmQ1NzZIXC9cL1V4R1FsZTRPVlRHaUwzRGs9IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIwMzllMDk0MC1lMzM2LTRjMDItOTFiMS03MmJhYjZmOTY1OGEiLCJhdWQiOiIzYmxqMzBrc2JqZDlvNGdndDgycmZkZGtnYSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJldmVudF9pZCI6IjI2NzJhNzQyLTM3ZWQtNGQ0ZC1hNGQ2LWRmNGUyODk3YTlhZSIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNjA3OTEzOTk0LCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAuYXAtbm9ydGhlYXN0LTEuYW1hem9uYXdzLmNvbVwvYXAtbm9ydGhlYXN0LTFfNXJIVzQxODc2IiwiY29nbml0bzp1c2VybmFtZSI6IjAzOWUwOTQwLWUzMzYtNGMwMi05MWIxLTcyYmFiNmY5NjU4YSIsImV4cCI6MTYwNzkxNzU5NCwiaWF0IjoxNjA3OTEzOTk1LCJlbWFpbCI6Imp1bnlhNDMxODYyOUB5YWhvby5jby5qcCJ9.R4VmM-i0z6yKjQgg9nhgcmEiF2-rBARUmjJ_uUiX4M6WctN4Oyr-6Y8Q9B9LL-HPZnBQvBzAIp4bp3yksn0em6sVMfNp3puAu-gJ3-4p_5Np8hST_aYlZ5Vwea7-tNjcCEFiyavNA9rIFyeGdawM7Ukk9aE27cd8bb1FvBgJzhQ9BGLCb3kMRbWTc7qU7VFzHOHRbMjTbQ_JqwHHcFjSoVXr3wsgfkUUV1AlyjIAolwufGmMWxOU2od-_-BZzoJwGqMu2-WHYplHoIwe01zByXDQw7HdFFE1wDZazZ9DaB-KCC-LuR_KPmuR_9kXnscJWnvKW8EQHV9gabSmpXbTmg"

	cognito := new(cognitoidentityprovider.AdminInitiateAuthOutput)
	authenticationResult := new(cognitoidentityprovider.AuthenticationResultType)
	authenticationResult.IdToken = &token
	cognito.AuthenticationResult = authenticationResult

	response := new(definition.PostSigninUserResponse)
	response.IdToken = *cognito.AuthenticationResult.IdToken

	mockedCognito := mock_cognito.NewMockICognito(ctrl)
	mockedCognito.EXPECT().AdminInitiateAuth(payload).Return(cognito, nil)

	domain := new(User)
	domain.MCognito = mockedCognito

	result := domain.PostSigninUser(payload)
	assert.Equal(t, 200, result.Code)
	assert.Equal(t, response, result.Data.(*definition.PostSigninUserResponse))
}

func TestPostSigninUser_Cognitoで問題が起きた場合エラーを返すか検証(t *testing.T) {
	s := GetNewUserService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	payload := new(definition.PostSigninUserRequestBody)
	payload.Email = "test@test.com"
	payload.Password = "Testtest="

	mockedCognito := mock_cognito.NewMockICognito(ctrl)
	mockedCognito.EXPECT().AdminInitiateAuth(payload).Return(&cognitoidentityprovider.AdminInitiateAuthOutput{}, &cognitoidentityprovider.InternalErrorException{})

	expect := new(error_handling.ErrorHandling)
	expect.Code = 400
	expect.ErrMessage = "InternalErrorException: "
	expect.ErrStack = errors.New("")
	domain := new(User)
	domain.MCognito = mockedCognito

	result := domain.PostSigninUser(payload)
	assert.Equal(t, *expect, result.ErrorHandling)
}
