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

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewBlockchainClient(conn)

	if *addFlag {
		b, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
			Data: time.Now().String(),
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(b)
	}

	if *listFlag {
		bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})

		if err != nil {
			log.Fatal(err)
		}

		for _, b := range bc.Blocks {
			fmt.Println(b)
		}
	}
}
