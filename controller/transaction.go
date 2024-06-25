package controller

import (
	"net/http"

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

	input.UID = lib.GenerateUID()

}