package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/4lexir4/blocksie/node"
	"github.com/4lexir4/blocksie/proto"
	"google.golang.org/grpc"
)

func main() {
	node := node.NewNode()

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("Node running on prot:", ":3000")
	grpcServer.Serve(ln)
}

func makeTransaction() {
	client, err := grpc.Dial(":3000")
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)

	_, err = c.HandleTransaction(context.TODO(), &proto.Transaction{})
	if err != nil {
		log.Fatal(err)
	}

}
