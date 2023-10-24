package basic

import (
	"log"
	hello "project01/hello/protogen"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	batman := hello.User{
		Id:       97,
		Username: "batman",
		IsActive: true,
		Password: []byte("batmanpassword"),
		Gender:   hello.Gender_GENDER_MALE,
	}

	nightwing := hello.User{
		Id:       96,
		Username: "nightwing",
		IsActive: true,
		Password: []byte("nightwingpassword"),
		Gender:   hello.Gender_GENDER_MALE,
	}

	robin := hello.User{
		Id:       96,
		Username: "robin",
		IsActive: true,
		Password: []byte("robinpassword"),
		Gender:   hello.Gender_GENDER_MALE,
	}

	batFamily := hello.UserGroup{
		GroupId:     999,
		GroupName:   "Bat Family",
		Users:       []*hello.User{&batman, &nightwing, &robin},
		Description: "The classic bat family",
	}

	jsonBytes, _ := protojson.Marshal(&batFamily)
	log.Println(string(jsonBytes))
}
