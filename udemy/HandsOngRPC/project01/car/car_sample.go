package car

import (
	"log"
	pb "project01/proto/car/protogen/car"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

func ValidateCar() {
	car := &pb.Car{
		CarId:           uuid.New().String(),
		Brand:           "Toyota",
		Model:           "Corolla",
		Price:           10000,
		ManufactureYear: 2021,
	}

	// Validate the car
	if err := car.ValidateAll(); err != nil {
		log.Fatalln("Validation failed", err)
	}

	log.Println(car)

	jsonBytes, _ := protojson.Marshal(car)
	log.Println(string(jsonBytes))
}
