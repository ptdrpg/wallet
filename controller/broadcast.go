package controller

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ptdrpg/wallet/entity"
)

func (c *Controller) Broadcast(data entity.Transaction) {
	for node := range c.B.Nodes {
		url := node
		tempdata, _:= json.Marshal(data)
		newData := string(tempdata)
		jsonStr, _:= json.Marshal(map[string]string{"data": newData})
		http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	}
}