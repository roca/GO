package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"udemy.com/gRPC/code/blog/blogpb"
)

func main() {

	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("count no connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	doUnary(c)
}

func doUnary(c blogpb.BlogServiceClient) {
	fmt.Println("Creating  the Blog")
	blog := &blogpb.Blog{
		AuthorId: "Stephane",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	req := &blogpb.CreateBlogRequets{
		Blog: blog,
	}

	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Blog RPC: %v", err)
	}
	log.Printf("Blog has been created: %v", res.Blog)
}
