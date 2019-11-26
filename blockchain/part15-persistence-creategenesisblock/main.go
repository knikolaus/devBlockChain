package main

import (
	"devBlockChain/blockchain/part15-persistence-creategenesisblock/BLC"
)

func main() {

	////data string,height int64,prevBlockHash []byte
	//block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	//fmt.Printf("%d\n",block.Nonce)
	//fmt.Printf("%x\n",block.Hash)
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

}
