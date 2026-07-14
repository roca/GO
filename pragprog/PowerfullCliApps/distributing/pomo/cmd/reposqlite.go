//go:build !inmemory && !containers
// +build !inmemory,!containers

package cmd

import (
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/distributing/pomo/pomodoro"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/distributing/pomo/pomodoro/repository"
	"github.com/spf13/viper"
)

func getRepo() (pomodoro.Repository, error) {
	repo, err := repository.NewSQLite3Repo(viper.GetString("db"))
	if err != nil {
		return nil, err
	}
	return repo, nil
}
