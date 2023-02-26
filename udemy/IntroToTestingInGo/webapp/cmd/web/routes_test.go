package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

// Test Application routes
func Test_application_routes(t *testing.T) {
	var registered = []struct {
		name   string
		route  string
		method string
	}{
		{"Home", "/", "GET"},
		{"Login", "/login", "POST"},
		{"static", "/static/*", "GET"},
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
