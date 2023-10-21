package basic

import (
	"log"
	hello "project01/hello/protogen"
)

func BasicUser() {
	u := hello.User{
		Username: "Superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
		Id:       11,
	}
	log.Println(&u)
}
