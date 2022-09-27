package server

import (
	"fmt"
	"net/http"
)

func (s *server) token(w http.ResponseWriter, r *http.Request) {

	var code string

	if code = r.URL.Query().Get("code"); code == "" {
		returnError(w, http.StatusBadRequest, fmt.Errorf("code is missing"))
		return
	}

	fmt.Println(s.Codes[code])

}
