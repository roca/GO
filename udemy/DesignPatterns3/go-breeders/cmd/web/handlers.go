package main

import (
	"fmt"
	"go-breaders/models"
	"go-breaders/pets"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/roca/go-toolkit/v2"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) DogOfMonth(w http.ResponseWriter, r *http.Request) {
	// Get the breed
	breed, _ := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")

	// get the dog of the month from the database
	dom, _ := app.App.Models.Dog.GetDogOfMonthByID(1)

	timeLayout := "2006-01-02"
	dob, _ := time.Parse(timeLayout, "2023-11-01")

	// Create dog and decorate it
	dog := &models.DogOfMonth{
		Dog: &models.Dog{
			ID: 1,
			PetProps: models.PetProps{
				Name:             "Sam",
				BreedID:          breed.ID,
				Color:            "Black & Tan",
				DataOfBirth:      dob,
				SpayedOrNeutered: 0,
				Description:      "A very friendly dog who loves to play fetch.",
				Weight:           75,
			},
			Breed: *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}

	// Serve the web page
	data := map[string]interface{}{
		"dog": dog,
	}

	app.render(w, "dog-of-month.page.gohtml", &templateData{ Data: data})
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.New("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.New("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools

	// create a dog with the builder pattern
	dog, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("A mixed breed of unknown origin. Probably has some German Shepherd heritage.").
		SetColor("Black and White").
		SetAge(3).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools

	// create a cat with the builder pattern
	cat, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("mixed breed").
		SetWeight(10).
		SetDescription("A mixed breed of unknown origin. Probably has some Siamese heritage.").
		SetColor("White").
		SetAge(2).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	catBreeds, err := app.App.CatService.CallAllBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools

	// Get species from URL itself
	species := chi.URLParam(r, "species")

	// Get breed from the URL.
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)

	// Create a pet from the abstract factory
	pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Write the result as JSON.
	_ = t.WriteJSON(w, http.StatusOK, pet)
}
