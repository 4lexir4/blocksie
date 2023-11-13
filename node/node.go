package node

import (
	"context"
	"fmt"

	"github.com/4lexir4/blocksie/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.None, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("Received tx from:", peer)
	return &proto.None{}, nil
}
