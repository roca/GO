package basic

import (
	"log"
	pb "project01/proto/basic/protogen"
)

func Basicpb() {
	h := pb.Hello{
		Name: "Clark Kent",
	}
	log.Println(&h)
}
