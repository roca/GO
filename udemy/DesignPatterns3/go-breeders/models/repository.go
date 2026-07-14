package models

import (
	"database/sql"
)

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
	GetDogOfMonthByID(id int) (*DogOfMonth, error)
}

type mysqlRepository struct {
	DB *sql.DB
}

func newMysqlRepository(db *sql.DB) Repository {
	return &mysqlRepository{DB: db}
}

type testRepository struct {
	DB *sql.DB
}

func newTestRepository(db *sql.DB) Repository {
	return &testRepository{DB: nil}
}
