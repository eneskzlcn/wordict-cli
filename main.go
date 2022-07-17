package main

import (
	"github.com/eneskzlcn/dictionary-app-cli/ask"
	"github.com/eneskzlcn/dictionary-app-cli/cli"
	"github.com/eneskzlcn/dictionary-app-cli/client"
	"github.com/eneskzlcn/dictionary-app-cli/config"
	"github.com/eneskzlcn/dictionary-app-cli/login"
	"github.com/eneskzlcn/dictionary-app-cli/root"
	"os"
)
var app *cli.App

func init() {
	rootService := root.NewService()
	rootHandler := root.NewHandler(rootService)

	askService := ask.NewService()
	askHandler := ask.NewHandler(askService)

	loginClient := client.NewClient("x")
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
