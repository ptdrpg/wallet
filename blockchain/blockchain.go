package blockchain

type Blockchain struct {
	GenesisBlock Block
	Chain        []Block
	Nodes  map[string]bool
	Difficulty   int
}
