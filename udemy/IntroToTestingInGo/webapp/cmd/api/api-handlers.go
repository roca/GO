package main

import "net/http"

func (app *application) authenticate(w http.ResponseWriter, r *http.Request){
	// read a json payload

	// look up the user by email address

	// check the password

	// generate a tokens

	// send token to user
}

func (app *application) refresh(w http.ResponseWriter, r *http.Request){}

func (app *application) allUser(w http.ResponseWriter, r *http.Request){}

func (app *application) getUser(w http.ResponseWriter, r *http.Request){}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request){}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request){}

func (app *application) insertUser(w http.ResponseWriter, r *http.Request){}
