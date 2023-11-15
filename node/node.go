package node

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/4lexir4/blocksie/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version string

	peerLock sync.RWMutex
	peers    map[proto.NodeClient]*proto.Version

	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		peers:   make(map[proto.NodeClient]*proto.Version),
		version: "blocksie-0.1",
	}
}

func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()

	fmt.Printf("New peer connected (%s) - height (%d)\n", v.ListenAddr, v.Height)

	n.peers[c] = v
}

func (n *Node) deletePeer(c proto.NodeClient) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()

	delete(n.peers, c)
}

func (n *Node) Start(listenAddr string) error {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	proto.RegisterNodeServer(grpcServer, n)

	fmt.Println("Node running on prot:", listenAddr)

	return grpcServer.Serve(ln)
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	c, err := makeNodeClient(v.ListenAddr)
	if err != nil {
		return nil, err
	}

	n.addPeer(c, v)

	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("Received tx from:", peer)
	return &proto.Ack{}, nil
}

func makeNodeClient(listenAddr string) (proto.NodeClient, error) {
	c, err := grpc.Dial(listenAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return proto.NewNodeClient(c), nil
}
