package main

/*
   Route: GET /auth
*/

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/dgrijalva/jwt-go"
	"udemy.com/sls/sls-notes-backend/api/utils"
)

type CognitoIdentity struct {
	CognitoData *cognitoidentity.GetCredentialsForIdentityOutput `json:"cognito_data"`
	UserName    interface{}                                      `json:"user_name"`
}

func JwtDecode(jwtStr string) (*jwt.Token, error) {
	key := "your-256-bit-secret"

	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return &jwt.Token{}, err
	}

	return token, nil
}

func handler(event *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	idToken := event.Headers["Authorization"]
	cognitoIdentity := cognitoidentity.CognitoIdentity{}
	identityPoolID := os.Getenv("COGNITO_IDENTITY_POOL_ID")

	idData, err := cognitoIdentity.GetId(&cognitoidentity.GetIdInput{
		IdentityPoolId: aws.String(identityPoolID),
		Logins: map[string]*string{
			"accounts.google.com": aws.String(idToken),
		},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	cognitoData, err := cognitoIdentity.GetCredentialsForIdentity(&cognitoidentity.GetCredentialsForIdentityInput{
		IdentityId: idData.IdentityId,
		Logins: map[string]*string{
			"accounts.google.com": aws.String(idToken),
		},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	decoded, err := JwtDecode(idToken)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	cognitoResponse := CognitoIdentity{
		CognitoData: cognitoData,
		UserName:    decoded.Header["name"],
	}
	b, err := json.Marshal(&cognitoResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    utils.GetResponseHeaders(),
		Body:       string(b),
	}

	return response, nil
}

func main() {
	lambda.Start(handler)

}
