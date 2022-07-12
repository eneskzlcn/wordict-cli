package root

import (
	"github.com/eneskzlcn/dictionary-app-cli/cli"
	"github.com/eneskzlcn/dictionary-app-cli/config"
	"os"
)
type RootService interface {
	LoadSession() error
}
type Handler struct {
	command *cli.Command
	service	RootService
}
func NewHandler(service RootService) *Handler {
	handler := &Handler{}
	handler.service = service
	handler.init()
	return handler
}
func (h *Handler) init() {
	h.command = cli.NewCommand(config.AppConfig.Name,ShortDesc,LongDesc,nil,h.CommandRun,nil)
}

const (
	Name = "wordict"
	ShortDesc = "Wordict is an application to learn english"
	LongDesc = `Wordict designed as a long term english learning process.`
)

func (h *Handler) CommandRun(command *cli.Command, args []string){
	if err := h.service.LoadSession(); err!= nil {
		print(err.Error())
		os.Exit(1)
	}
	print("Root command started")
}
func (h *Handler) ExecuteCommand() error {
	return h.command.Execute()
}
func (h *Handler) GetName() string {
	return Name
}
func (h *Handler) GetCommand() *cli.Command{
	return h.command
}
func (h *Handler) AddSubHandler(handler cli.Handler) {
	h.command.AddCommand(handler.GetCommand())
}