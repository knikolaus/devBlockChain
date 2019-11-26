package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

/**
 * @Author: singHwang
 * @File:  BlockchainIterator
 * @Version: 1.0.0
 * @Date: 11/26/19 5:37 PM
 */

type BlockchainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

//　迭代器
func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}
func (blockchainIterator *BlockchainIterator) Next() *Block {
	var block *Block

	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			currentBlockBytes := b.Get(blockchainIterator.CurrentHash)

			//获取当前迭代器里面的currentHash对应的区块
			block = DeserializeBlock(currentBlockBytes)

			blockchainIterator.CurrentHash = block.PrevBlockHash
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return block
}
