package basic

import (
	"log"
	hello "project01/hello/protogen"

	"google.golang.org/protobuf/encoding/protojson"
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

func ProtoToJsonUser() {
	u := hello.User{
		Id:       99,
		Username: "wonderwoman",
		IsActive: true,
		Password: []byte("wonderwomanpassword"),
		Emails:   []string{"wonderwoman@movie.com", "wonderwoman@dc.com"},
		Gender:   hello.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	jsonBytes := []byte(`{
		"id": 97,
		"username": "batman",
		"is_active": true,
		"password": "batmannpassword",
		"emails": [
			"batman@movie.com",
			"batman@dc.com"
		],
		"gender": "GENDER_MALE"
	}`)

	var p hello.User

	err := protojson.Unmarshal(jsonBytes, &p)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(&p)

}
