package proverbial_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jboursiquot/proverbial"
)

func TestHandler(t *testing.T) {
	tests := map[string]struct {
		req events.APIGatewayProxyRequest
		res events.APIGatewayProxyResponse
	}{
		"random": {
			req: events.APIGatewayProxyRequest{},
			res: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			h := proverbial.NewHandler()

			r, err := h.Handle(context.Background(), &tc.req)
			if err != nil {
				t.Fatalf("Got an err: %s", err)
			}

			if r.StatusCode != http.StatusOK {
				t.Fatalf("Expected 200 OK but got %d", r.StatusCode)
			}
		})
	}
}
