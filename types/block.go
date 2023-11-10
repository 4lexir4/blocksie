package types

import (
	"crypto/sha256"

	"github.com/4lexir4/blocksie/crypto"
	"github.com/4lexir4/blocksie/proto"
	//pb "github.com/golang/protobuf/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	return pk.Sign(HashBlock(b))
}

func HashBlock(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}