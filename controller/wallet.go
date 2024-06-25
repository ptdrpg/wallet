package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/wallet/entity"
	"github.com/ptdrpg/wallet/lib"
)

func (c *Controller) FindAllWallet(ctx *gin.Context) {
	wallets, err := c.R.FindAllWallet()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, wallets)
}

func (c *Controller) FindWalletById(ctx *gin.Context) {
	walletId := ctx.Param("uid")
	wallet, err := c.R.FindWalletById(walletId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, wallet)
}

type WalletInput struct {
	OwnerUID  string `json:"owner_uid"`
	PublicKey string `json:"public_key"`
}

func (c *Controller) CreateWalletWithInput(input WalletInput) (*entity.Wallet, error) {
	wallet := &entity.Wallet{
		UID:       lib.GenerateUID(),
		OwnerUID:  input.OwnerUID,
		Balance:   0,
		CreatedAt: time.Now().String(),
		PublicKey: input.PublicKey,
	}

	if err := c.R.CreateWallet(wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}

func (c *Controller) CreateWallet(ctx *gin.Context) {
	var input WalletInput
	err := ctx.ShouldBindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	wallet, cErr := c.CreateWalletWithInput(input)
	if cErr != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": cErr.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, wallet)
}
