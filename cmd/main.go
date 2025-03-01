package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	server := grpc.NewServer()

	lis, err := net.Listen("tcp", ":55053")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	fmt.Println(lis, "Connection started at the endpoint 55053")

	if err := server.Serve(lis); err != nil {
		log.Fatalln("failed to serv", err)
	}

}
