package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-github/v47/github"
	"k8s.io/client-go/kubernetes"
)

type server struct {
	client           *kubernetes.Clientset
	githubClient     *github.Client
	webhookSecretKey string
}

func (s server) webhook(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
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
		files := getFiles(event.Commits)
		fmt.Printf("Files: %s\n", strings.Join(files, ", "))
		for _, filename := range files {
			if strings.HasSuffix(filename, ".yaml") {
				downLoadedFile, _, err := s.githubClient.Repositories.DownloadContents(
					ctx,
					*event.Repo.Owner.Name,
					*event.Repo.Name,
					filename,
					&github.RepositoryContentGetOptions{
						Ref: "staging",
					},
				)
				if err != nil {
					w.WriteHeader(500)
					fmt.Printf("DownloadContents error: %v\n", err)
					return
				}
				defer downLoadedFile.Close()
				fileBody, err := io.ReadAll(downLoadedFile)
				if err != nil {
					w.WriteHeader(500)
					fmt.Printf("ReadAll error: %v\n", err)
					return
				}
				_, _, err = deploy(ctx, s.client, fileBody)
				if err != nil {
					w.WriteHeader(500)
					fmt.Printf("Deployment error: %v\n", err)
					return
				}
				fmt.Printf("Deployment of %s finished\n body: \n %s", filename, string(fileBody))
			}
		}
	default:
		w.WriteHeader(500)
		fmt.Printf("Event not found %s\n", event)
		return
	}
}

func getFiles(commits []*github.HeadCommit) []string {
	allFiles := []string{}
	for _, commit := range commits {
		allFiles = append(allFiles, commit.Added...)
		allFiles = append(allFiles, commit.Modified...)
	}
	allUniqueFilesMap := make(map[string]bool)
	for _, file := range allFiles {
		allUniqueFilesMap[file] = true
	}
	allUniqueFilesSlice := []string{}
	for filename := range allUniqueFilesMap {
		allUniqueFilesSlice = append(allUniqueFilesSlice, filename)
	}
	return allUniqueFilesSlice
}
