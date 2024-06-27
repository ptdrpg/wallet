package blockchain

func (b *Blockchain) IsValid() bool {
	for i := range b.Chain[1:] {
		previousBlock := b.Chain[i]
		currentBlock := b.Chain[i+1]
		if currentBlock.Hash != currentBlock.CalculeHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}

	return true
}