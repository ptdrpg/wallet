package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ptdrpg/wallet/entity"
)

func History(data Block) error {
	outputDir := "upload/transaction/"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return errors.New(err.Error())
		}
	}

	filename := fmt.Sprintf("%v%v.json", data.Data.Type, data.Data.UID)
	filepath := outputDir + filename
	_, Create := os.Create(filepath)
	if Create != nil {
		return errors.New(Create.Error())
	}

	newData, err := json.Marshal(data)
	if err != nil {
		return errors.New(err.Error())
	}

	WriteFile := os.WriteFile(filepath, newData, 0755)
	if WriteFile != nil {
		return errors.New(WriteFile.Error())
	}

	return nil
}

func (b *Blockchain) AddBlock(data entity.Transaction) {
	lastBlock := b.Chain[len(b.Chain)-1]
	newBlock := Block{
		PreviousHash: lastBlock.Hash,
		Data:         data,
		TimeStampe:   time.Now(),
	}
	newBlock.Mine(b.Difficulty)
	History(Block{
		PreviousHash: newBlock.PreviousHash,
		Data:         newBlock.Data,
	})
	b.Chain = append(b.Chain, newBlock)
}
