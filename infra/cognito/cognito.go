package cognito

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/midnight-trigger/todo/api/definition"
)

type Cognito struct {
}

//go:generate mockgen -source cognito.go -destination mock_cognito/mock_cognito.go
type ICognito interface {
	SignUp(body *definition.PostSignupUserRequestBody) (response *cognitoidentityprovider.SignUpOutput, err error)
	AdminInitiateAuth(body *definition.PostSigninUserRequestBody) (response *cognitoidentityprovider.AdminInitiateAuthOutput, err error)
}

func GetNewCognito() *Cognito {
	return &Cognito{}
}

// Cognito会員登録
func (m *Cognito) SignUp(body *definition.PostSignupUserRequestBody) (response *cognitoidentityprovider.SignUpOutput, err error) {
	svc := cognitoidentityprovider.New(
		session.New(),
		&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))},
	)

	params := new(cognitoidentityprovider.SignUpInput)
	params.ClientId = aws.String(os.Getenv("APPLICATION_CLIENT"))
	params.Password = aws.String(body.Password)
	params.Username = aws.String(body.Email)
	response, err = svc.SignUp(params)
	return
}

// Cognitoログイン
func (m *Cognito) AdminInitiateAuth(body *definition.PostSigninUserRequestBody) (response *cognitoidentityprovider.AdminInitiateAuthOutput, err error) {
	svc := cognitoidentityprovider.New(
		session.New(),
		&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))},
	)

	params := new(cognitoidentityprovider.AdminInitiateAuthInput)
	params.AuthFlow = aws.String("ADMIN_NO_SRP_AUTH")
	params.AuthParameters = map[string]*string{
		"USERNAME": aws.String(body.Email),
		"PASSWORD": aws.String(body.Password),
	}
	params.ClientId = aws.String(os.Getenv("APPLICATION_CLIENT"))
	params.UserPoolId = aws.String(os.Getenv("USER_POOL_ID"))
	response, err = svc.AdminInitiateAuth(params)
	return
}
