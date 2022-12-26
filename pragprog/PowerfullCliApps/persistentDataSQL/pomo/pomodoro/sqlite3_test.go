//go:build !inmemory
// +build !inmemory

package pomodoro_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/pomodoro"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	t.Log("Using SQLite3 repo")

	tf, err := ioutil.TempFile("", "pomo")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()

	dbRepo, err := repository.NewSQLite3Repo(tf.Name())
	if err != nil {
		t.Fatal(err)
	}

	return dbRepo, func() {
		os.Remove(tf.Name())
	}
}
