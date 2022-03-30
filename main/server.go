package main

import (
	pb "exercise/gRPC_test/gen"
	"exercise/gRPC_test/handler"
	"exercise/gRPC_test/model"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strings"
)

const (
	port = ":50051"
)

type UserServiceImplement struct{}

func (s *UserServiceImplement) Upper(ctx context.Context, in *pb.UpperRequest) (*pb.UpperResponse, error) {
	log.Printf("Received: %s", in.Name)
	return &pb.UpperResponse{Message: strings.ToUpper(in.Name)}, nil
}

func (s *UserServiceImplement) QueryUserInfo(ctx context.Context, req *pb.QueryUserInfoReq) (*pb.QueryUserInfoResp, error) {
	log.Printf("Received: %s", req.Name)
	return handler.NewUserInfoHandler().QueryUserInfo(ctx, req)
}

func main() {
	// init sql
	initSQL()
	// init log
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	// init UserServiceImplement
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterToUpperServer(s, &UserServiceImplement{})
	// Register reflection service on gRPC UserServiceImplement.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSQL() {
	err := model.InitDatabase()
	if err != nil {
		fmt.Fprintf(os.Stderr, "init mysql error， err=[%v]\n", err)
		os.Exit(-1)
	}
}
