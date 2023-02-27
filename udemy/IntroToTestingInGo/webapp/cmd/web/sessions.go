package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

func getSession() *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}
