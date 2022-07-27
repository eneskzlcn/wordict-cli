package root

import (
	"github.com/eneskzlcn/dictionary-app-cli/cli"
	"os"
)
type RootService interface {
	LoadSession() error
}
type Handler struct {
	command *cli.Command
	service RootService
}
func NewHandler(service RootService) *Handler {
	handler := &Handler{}
	handler.service = service
	handler.init()
	return handler
}
func (h *Handler) init() {
	h.command = NewRootCommand(h.CommandRun)
}

const (
	Name = "wordict"

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