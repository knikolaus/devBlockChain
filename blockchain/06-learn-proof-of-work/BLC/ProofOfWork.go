package BLC

import (
	"math/big"
)

const targetBit = 16

type ProofOfWork struct {
	Block *Block

	target *big.Int
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
}

//　创建工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {

	target := big.NewInt(1)

	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}
