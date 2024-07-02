package blockchain

import (
	"bytes"
	"fmt"
	"net"
)

func (b *Blockchain) HandleConnection(conn net.Conn) {
	defer conn.Close()

	var buff bytes.Buffer
	_, err := buff.ReadFrom(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

}