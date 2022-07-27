package cli

import (
	"errors"
)
type Handler interface {
	CommandRun(command *Command, args []string)
	ExecuteCommand() error
	GetName() string
	GetCommand() *Command
}
type RootHandler interface {
	CommandRun(command *Command, args []string)
	ExecuteCommand() error
	GetName() string
	AddSubHandler(handler Handler)
}
type App struct {
	handlers map[string]Handler
	name string
	rootHandler RootHandler
}
func NewApp(handlers []Handler, rootHandler RootHandler, name string) *App {
	handlerMap := make(map[string]Handler,1)
	//all the other commands except root, need to be child of the root command to be used as <app-name> <command>
	for _, handler := range handlers {
		handlerMap[handler.GetName()] = handler
		rootHandler.AddSubHandler(handler)
	}

	return &App{handlers: handlerMap, name: name, rootHandler: rootHandler}
}

func (app *App) GetName() string {
	return app.name
}
func (app *App) AddHandler(handler Handler) error {
	if app.rootHandler == nil {
		return errors.New("there is no root handler registered to app")
	}
	app.rootHandler.AddSubHandler(handler)
	return nil
}
func (app *App) AddRootHandler(handler RootHandler) {
	app.rootHandler = handler
}
func(app *App) Start() error{
	if app.rootHandler != nil {
		return app.rootHandler.ExecuteCommand()
	}
	return errors.New("root handler not found")
}
