package basic

import (
	"log"
	pb "project01/proto/basic/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func BasicUser() {

	addr := &pb.Address{
		Street:     "Daily Planet",
		City:       "Metropolis",
		Country:    "US",
		Coordinate: &pb.Address_Coordinate{Latitude: 40.70797893425118, Longitude: -74.01163838107261},
	}

	a := &anypb.Any{}
	paperMail := &pb.PaperMail{PaperMailAddress: "Daily Planet, Metropolis, US"}
	err := anypb.MarshalFrom(a, paperMail, proto.MarshalOptions{})
	if err != nil {
		log.Fatal(err)
	}

	socialMedia := &pb.SocialMedia{
		SocialMediaPlatform: "Twitter",
		SocialMediaUsername: "batman",
	}

	skills := map[string]uint32{
		"Go":     9,
		"Python": 8,
		"Java":   6,
		"Kotlin": 7,
		"Ruby":   9,
	}
	

	u := pb.User{
		Id:                   11,
		Username:             "Superman",
		IsActive:             true,
		Password:             []byte("supermanpassword"),
		Emails:               []string{"superman@movie.com", "superman@dc.com"},
		Gender:               pb.Gender_GENDER_MALE,
		Address:              addr,
		CommunicationChannel: a,
		ElectronicCommChannel: &pb.User_SocialMedia{
			SocialMedia: socialMedia,
		},
		SkillRating: skills,
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

func BasicUnMarshalAynKnown() {
	socialMedia := &pb.SocialMedia{
		SocialMediaPlatform: "Twitter",
		SocialMediaUsername: "batman",
	}

	var a anypb.Any
	err := anypb.MarshalFrom(&a, socialMedia, proto.MarshalOptions{})
	if err != nil {
		log.Fatal(err)
	}

	sm := &pb.SocialMedia{}

	if err := a.UnmarshalTo(sm); err != nil {
		log.Fatal(err)
	}

	jsonBytes, _ := protojson.Marshal(sm)
	log.Println(string(jsonBytes))

	if a.MessageIs(&pb.SocialMedia{}) {
		unmarsheled, err := a.UnmarshalNew()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Unmarshled as:", unmarsheled.ProtoReflect().Descriptor().FullName())
		log.Println("It's a SocialMedia type")
	}
}
