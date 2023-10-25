package basic

import (
	"log"
	pb "project01/proto/basic/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {

	addr := &pb.Address{
		Street: "Daily Planet",
		City: "Metropolis",
		Country: "US",
		Coordinate: &pb.Address_Coordinate{
			Latitude: 40.70797893425118,
			Longitude: -74.01163838107261,
		},
	}

	u := pb.User{
		Id:       11,
		Username: "Superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
		Emails:   []string{"superman@movie.com", "superman@dc.com"},
		Gender:   pb.Gender_GENDER_MALE,
		Address: addr,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	u := pb.User{
		Id:       99,
		Username: "wonderwoman",
		IsActive: true,
		Password: []byte("wonderwomanpassword"),
		Emails:   []string{"wonderwoman@movie.com", "wonderwoman@dc.com"},
		Gender:   pb.Gender_GENDER_FEMALE,
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

	var p pb.User

	err := protojson.Unmarshal(jsonBytes, &p)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(&p)

}
