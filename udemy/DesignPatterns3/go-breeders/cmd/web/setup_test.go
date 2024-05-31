package main

import (
	"go-breaders/configuration"
	"go-breaders/models"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {

	testBackend := &TestBackend{}
	testAdapter := &RemoteService{Remote: testBackend}

	testApp = application{
		App:        configuration.New(nil),
		catService: testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	return []*models.CatBreed{
		&models.CatBreed{
			ID: 1,
			BreadProps: models.BreadProps{
				Breed:   "Tomcat",
				Details: "Tomcat is a breed of cat",
			},
		},
	}, nil
}
