package node

import (
	"context"
	"encoding/hex"
	"net"
	"sync"

	"github.com/4lexir4/blocksie/proto"
	"github.com/4lexir4/blocksie/types"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Mempool struct {
	txx map[string]*proto.Transaction
}

func NewMempool() *Mempool {
	return &Mempool{
		txx: make(map[string]*proto.Transaction),
	}
}

func (pool *Mempool) Has(tx *proto.Transaction) bool {
	hash := hex.EncodeToString(types.HashTransaction(tx))
	_, ok := pool.txx[hash]
	return ok
}

func (pool *Mempool) Add(tx *proto.Transaction) bool {
	if pool.Has(tx) {
		return false
	}
	hash := hex.EncodeToString(types.HashTransaction(tx))
	pool.txx[hash] = tx
	return true
}

type Node struct {
	version    string
	listenAddr string
	logger     *zap.SugaredLogger

	peerLock sync.RWMutex
	peers    map[proto.NodeClient]*proto.Version
	mempool  *Mempool

	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.EncoderConfig.TimeKey = ""
	logger, _ := loggerConfig.Build()
	return &Node{
		peers:   make(map[proto.NodeClient]*proto.Version),
		version: "blocksie-0.1",
		logger:  logger.Sugar(),
		mempool: NewMempool(),
	}
}

func (n *Node) Start(listenAddr string, bootstrapNodes []string) error {
	n.listenAddr = listenAddr
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	proto.RegisterNodeServer(grpcServer, n)

	n.logger.Infow("Node started...", "port", n.listenAddr)

	// Bootstrap the network with a list of already known nodes in the network
	if len(bootstrapNodes) > 0 {
		go n.bootstrapNetwork(bootstrapNodes)
	}

	return grpcServer.Serve(ln)
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	c, err := makeNodeClient(v.ListenAddr)
	if err != nil {
		return nil, err
	}

	n.addPeer(c, v)

	return n.getVersion(), nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	hash := hex.EncodeToString(types.HashTransaction(tx))

	if n.mempool.Add(tx) {
		n.logger.Debugw("Received transaxtion", "from", peer.Addr, "hash", hash, "we", n.listenAddr)
		go func() {
			if err := n.broadcast(tx); err != nil {
				n.logger.Errorw("Broadcast error", "err", err)
			}
		}()
	}
	return &proto.Ack{}, nil
}

func (n *Node) broadcast(msg any) error {
	for peer := range n.peers {
		switch v := msg.(type) {
		case *proto.Transaction:
			_, err := peer.HandleTransaction(context.Background(), v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()

	n.peers[c] = v

	// Connect to all the peers in the received peer list
	if len(v.PeerList) > 0 {
		go n.bootstrapNetwork(v.PeerList)
	}

	n.logger.Debugw("New peer successfully connected",
		"we", n.listenAddr,
		"remoteNode", v.ListenAddr,
		"height", v.Height)
}

func (n *Node) deletePeer(c proto.NodeClient) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()

	delete(n.peers, c)
}

func (n *Node) bootstrapNetwork(addrs []string) error {
	for _, addr := range addrs {
		if !n.canConnectWith(addr) {
			continue
		}
		n.logger.Debugw("Dialing remote node", "we", n.listenAddr, "remote", addr)
		c, v, err := n.dialRemoteNode(addr)
		if err != nil {
			return err
		}
		n.addPeer(c, v)
	}
	return nil
}

func (n *Node) dialRemoteNode(addr string) (proto.NodeClient, *proto.Version, error) {
	c, err := makeNodeClient(addr)
	if err != nil {
		return nil, nil, err
	}

	v, err := c.Handshake(context.Background(), n.getVersion())
	if err != nil {
		return nil, nil, err
	}
	return c, v, nil
}

func (n *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version:    "blocksie-0.1",
		Height:     0,
		ListenAddr: n.listenAddr,
		PeerList:   n.getPeerList(),
	}
}

func (n *Node) canConnectWith(addr string) bool {
	if n.listenAddr == addr {
		return false
	}
	connectedPeers := n.getPeerList()
	for _, connectedAddr := range connectedPeers {
		if addr == connectedAddr {
			return false
		}
	}
	return true
}

func (n *Node) getPeerList() []string {
	n.peerLock.RLock()
	defer n.peerLock.RUnlock()

	peers := []string{}
	for _, version := range n.peers {
		peers = append(peers, version.ListenAddr)
	}

	return peers
}

func makeNodeClient(listenAddr string) (proto.NodeClient, error) {
	c, err := grpc.Dial(listenAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return proto.NewNodeClient(c), nil
}
