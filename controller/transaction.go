package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/wallet/entity"
	"github.com/ptdrpg/wallet/lib"
)

func (c *Controller) FindAllTransaction(ctx *gin.Context){
	transactions, err := c.R.FindAllTransaction()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, transactions)
}

func (c *Controller) CreateTransaction(ctx *gin.Context) {
	var input entity.Transaction
	binderr := ctx.ShouldBindJSON(&input)
	if binderr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": binderr.Error(),
		})
		return
	}

	sender, sendErr := c.R.FindWalletById(input.From)
	if sendErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": sendErr.Error(),
		})
		return
	}

	if sender.Balance < input.Amount {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid operation, balance insufficient",
		})
		return
	}

	receiver, receiveErr := c.R.FindWalletById(input.To)
	if receiveErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": receiveErr.Error(),
		})
		return
	}

	input.UID = lib.GenerateUID()
	input.CreatedAt = time.Now()
	input.Status = "Pending"
	c.B.AddBlock(input)
	if !c.B.IsValid() {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalide chain",
		})
		return
	}


	sender.Balance = sender.Balance - input.Amount
	receiver.Balance = receiver.Balance + input.Amount
	updateSender := c.R.UpdateWallet(&sender)
	if updateSender != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": updateSender.Error(),
		})
		return
	}

	updateReceiver := c.R.UpdateWallet(&receiver)
	if updateReceiver != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": updateReceiver.Error(),
		})
		return
	}

	err := c.R.CreateTransaction(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, input)
}