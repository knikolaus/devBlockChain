package BLC

type Blockchain struct {
	Blocks []*Block
}

func CreateBlockchainWithGenesisBlock() *Blockchain {
	genesisBlock := CreateGenesisBlock("hello world.....")

	return &Blockchain{[]*Block{genesisBlock}}
}
