package proverbial

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jboursiquot/go-proverbs"
)

// Handler handles incoming request.
type Handler struct {
}

// NewHandler returns a new handler.
func NewHandler() *Handler {
	return &Handler{}
}

// Handle handles the request.
func (h *Handler) Handle(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	p := proverbs.Random()
	body, err := json.Marshal(p)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	headers := map[string]string{"Content-Type": "application/json"}
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       string(body),
		StatusCode: http.StatusOK,
	}, nil
}
