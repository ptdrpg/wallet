package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/wallet/app"
	"github.com/ptdrpg/wallet/blockchain"
	"github.com/ptdrpg/wallet/controller"
	"github.com/ptdrpg/wallet/repository"
	"github.com/ptdrpg/wallet/router"
	"github.com/urfave/cli"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	block := blockchain.CreateBlockChain(5)
	mainR := gin.Default()
	app.Connexion()
	db := app.DB
	repo := repository.NewRepository(db)
	c := controller.NewController(db, repo, &block)
	r := router.NewRouter(mainR, c)
	r.RegisterRouter()
	cli := &cli.App{
		Name:  "Wallet-cli",
		Usage: "provisooir wallet",
		Commands: []cli.Command{
			{
				Name:    "create",
				Aliases: []string{"C"},
				Usage:   "create new wallet",
				Action: func(cli *cli.Context) error {
					// ownerUID := cli.String("o")
					// publicKey := cli.String("k")

					input := controller.WalletInput{
						OwnerUID:  "qwoeqwer",
						PublicKey: "abcdefghijklmnopqrstuvwxyz",
					}
					c.CreateWalletWithInput(input)
					return nil
				},
			},
		},
	}

	err := cli.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}

	r.R.Run(":4400")
}
