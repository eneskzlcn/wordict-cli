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
	flags := []*cli.Flag{
		cli.NewStringFlag(QuestionFlagNoOptionValue,QuestionTypeFlagName,"t",QuestionTypeFlagUsage,QuestionFlagNoOptionValue,true),
	}
	h.command = cli.NewCommand(Name,ShortDesc,LongDesc,nil,h.CommandRun,flags)
}

const (
	Name = "ask"
	ShortDesc = "Asks you about a word."
	LongDesc = "That commands asks you a question arround your words. Question type can be specified by setting --type flag."
	QuestionTypeFlagName = "type"
	QuestionTypeFlagUsage = `Specify the asking question type. Default chooses randomly. If you do not specify
	an option to flag, it will be set as random too.

	synonym: The question will be asked is the synonym of a random word.
	opposite: The question will be asked is the opposite of a random word.
	translation: The question will be asked is the translation of a random word.
`
	QuestionFlagNoOptionValue = "random"
	QuestionFlagTranslationOption = "translation"
	QuestionFlagOppositeOption = "opposite"
	QuestionFlagSynonymOption = "synonym"
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
