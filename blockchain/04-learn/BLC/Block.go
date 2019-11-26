package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	// 1. 当前区块的高度
	Height int64

	//2. 上一个区块的哈希
	PrveBlockHash []byte

	//3. 交易数据
	Data []byte

	//4. 时间戳
	Timestamp int64

	//5.哈希
	Hash []byte
}

func (block *Block) setHash() {
	height := InttoHex(block.Height)
	fmt.Println(height)
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)
	//fmt.Printf(timeString)
	blockBytes := bytes.Join([][]byte{height, block.PrveBlockHash, block.Data, timeBytes}, []byte{})

	hash := sha256.Sum256(blockBytes)

	//分片
	block.Hash = hash[:]

}

func NewBlock(data string, height int64, preBlockHash []byte) *Block {

	block := &Block{height, preBlockHash, []byte(data), time.Now().Unix(), nil}

	block.setHash()
	return block
}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
