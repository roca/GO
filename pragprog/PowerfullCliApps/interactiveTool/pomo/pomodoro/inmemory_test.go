package pomodoro_test

import (
	"testing"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/pomodoro"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
