//go:build inmemory || containers
// +build inmemory containers

package cmd

import (
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/distributing/pomo/pomodoro"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/distributing/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
