package main

import (
	"log"
	"net"

	"../proto"
	"./blockchain"
	"./config"
	"./grpcserver"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.NewConfig()

	l, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &grpcserver.Server{
		Blockchain: blockchain.NewBlockChain(),
	})

	log.Printf("tcp server was started on: %v", cfg.Port)
	srv.Serve(l)
}
