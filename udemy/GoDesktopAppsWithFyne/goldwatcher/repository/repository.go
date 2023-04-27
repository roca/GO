package repository

import (
	"database/sql"
	"errors"
)

var (
	errUpdateFailed = errors.New("update failed")
	errDeleteFailed = errors.New("delete failed")
)

type DatabaseRepo interface {
	Connection() *sql.DB
}
