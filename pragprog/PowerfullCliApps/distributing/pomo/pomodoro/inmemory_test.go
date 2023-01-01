//go:build inmemory
// +build inmemory

package pomodoro_test

import (
	"testing"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/distributing/pomo/pomodoro"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/distributing/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	t.Log("Using in-memory repo")

	return repository.NewInMemoryRepo(), func() {}
}
