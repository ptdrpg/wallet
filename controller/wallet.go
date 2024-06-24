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

type Input struct {
	OwnerUID   string `json:"owner_uid"`
	PublicKey string `json:"security_code"`
}

func (c *Controller) CreateWallet(ctx *gin.Context) {
	var wallet entity.Wallet
	var input Input
	err := ctx.ShouldBindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	wallet.UID = lib.GenerateUID()
	wallet.OwnerUID = input.OwnerUID
	wallet.Balance = 0
	wallet.CreatedAt = time.Now().String()
	wallet.PublicKey = input.PublicKey

	err = c.R.CreateWallet(&wallet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, wallet)
}
