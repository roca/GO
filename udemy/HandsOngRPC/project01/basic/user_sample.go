package basic

import (
	"log"
	hello "project01/hello/protogen"
)

func BasicUser() {
	u := hello.User{
		Id:       11,
		Username: "Superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
		Emails:   []string{"superman@movie.com", "superman@dc.com"},
		Gender:   hello.Gender_GENDER_MALE,
	}
	log.Println(&u)
}
