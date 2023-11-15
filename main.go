package main

import (
	"context"
	"log"
	"time"

	"github.com/4lexir4/blocksie/node"
	"github.com/4lexir4/blocksie/proto"
	"google.golang.org/grpc"
)

func main() {
	node := node.NewNode()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			makeTransaction()
		}
	}()

	log.Fatal(node.Start(":3000"))
}

func makeNode(listenAddr string, boostrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(listenAddr)
	if err := n.BootstrapNetwork(boostrapNodes); err != nil {
		log.Fatal(err)
	}

	return n
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version:    "blocksie-0.1",
		Height:     1,
		ListenAddr: ":4000",
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}

}
