package ask

import (
	"github.com/eneskzlcn/dictionary-app-cli/cli"
	"log"
)
type AskService interface {
	Ask(questionType string) error
}
type Handler struct {
	command *cli.Command
	service	AskService
}

func NewHandler(service AskService) *Handler {
	handler := &Handler{}
	handler.service = service
	handler.init()
	return handler
}
func (h *Handler) init() {
	h.command = NewAskCommand(h.CommandRun)
}

const (
	Name = "ask"
)

func (h *Handler) CommandRun(command *cli.Command, args []string){
	questionType := command.GetFlag(QuestionTypeFlagName).Value.String()
	err := h.service.Ask(questionType)
	if err != nil {
		log.Println(err)
	}
}
func (h *Handler) ExecuteCommand() error {
	return h.command.Execute()
}
func (h *Handler) AddSubHandler(handler Handler) {
	h.command.AddCommand(handler.command)
}
func (h *Handler) GetName() string {
	return Name
}
func (h *Handler) GetCommand() *cli.Command{
	return h.command
}
