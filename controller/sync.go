package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ptdrpg/wallet/blockchain"
)

func (c *Controller) SynchBlockchain() {
	for node := range c.B.Nodes {
		response, _ := http.Get(node)
		fmt.Println(response.Body)
		var otherBlocks []blockchain.Block
		if err := json.NewDecoder(response.Body).Decode(&otherBlocks); err == nil {
			if len(otherBlocks) > len(c.B.Chain) {
				c.B.Chain = otherBlocks
			}
		}
	}
}
