package main

import "devBlockChain/blockchain/part24-persistence-cli/BLC"

/**
 * @Author: singHwang
 * @File:  main
 * @Version: 1.0.0
 * @Date: 11/26/19 7:49 PM
 */

func main() {
	blockchain := BLC.CreateBlockchainWithGenesisBlock()

	cli := BLC.CLI{blockchain}

	cli.Run()
}
