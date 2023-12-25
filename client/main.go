package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"unary_example/unary_grpc_example/server/protofiles/greetpb"
)

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	getGreeting("Jack", "us", c)
	getGreeting("Jose", "mx", c)
}

func getGreeting(name, countryCode string, c greetpb.GreetServiceClient) {
	log.Println("creating greeting")

	res, err := c.Greet(context.Background(), &greetpb.GreetRequest{
		CountryCode: countryCode,
		UserName:    name,
	})

	if err != nil {
		log.Println("error", err)
		panic(err)
	}
	log.Println(res.Result)
}
