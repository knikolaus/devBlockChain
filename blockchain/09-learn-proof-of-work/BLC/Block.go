package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
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

	//6. 工作量证明
	Nonce int64
}

//func (block *Block) setHash() {
//	height := InttoHex(block.Height)
//	fmt.Println(height)
//	timeString := strconv.FormatInt(block.Timestamp,2)
//	timeBytes := []byte(timeString)
//	//fmt.Printf(timeString)
//	blockBytes := bytes.Join([][]byte{height,block.PrveBlockHash,block.Data,timeBytes}, []byte{})
//
//
//	hash := sha256.Sum256(blockBytes)
//
//
//	//分片
//	block.Hash = hash[:]
//
//
//
//}

func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//反序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

func NewBlock(data string, height int64, preBlockHash []byte) *Block {

	block := &Block{height, preBlockHash, []byte(data), time.Now().Unix(), nil, 0}

	// 创建新区块的时候添加工作量证明
	pow := NewProofOfWork(block)

	hash, nonce := pow.Run()

	block.Hash = hash[:]

	block.Nonce = nonce

	return block
}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
