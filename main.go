package main

import (
	cli "github.com/eneskzlcn/incli"
	"github.com/eneskzlcn/dictionary-app-cli/client"
	"github.com/eneskzlcn/dictionary-app-cli/config"
	ask "github.com/eneskzlcn/dictionary-app-cli/internal/ask"
	login "github.com/eneskzlcn/dictionary-app-cli/internal/login"
	root "github.com/eneskzlcn/dictionary-app-cli/internal/root"
	"os"
)
var app *cli.App

func init() {
	rootService := root.NewService()
	rootHandler := root.NewHandler(rootService)

	askService := ask.NewService()
	askHandler := ask.NewHandler(askService)

	restClient := client.NewClient("x")

	loginClient := login.NewClient(restClient)
	loginService := login.NewService(loginClient)
	loginHandler := login.NewHandler(loginService)

	app = cli.NewApp([]cli.Handler{
		askHandler,
		loginHandler,

	},rootHandler, config.AppConfig.Name)
}
func main() {

	if err := app.Start(); err != nil {
		print(err.Error())
		os.Exit(1)
	}
}
