//go:build inmemory
// +build inmemory

package cmd

import (
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/persistentDataSQL/pomo/pomodoro"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/persistentDataSQL/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
