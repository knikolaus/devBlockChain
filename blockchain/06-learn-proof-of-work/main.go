package main

import (
	"devBlockChain/blockchain/06-learn-proof-of-work/BLC"
	"fmt"
)

func main() {
	blockchain := BLC.CreateBlockchainWithGenesisBlock()

	//新区块
	blockchain.AddBlockToBlackchain("send 1RMB to xinghua.huang", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlackchain("send 2RMB to xinghua.huang", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlackchain("send 3RMB to xinghua.huang", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlackchain("send 4RMB to xinghua.huang", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	fmt.Println(blockchain)
}
