package basic

import (
	"log"
	"project01/proto/basic/protogen/basic"
)

func BasicUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 1,
		Slug:          "slug-v1",
		Title:         "title",
		HtmlContent:   "content",
		AuthorId:      1,
		//Category: "Fiction",
		//SubCategory: "Science Fiction",
	}
	WriteProtoToFile(&uc, "basic_user_content_v1.bin")
}

func BasicReadUserContentV1() {
	uc := basic.UserContent{}
	ReadProtoFromFile("basic_user_content_v1.bin", &uc)

	log.Println(&uc)
}
