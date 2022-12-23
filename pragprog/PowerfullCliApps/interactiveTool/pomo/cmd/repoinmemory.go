package cmd

import (
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/interactiveTool/pomo/pomodoro"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/interactiveTool/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
