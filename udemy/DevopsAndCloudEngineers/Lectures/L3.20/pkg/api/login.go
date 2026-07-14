package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type LoginRequest struct {
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

func doLoginRequest(client ClientInterface, requestURL, password string) (string, error) {
	loginRequest := LoginRequest{
		Password: password,
	}

	body, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("Marshal error %s", err)
	}

	resp, err := client.Post(requestURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("http Post error: %s\n", err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ReadAll error: %s\n", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Invalid output (HTTP Code %d): %s\n", resp.StatusCode, respBody)
	}

	if !json.Valid(respBody) {
		return "", RequestError{
			HTTPCode: resp.StatusCode,
			Body:     string(respBody),
			Err:      fmt.Sprintf("No valid JSON rerturned"),
		}
	}

	var loginResponse LoginResponse
	if err := json.Unmarshal(respBody, &loginResponse); err != nil {
		return "", RequestError{
			HTTPCode: resp.StatusCode,
			Body:     string(respBody),
			Err:      fmt.Sprintf("Page unmarshal error: %s\n", err),
		}
	}

	if loginResponse.Token == "" {
		return "", RequestError{
			HTTPCode: resp.StatusCode,
			Body:     string(respBody),
			Err:      "Empty token replied",
		}
	}

	return loginResponse.Token, nil
}
