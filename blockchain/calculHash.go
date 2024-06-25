package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
)

func (b *Block) CalculeHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + b.TimeStampe.String() + strconv.Itoa(b.Pow)
	blockHash := sha256.Sum256([]byte(blockData))

	return fmt.Sprintf("%x", blockHash)
}