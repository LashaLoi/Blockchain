package main

import (
	"context"
	"log"
	"net"

	"../proto"
	"./blockchain"
	"google.golang.org/grpc"
)

// GRPCServer ...
type GRPCServer struct {
	Blockchain *blockchain.Blockchain
}

// AddBlock ...
func (s *GRPCServer) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)

	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

// GetBlockchain ...
func (s *GRPCServer) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponce, error) {
	resp := new(proto.GetBlockchainResponce)

	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			Hash:          b.Hash,
			PrevBlockHash: b.PrevBlockHash,
			Data:          b.Data,
		})
	}

	return resp, nil
}

func main() {
	port := ":8080"

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &GRPCServer{
		Blockchain: blockchain.NewBlockChain(),
	})

	log.Printf("tcp server was started on: %v", port)
	srv.Serve(l)
}
