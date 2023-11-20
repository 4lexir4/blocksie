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
	makeNode(":3000", []string{})
	time.Sleep(time.Second)
	makeNode(":4000", []string{":3000"})
	time.Sleep(time.Second)
	makeNode(":5000", []string{":4000"})

	time.Sleep(time.Second)
	makeTransaction()

	select {}
}

func makeNode(listenAddr string, boostrapNodes []string) *node.Node {
	n := node.NewNode()
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
