package main

import (
	"log"
	"net"

	pb "github.com/amanjots19/grpc-search/proto"
	"google.golang.org/grpc"
)

func main() {
	dic := newDIContainer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransacUserServer(s, dic.server())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
