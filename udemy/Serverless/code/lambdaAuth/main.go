package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
)

// Help function to generate an IAM policy
func generatePolicy(principalId, effect, resource string, claims jwt.MapClaims) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"sub":  claims["sub"],
		"name": claims["name"],
		"data": claims["data"],
	}
	return authResponse
}

func ExtractClaims(jwtStr string) (jwt.MapClaims, error) {
	key := "your-256-bit-secret"

	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return jwt.MapClaims{}, errors.New("Claims could not be extracted")
	}

}

func handler(event *events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {

	tokenString := event.AuthorizationToken
	log.Printf("AuthorizationToken: %s", tokenString)

	claims, err := ExtractClaims(tokenString)
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, err
	}

	log.Println(claims)

	if claims["sub"] != "user1" {
		return generatePolicy("user1", "Deny", event.MethodArn, claims), nil
	}

	return generatePolicy("user1", "Allow", event.MethodArn, claims), nil

}

func main() {
	lambda.Start(handler)

}
