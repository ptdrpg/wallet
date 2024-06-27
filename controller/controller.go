package controller

import (
	"github.com/ptdrpg/wallet/blockchain"
	"github.com/ptdrpg/wallet/repository"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
	R  *repository.Repository
	B  *blockchain.Blockchain
}

func NewController(db *gorm.DB, r *repository.Repository, b *blockchain.Blockchain) *Controller {
	return &Controller{
		DB: db,
		R:  r,
		B:  b,
	}
}
