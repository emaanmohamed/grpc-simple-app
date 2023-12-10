package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/emaanmohamed/grpc-go/server/protofiles/greetpb"
	"google.golang.org/grpc"
)

// this struct could be used to dependency injection
type server struct{}

// handler
func (s server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	log.Println("User name: ", req.UserName)
	log.Println("Country code: ", req.CountryCode)

	var greeting string

	switch req.CountryCode {
	case "us":
		greeting = "hello " + req.UserName
	case "mx":
		greeting = "Hola " + req.UserName
	default:
		greeting = "Hola/Hello " + req.UserName
	}

	return &greetpb.GreetResponse{
		Result: greeting,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8060")

	if err != nil {
		panic(err)
	}

	fmt.Println("starting server")

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, server{})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
