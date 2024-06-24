package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/wallet/controller"
)

type Router struct {
	R *gin.Engine
	C *controller.Controller
}

func NewRouter(r *gin.Engine, c *controller.Controller) *Router {
	return &Router{
		R: r,
		C: c,
	}
}

func (r *Router) RegisterRouter() {
	
}
