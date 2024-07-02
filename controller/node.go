package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) NodeRegister(ctx *gin.Context) {
	var node struct {
		Node string `json:"node"`
	}
	if err:= ctx.ShouldBindJSON(&node); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.B.Nodes[node.Node] = true
	ctx.JSON(http.StatusOK, c.B)
}