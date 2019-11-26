package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//数据库名
const dbName = "blockchain.db"

//表名
const blockTableName = "blocks"

type Blockchain struct {
	//Blocks []*Block  // 存储有序的区块
	//存储最新区块的hash
	Tip []byte
	DB  *bolt.DB
}

// 增加区块到区块链里面

//func (blc *Blockchain) AddBlockToBlockchain(data string,height int64,prehash []byte){
//	newBlock := NewBlock(data,height,prehash)
//	//在blockChain最后添加一个区块
//	blc.Blocks = append(blc.Blocks,newBlock)
//}

func (blc *Blockchain) AddBlockToBlockchain(data string) {
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//通过hash获取到最新区块
			blockBytes := b.Get(blc.Tip)
			block := DeserializeBlock(blockBytes)

			newBlock := NewBlock(data, block.Height+1, block.Hash)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			blc.Tip = newBlock.Hash
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	// 创建创世区块 没有就创建数据库名不是表名
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic(err)
		}

		if b == nil {
			genesisBlock := CreateGenesisBlock("hello world")
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			blockHash = genesisBlock.Hash
		}
		return nil
	})
	// 返回区块链对象
	return &Blockchain{blockHash, db}
}
