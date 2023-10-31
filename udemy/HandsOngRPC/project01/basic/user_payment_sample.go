package basic

import (
	"log"
	"project01/proto/basic/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicReadUserPayment() {
	log.Println("Read User Payment")

	up := basic.UserPayment{}

	ReadProtoFromFile("basic_user_content_v1.bin", &up)

	log.Println(&up)

	jsonBytes, _ := protojson.Marshal(&up)
	log.Println(string(jsonBytes))
	log.Println()
}
