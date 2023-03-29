package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

// Test Application routes
func Test_app_routes(t *testing.T) {
	var registered = []struct {
		name   string
		route  string
		method string
	}{
		{"Auth", "/auth", "POST"},
		{"refresh token", "/refresh", "POST"},
		// User routes
		{"Get all users", "/users/", "GET"},
		{"Get a user", "/users/{userID}", "GET"},
		{"Delete a user", "/users/{userID}", "DELETE"},
		{"Add a user", "/users/", "PUT"},
		{"Update a user", "/users/", "PATCH"},
	}

	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, r := range registered {
		t.Run(r.name, func(t *testing.T) {
			if !routeExists(r.route, r.method, chiRoutes) {
				t.Errorf("Route '%s' with method '%s' does not exist", r.route, r.method)
			}
		})
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})

	return found
}
