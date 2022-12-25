//go:build !inmemory
// +build !inmemory

package repository

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/pomodoro"
)

const (
	createIntervalTable string = ` CREATE TABLE IF NOT EXISTS "interval" (
		"id" INTEGER,
		"start_time" DATETIME NOT NULL,
		"planned_duration" INTEGER DEFAULT 0,
		"actual_duration" INTEGER DEFAULT 0,
		"category" TEXT NOT NULL,
		"state" INTEGER DEFAULT 1,
		PRIMARY KEY("id")
	);`
)

type dbRepo struct {
	sync.RWMutex
	db *sql.DB
}

func NewSQLite3MemoryRepo(dbfile string) (*dbRepo, error) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := db.Exec(createIntervalTable); err != nil {
		return nil, err
	}

	return &dbRepo{
		db: db,
	},nil
}

func (r *dbRepo) Create(i pomodoro.Interval) (int64, error) {
	r.Lock()
	defer r.Unlock()

	return 0, nil
}

func (r *dbRepo) Update(i pomodoro.Interval) error {
	r.Lock()
	defer r.Unlock()

	return nil
}

func (r *dbRepo) ByID(id int64) (pomodoro.Interval, error) {
	r.RLock()
	defer r.RUnlock()

	i := pomodoro.Interval{}

	return i, nil
}

func (r *dbRepo) Last() (pomodoro.Interval, error) {
	r.RLock()
	defer r.RUnlock()

	i := pomodoro.Interval{}

	return i, nil
}

func (r *dbRepo) Breaks(n int) ([]pomodoro.Interval, error) {
	r.RLock()
	defer r.RUnlock()

	data := []pomodoro.Interval{}

	return data, nil
}
