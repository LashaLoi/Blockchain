package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"../proto"
	"google.golang.org/grpc"
)

var (
	addFlag  = flag.Bool("add", false, "add new block")
	listFlag = flag.Bool("list", false, "get the blockchain")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}

	client := proto.NewBlockchainClient(conn)

	if *addFlag {
		b, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
			Data: time.Now().String(),
		})

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(b)
		return
	}

	if *listFlag {
		bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})

		if err != nil {
			log.Fatal(err)
			return
		}

		for _, b := range bc.Blocks {
			fmt.Println(b)
		}

		return
	}

	log.Fatal("You are not provide any flat, please use --add or --list flag to make a request")
}
