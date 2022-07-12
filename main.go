package main

import (
	"github.com/eneskzlcn/dictionary-app-cli/ask"
	"github.com/eneskzlcn/dictionary-app-cli/cli"
	"github.com/eneskzlcn/dictionary-app-cli/config"
	"github.com/eneskzlcn/dictionary-app-cli/root"
	"os"
)
var app *cli.App

func init() {
	askService := ask.NewService()
	askHandler := ask.NewHandler(askService)
	rootService := root.NewService()
	rootHandler := root.NewHandler(rootService)


	app = cli.NewApp([]cli.Handler{
		askHandler,
	},rootHandler, config.AppConfig.Name)
}
func main() {

	if err := app.Start(); err != nil {
		print(err.Error())
		os.Exit(1)
	}
}
