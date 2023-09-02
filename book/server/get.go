package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	pb "github.com/julieqiu/scratch/book/proto"
)

func (s *Server) GetBook(ctx context.Context, in *pb.GetBookRequest) (*pb.Book, error) {
	log.Printf("GetBook was invoked with %v\n", in)

	if !strings.HasPrefix(in.Name, "books/") {
		return nil, fmt.Errorf("invalid book name: %q", in.Name)
	}

	// fetch the book...

	return &pb.Book{Name: in.Name}, nil
}
