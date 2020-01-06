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
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/dgrijalva/jwt-go"
	"udemy.com/sls/sls-notes-backend/api/utils"
)

type CognitoResponse struct {
	CognitoData *cognitoidentity.GetCredentialsForIdentityOutput `json:"cognito_data"`
	UserName    interface{}                                      `json:"user_name"`
}

var sess *session.Session
var svc *cognitoidentity.CognitoIdentity

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = cognitoidentity.New(sess)
}

func JwtDecode(jwtStr string) (*jwt.Token, error) {

	token, _, err := new(jwt.Parser).ParseUnverified(jwtStr, jwt.MapClaims{})
	if err != nil {
		return &jwt.Token{}, err
	}

	return token, nil
}

func handler(event *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	idToken := event.Headers["Authorization"]
	identityPoolID := os.Getenv("COGNITO_IDENTITY_POOL_ID")

	idData, err := svc.GetId(&cognitoidentity.GetIdInput{
		IdentityPoolId: aws.String(identityPoolID),
		Logins: map[string]*string{
			"accounts.google.com": aws.String(idToken),
		},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	cognitoData, err := svc.GetCredentialsForIdentity(&cognitoidentity.GetCredentialsForIdentityInput{
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

    claims, _ := decoded.Claims.(jwt.MapClaims)

	cognitoResponse := CognitoResponse{
		CognitoData: cognitoData,
		UserName:    claims["name"],
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
