package main

import (
	"fmt"
	"net/http"

	"k8s.io/client-go/kubernetes"
	"github.com/google/go-github/v47/github"
)

type server struct {
	client *kubernetes.Clientset
	githubClient *github.Client
}

func (s server) webhook(req http.ResponseWriter, w *http.Request) {
	payload, err := github.ValidatePayload(req, s.webhookSecretKey)
	if err != nil { ... }
	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil { ... }
	switch event := event.(type) {
	case *github.CommitCommentEvent:
		processCommitCommentEvent(event)
	case *github.CreateEvent:
		processCreateEvent(event)
	...
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
