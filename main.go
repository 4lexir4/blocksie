package main

import (
	"context"
	"log"
	"time"

	//"time"

	"github.com/4lexir4/blocksie/crypto"
	"github.com/4lexir4/blocksie/node"
	"github.com/4lexir4/blocksie/proto"
	"github.com/4lexir4/blocksie/util"
	"google.golang.org/grpc"
)

func main() {
	makeNode(":3000", []string{}, true)
	time.Sleep(time.Second)
	makeNode(":4000", []string{":3000"}, false)
	time.Sleep(time.Second)
	makeNode(":5000", []string{":4000"}, false)

	for {
		time.Sleep(time.Second)
		makeTransaction()
	}
}

func makeNode(listenAddr string, boostrapNodes []string, isValidator bool) *node.Node {
	cfg := node.ServerConfig{
		Version:    "blocksie-0.1",
		ListenAddr: listenAddr,
	}
	if isValidator {
		cfg.PrivateKey = crypto.GeneratePrivateKey()
	}
	n := node.NewNode(cfg)
	go n.Start(listenAddr, boostrapNodes)
	return n
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)
	prvKey := crypto.GeneratePrivateKey()

	tx := &proto.Transaction{
		Version: 1,
		Inputs: []*proto.TxInput{
			{
				PrvHash:     util.RandomHash(),
				PrvOutIndex: 0,
				PublicKey:   prvKey.Public().Bytes(),
			},
		},
		Outputs: []*proto.TxOutput{
			{
				Amount:  99,
				Address: prvKey.Public().Address().Bytes(),
			},
		},
	}

	_, err = c.HandleTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatal(err)
	}

}
