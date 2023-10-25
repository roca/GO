package basic

import (
	"log"
	pb "project01/proto/basic/protogen"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	batman := pb.User{
		Id:       97,
		Username: "batman",
		IsActive: true,
		Password: []byte("batmanpassword"),
		Gender:   pb.Gender_GENDER_MALE,
	}

	nightwing := pb.User{
		Id:       96,
		Username: "nightwing",
		IsActive: true,
		Password: []byte("nightwingpassword"),
		Gender:   pb.Gender_GENDER_MALE,
	}

	robin := pb.User{
		Id:       96,
		Username: "robin",
		IsActive: true,
		Password: []byte("robinpassword"),
		Gender:   pb.Gender_GENDER_MALE,
	}

	batFamily := pb.UserGroup{
		GroupId:     999,
		GroupName:   "Bat Family",
		Users:       []*pb.User{&batman, &nightwing, &robin},
		Description: "The classic bat family",
	}

	jsonBytes, _ := protojson.Marshal(&batFamily)
	log.Println(string(jsonBytes))
}
