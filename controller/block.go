package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func(c *Controller) GetBlock(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.B.Chain)
}