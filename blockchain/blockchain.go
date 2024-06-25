package blockchain

type Blockchain struct {
	GenesisBlock Block
	Chain        []Block
	Difficulty   int
}
