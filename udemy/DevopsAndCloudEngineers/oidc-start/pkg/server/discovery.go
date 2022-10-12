package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/oidc"
)

func (s *server) discovery(w http.ResponseWriter, r *http.Request) {
	discovery := oidc.Discovery{
		Issuer:                            s.Config.Url,
		AuthorizationEndpoint:             fmt.Sprintf("%s/authorization", s.Config.Url),
		TokenEndpoint:                     fmt.Sprintf("%s/token", s.Config.Url),
		UserinfoEndpoint:                  fmt.Sprintf("%s/userinfo", s.Config.Url),
		JwksURI:                           fmt.Sprintf("%s/jwks.json", s.Config.Url),
		ScopesSupported:                   []string{"openid"},
		ResponseTypesSupported:            []string{"code"},
		TokenEndpointAuthMethodsSupported: []string{"client_secret"},
		IDTokenSigningAlgValuesSupported:  []string{"RS256"},
		ClaimsSupported:                   []string{"iss", "sub", "aud", "exp", "nbf", "iat"},
		SubjectTypesSupported:             []string{"public"},
	}

	body, err := json.Marshal(discovery)
	if err != nil {
		returnError(w, http.StatusInternalServerError, fmt.Errorf("json.MarshalIndent error: %s", err))
		return
	}

	w.Write(body)

}
