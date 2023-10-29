package basic

import (
	"log"
	"project01/proto/basic/protogen/basic"
)

func WriteBasicUserContentSample() {
	uc := basic.UserContent{
		UserContentId: 1,
		Slug:          "slug",
		Title:         "title",
		HtmlContent:   "content",
		AuthorId:      1,
	}
	WriteProtoToFile(&uc, "basic_user_content.bin")
}

func ReadBasicUserContentSample() {
	uc := basic.UserContent{}
	ReadProtoFromFile("basic_user_content.bin", &uc)

	log.Println(&uc)
}
