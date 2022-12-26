//go:build inmemory
// +build inmemory

package pomodoro_test

import (
	"testing"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/pomodoro"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	t.Log("Using in-memory repo")

	return repository.NewInMemoryRepo(), func() {}
}
