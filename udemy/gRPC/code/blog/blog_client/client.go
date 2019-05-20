package main

import (
	"context"
	"fmt"
	"io"
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

	req := &blogpb.CreateBlogRequest{
		Blog: blog,
	}

	createBlogRes, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Blog RPC: %v", err)
	}
	log.Printf("Blog has been created: %v", createBlogRes.Blog)
	blogID := createBlogRes.GetBlog().GetId()

	// read Blog
	fmt.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "5cdb08200b9d068dcc9b3eff",
	})
	if err2 != nil {
		log.Printf("error while calling ReadBlog RPC: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, err := c.ReadBlog(context.Background(), readBlogReq)
	if err != nil {
		log.Printf("error while calling ReadBlog RPC: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", readBlogRes)

	// update Blog

	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awsome additions",
	}

	updateBlogReq := &blogpb.UpdateBlogRequest{Blog: newBlog}
	updateBlogRes, err := c.UpdateBlog(context.Background(), updateBlogReq)
	if err != nil {
		log.Printf("error while calling UpdateBlog RPC: %v\n", err)
	}

	log.Printf("Blog was updated: %v\n", updateBlogRes)

	// Delete blog
	deleteBlogRes, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId: blogID,
	})
	if err != nil {
		log.Printf("error while calling DeleteBlog RPC: %v\n", err)
	}

	log.Printf("Blog %v was delete: \n", deleteBlogRes.GetBlogId())

	// List Blogs

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from listBlog: %v", res.GetBlog())
	}

}
