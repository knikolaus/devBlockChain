package BLC

type Blockchain struct {
	Blocks []*Block
}

// 创建新的区块
func (blc *Blockchain) AddBlockToBlackchain(data string, height int64, preHash []byte) {
	newBlock := NewBlock(data, height, preHash)
	//newBlock1 := NewBlock(data,height,preHash)
	blc.Blocks = append(blc.Blocks, newBlock)

}

//　创造一个有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	genesisBlock := CreateGenesisBlock("hello world.....")

	return &Blockchain{[]*Block{genesisBlock}}
}
