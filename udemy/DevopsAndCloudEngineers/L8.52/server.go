package main

import (
	"fmt"
	"net/http"

	"k8s.io/client-go/kubernetes"
)

type server struct {
	client *kubernetes.Clientset
}

func (s server) webhook(req http.ResponseWriter, w *http.Request) {
	fmt.Printf("test\n")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, s.webhookSecretKey)
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
