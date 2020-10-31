package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// see: https://docs.aws.amazon.com/code-samples/latest/catalog/go-cognito-CognitoListUsers.go.html
func NewCognitoSession()*cognitoidentityprovider.CognitoIdentityProvider {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return cognitoidentityprovider.New(sess)
}
func findSubs(userList *cognitoidentityprovider.ListUsersOutput) []string {
	var userSubs []string

	for _, user := range userList.Users {
		for _, attr := range user.Attributes {
			if *attr.Name == "sub" {
				userSubs = append(userSubs, *attr.Value)
			}
		}
	}
	return userSubs
}

func Handler() error {
	userPoolId := "ADD ME!!"

	userList, cognitoErr := NewCognitoSession().
		ListUsers(&cognitoidentityprovider.ListUsersInput{UserPoolId: &userPoolId})
	if cognitoErr != nil {
		return cognitoErr
	}

	userSubs := findSubs(userList)

	fmt.Println("userSubs: ", userSubs)

	return nil
}

func main() {
	lambda.Start(Handler)
}
