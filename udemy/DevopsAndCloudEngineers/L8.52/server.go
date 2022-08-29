package main

import (
	"fmt"
	"net/http"

	"github.com/google/go-github/v47/github"
	"k8s.io/client-go/kubernetes"
)

type server struct {
	client           *kubernetes.Clientset
	githubClient     *github.Client
	webhookSecretKey string
}

func (s server) webhook(w http.ResponseWriter, req *http.Request) {
	payload, err := github.ValidatePayload(req, []byte(s.webhookSecretKey))
	if err != nil {
		w.WriteHeader(500)
		fmt.Printf("ValidatePayload error: %v\n", err)
		return
	}
	event, err := github.ParseWebHook(github.WebHookType(req), payload)
	if err != nil {
		w.WriteHeader(500)
		fmt.Printf("ParseWebHook error: %v\n", err)
		return
	}
	switch event := event.(type) {
	case *github.PushEvent:

	default:
		w.WriteHeader(500)
		fmt.Printf("Event not found %s\n", event)
		return
	}
}
