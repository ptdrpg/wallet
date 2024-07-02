package blockchain

import "time"

func CreateBlockChain (difficulty int) Blockchain {
	genesisBlock := Block{
		Hash: "0",
		TimeStampe: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		make(map[string]bool),
		difficulty,
	}
}