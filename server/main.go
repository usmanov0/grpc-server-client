package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"unary_example/unary_grpc_example/server/protofiles/greetpb"
)

type server struct{}

func (s server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("User name: ", req.UserName)
	log.Println("Country: ", req.CountryCode)

	var greeting string
	switch req.CountryCode {
	case "us":
		greeting = "Hello" + req.UserName
	case "uz":
		greeting = "Salom" + req.UserName
	default:
		greeting = "Hello/Salom" + req.UserName
	}
	return &greetpb.GreetResponse{
		Result: greeting,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("starting server...")

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, server{})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
