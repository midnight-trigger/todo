package cognito

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/midnight-trigger/todo/api/definition"
)

func SignUp(body *definition.PostUserRequestBody) (cognitoResponse *cognitoidentityprovider.SignUpOutput, err error) {
	svc := cognitoidentityprovider.New(
		session.New(),
		&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))},
	)

	params := new(cognitoidentityprovider.SignUpInput)
	params.ClientId = aws.String(os.Getenv("APPLICATION_CLIENT"))
	params.Password = aws.String(body.Password)
	params.Username = aws.String(body.Email)
	cognitoResponse, err = svc.SignUp(params)

	return
}
