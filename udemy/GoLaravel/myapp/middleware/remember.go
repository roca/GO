package middleware

import (
	"fmt"

	"myapp/data"
	"net/http"
)

func (m *Middleware) CheckRemember(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.App.Session.Exists(r.Context(), "userID") {
			// user is not logged in
			cookie, err := r.Cookie(fmt.Sprintf("_%s_remember", m.App.AppName))
			if err != nil {
				// no cookie, so on to the next middleware
				next.ServeHTTP(w, r)
			} else {
				// we found a cookie, so check it
				key := cookie.Value
				var u data.User
				if len(key) > 0 {
					// cookie has some data, so validate it
				}

			}

		} else {
			// user is logged in
			next.ServeHTTP(w, r)
		}
	})
}
