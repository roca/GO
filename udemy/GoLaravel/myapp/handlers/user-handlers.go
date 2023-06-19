package handlers

import (
	"fmt"
	"myapp/data"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := data.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "me@here.com",
		Active:    1,
		Password:  "password",
	}

	id, err := h.Models.Users.Insert(u)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return
	}

	fmt.Fprintf(w, "User %s created with id: %d", u.FirstName, id)

}

func (h *Handlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	u, err := h.Models.Users.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return
	}

	fmt.Fprintf(w, "User %s %s %s", u.FirstName, u.LastName, u.Email)
}

func (h *Handlers) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	u, err := h.Models.Users.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return
	}

	u.LastName = h.App.RandomString(10)

	err = h.Models.Users.Update(*u)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return
	}

	fmt.Fprintf(w, "Users %s last name updated to %s", u.FirstName, u.LastName)
}
