package BLC

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

const targetBit = 16

type ProofOfWork struct {
	Block *Block

	target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{pow.Block.PrveBlockHash, pow.Block.Data, InttoHex(pow.Block.Timestamp), InttoHex(int64(targetBit)), InttoHex(int64(nonce)), InttoHex(int64(pow.Block.Height))},
		[]byte{})
	return data

}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0

	var hashInt big.Int
	var hash [32]byte

	for {
		dataBytes := proofOfWork.prepareData(nonce)

		hash = sha256.Sum256(dataBytes)

		hashInt.SetBytes(hash[:])

		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce = nonce + 1

	}
	return hash[:], int64(nonce)
}

//　创建工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {

	target := big.NewInt(1)

	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}
