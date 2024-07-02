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
	apiR := r.R.Group("api")
	v1 := apiR.Group("v1")
	
	wr := v1.Group("/wallet")
	wr.GET("/", r.C.FindAllWallet)
	wr.GET("/:uid", r.C.FindWalletById)
	wr.POST("/", r.C.CreateWallet)

	br := v1.Group("/blockchain")
	br.GET("/", r.C.GetBlock)


	tr := v1.Group("/transaction")
	tr.POST("/", r.C.CreateTransaction)

}

func (r *Router) FNRouter() {
	// apiR := r.R.Group("api")
	// v1 := apiR.Group("v1")


}

func (r *Router) MNRouter() {
	apiR := r.R.Group("api")
	v1 := apiR.Group("v1")

	nr := v1.Group("/register")
	nr.POST("/", r.C.NodeRegister)
	go r.C.SynchBlockchain()
}
