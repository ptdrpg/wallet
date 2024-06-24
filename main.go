package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/wallet/app"
	"github.com/ptdrpg/wallet/controller"
	"github.com/ptdrpg/wallet/repository"
	"github.com/ptdrpg/wallet/router"
)

func main () {
	// gin.SetMode(gin.ReleaseMode)
	mainR:= gin.Default()
	app.Connexion()
	db := app.DB
	repo := repository.NewRepository(db)
	controller := controller.NewController(db, repo)
	r := router.NewRouter(mainR, controller)
	r.RegisterRouter()

	r.R.Run(":4400")
}