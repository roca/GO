package basic

import (
	"log"
	pb "project01/proto/basic/protogen/basic"
)

func Basicpb() {
	h := pb.Hello{
		Name: "Clark Kent",
	}
	log.Println(&h)
}
