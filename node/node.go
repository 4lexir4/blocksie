package node

import (
	"context"
	"fmt"

	"github.com/4lexir4/blocksie/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version string
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		version: "blocksie-0.1",
	}
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	p, _ := peer.FromContext(ctx)

	fmt.Printf("Received version from %s: %+v\n", v, p.Addr)

	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("Received tx from:", peer)
	return &proto.Ack{}, nil
}
