package main

import (
	"devBlockChain/blockchain/02-learn/BLC"
	"fmt"
)

func main() {
	genesisBlock := BLC.CreateBlockchainWithGenesisBlock()

	fmt.Println(genesisBlock)

	fmt.Println(genesisBlock.Blocks)

	fmt.Println(genesisBlock.Blocks[0])

}
