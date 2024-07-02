package node

import (
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/wallet/app"
	"github.com/ptdrpg/wallet/blockchain"
	"github.com/ptdrpg/wallet/controller"
	"github.com/ptdrpg/wallet/repository"
	"github.com/ptdrpg/wallet/router"
)

func FullNode() {
	block := blockchain.CreateBlockChain(2)
	mainR := gin.Default()
	app.Connexion()
	db := app.DB
	repo := repository.NewRepository(db)
	c := controller.NewController(db, repo, &block)
	r := router.NewRouter(mainR, c)
	r.FNRouter()
	r.R.Run(":4401")
}
