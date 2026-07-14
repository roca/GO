package basic

import (
	"log"
	"os"
	"project01/proto/basic/protogen/basic"
	pb "project01/proto/basic/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// WriteToFile writes a proto message to a file
func WriteProtoToFile(msg proto.Message, fname string) {
	// Open a file for writing.
	out, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatal(err)
	}
}

func WriteProtoJSONToFile(msg proto.Message, fname string) {
	// Open a file for writing.
	out, err := protojson.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatal(err)
	}
}

// ReadFromFile reads a proto message from a file
func ReadProtoFromFile(fname string, dest proto.Message) {
	// Open a file for reading.
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Read the protobuf
	if err := proto.Unmarshal(in, dest); err != nil {
		log.Fatal(err)
	}

	// log.Println(dest)
}

func ReadProtoJSONFromFile(fname string, dest proto.Message) {
	// Open a file for reading.
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Read the protobuf
	if err := protojson.Unmarshal(in, dest); err != nil {
		log.Fatal(err)
	}

	// log.Println(dest)
}

func dummyUser() basic.User {
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

	return u
}

func WriteToFileSample() {
	u := dummyUser()
	WriteProtoToFile(&u, "superman_file.bin")
}

func ReadFromFileSample() {
	u := basic.User{}
	ReadProtoFromFile("superman_file.bin", &u)

	log.Println(&u)
}

func WriteToFileJSONSample() {
	u := dummyUser()
	WriteProtoJSONToFile(&u, "superman_file.json")
}

func ReadFromFileJSONSample() {
	u := basic.User{}
	ReadProtoJSONFromFile("superman_file.json", &u)

	log.Println(&u)
}

