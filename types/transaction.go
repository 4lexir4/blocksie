package types

import (
	"crypto/sha256"

	"github.com/4lexir4/blocksie/proto"

	pb "google.golang.org/protobuf/proto"
)

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}
