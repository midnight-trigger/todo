package cognito

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/midnight-trigger/todo/api/definition"
)

func SignUp(body *definition.PostUserRequestBody) (response *cognitoidentityprovider.SignUpOutput, err error) {
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

func AdminInitiateAuth(body *definition.PostSigninUserRequestBody) (response *cognitoidentityprovider.AdminInitiateAuthOutput, err error) {
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
