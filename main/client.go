package main

import (
	"log"
	"os"

	pb "exercise/gRPC_test/gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the UserServiceImplement.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewToUpperClient(conn)

	// Contact the UserServiceImplement and print out its response.
	name := "jsm"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.QueryUserInfo(context.Background(), &pb.QueryUserInfoReq{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %v", r.UserInfo)
}
