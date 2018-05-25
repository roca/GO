package viewmodels

import (

)

type Login struct {
	Title string
	Active string
}

func GetLogin() Login {
	result := Login{
		Title: "Lemonade Stand Society - Login",
		Active: "",
	}
	
	return result
}