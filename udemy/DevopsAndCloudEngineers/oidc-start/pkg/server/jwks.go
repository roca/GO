package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/oidc"
)

func (s *server) jwks(w http.ResponseWriter, r *http.Request) {

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(s.PrivateKey)
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("Private key parsing error: %s", err))
		return
	}
	publicKey := privateKey.PublicKey

	jwks := oidc.Jwks{
		Keys: []oidc.JwksKey{
			{
				Kid: "0-0-0-1",
				Alg: "RS256",
				Kty: "RSA",
				Use: "sig",
				N:   base64.StdEncoding.EncodeToString(publicKey.N.Bytes()),
				E:   "AQAB",
			},
		},
	}

	out , err := json.Marshal(jwks)
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("jwks marshall error: %s", err))
		return
	}

	w.Write(out)

}
