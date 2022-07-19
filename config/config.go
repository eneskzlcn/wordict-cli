package config

type App struct {
	Name string
}

var AppConfig *App

func init() {
	AppConfig = &App{Name: "wordict"}
}
