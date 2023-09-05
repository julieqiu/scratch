package main

import (
	"context"
	"log"

	pb "github.com/julieqiu/scratch/book/proto"
)

func doGetBook(c pb.BookServiceClient) {
	log.Println("doGetBook was invoked")
	ctx := context.Background()
	res, err := c.GetBook(ctx, &pb.GetBookRequest{
		Name: "books/julie",
	})
	if err != nil {
		log.Fatalf("GetBook: %v\n", err)
	}
	log.Printf("Book: %s\n", res.Name)
}
