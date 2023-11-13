package types

import (
	"testing"

	"github.com/4lexir4/blocksie/crypto"
	"github.com/4lexir4/blocksie/proto"
	"github.com/4lexir4/blocksie/util"
)

func TestNewTransaction(t *testing.T) {
	fromPrvKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrvKey.Public().Address().Bytes()

	toPrvKey := crypto.GeneratePrivateKey()
	toAddress := toPrvKey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrvHash:     util.RandomHash(),
		PrvOutIndex: 0,
		PublicKey:   fromPrvKey.Public().Bytes(),
	}

	output1 := &proto.TxOutput{
		Amout:   4,
		Address: toAddress,
	}

	output2 := &proto.TxOutput{
		Amout:   95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}
}
