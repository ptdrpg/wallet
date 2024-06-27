package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ptdrpg/wallet/blockchain"
)

func History (data blockchain.Block) error {
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

	WriteFile:= os.WriteFile(filepath, newData, 0755)
	if WriteFile != nil{
		return errors.New(WriteFile.Error())
	}

	return nil
}