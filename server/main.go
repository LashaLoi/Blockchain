package main

import (
	"context"
	"log"
	"net"

	"../proto"
	"google.golang.org/grpc"
)

// GRPCServer ...
type GRPCServer struct{}

// AddBlock ...
func (s *GRPCServer) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	return new(proto.AddBlockResponse), nil
}

// GetBlockchain ...
func (s *GRPCServer) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponce, error) {
	return new(proto.GetBlockchainResponce), nil
}

func main() {
	port := ":8080"

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &GRPCServer{})

	log.Printf("tcp server was started on: %v", port)
	srv.Serve(l)
}
