package main

import (
	"testing"

	"github.com/google/go-github/v47/github"
)

func TestGetFiles(t *testing.T) {
	commits := []*github.HeadCommit{}
	c1 := github.HeadCommit{
		Added:    []string{"file1", "file2"},
		Modified: []string{},
	}
	headCommits := append(commits, &c1)

	files := getFiles(headCommits)

	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
	if files[0] != "file1" {
		t.Errorf("Expected file1, got %s", files[0])
	}
}
