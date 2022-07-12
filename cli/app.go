package cli

import (
	"errors"
)

type Handler interface {
	CommandRun(command *Command, args []string)
	ExecuteCommand() error
	GetName() string
}

type App struct {
	handlers map[string]Handler
	name string
}
func NewApp(handlers []Handler, name string) *App {
	handlerMap := make(map[string]Handler,1)
	for _, handler := range handlers {
		handlerMap[handler.GetName()] = handler
	}
	return &App{handlers: handlerMap, name: name}
}

func (app *App) GetName() string {
	return app.name
}

func(app *App) Start() error{
	rootHandler := app.handlers[app.name]
	if rootHandler != nil {
		return rootHandler.ExecuteCommand()
	}
	return errors.New("root handler not found")
}
