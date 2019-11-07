package grpcserver

import (
	"context"

	"../../proto"
	"../blockchain"
)

// Server is a main element witch allow you to put in itto groc register function
type Server struct {
	Blockchain *blockchain.Blockchain
}

// AddBlock allow to male a call to add block into blockchain
func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)

	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

// GetBlockchain allow you to make a call to show a blockchain
func (s *Server) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponce, error) {
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
