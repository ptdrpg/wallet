package blockchain

import (
	"time"

	"github.com/ptdrpg/wallet/entity"
)

type Block struct {
	PreviousHash string
	Data         entity.Transaction
	Hash         string
	TimeStampe   time.Time
	Pow          int
}
