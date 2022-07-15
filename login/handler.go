package login

import (
	"github.com/eneskzlcn/dictionary-app-cli/cli"
	"github.com/eneskzlcn/dictionary-app-cli/util"
)

type LoginService interface {
	Login(request SignInRequest) error
}
type Handler struct {
	service LoginService
	command *cli.Command
}

const (
	Name = "login"
)


func (h *Handler) init() {
	h.command = NewLoginCommand(h.CommandRun)
}

func NewHandler(service LoginService) *Handler {
	handler := &Handler{
		service: service,
	}
	handler.init()
	return handler
}

func (h *Handler) CommandRun(command *cli.Command, args []string) {
	email := Email(h.command.GetFlag("email").Value.String())
	password := Password(h.command.GetFlag("password").Value.String())

	if len(email) > 0 && len(password) > 0 {
		if err := email.Validate(); err != nil {
			print(err)
			return
		}
		if err := password.Validate(); err != nil {
			print(err)
			return
		}
		if err := h.service.Login(SignInRequest{
			Email:    email,
			Password: password,
		}); err!= nil {
			print("Error occurred when logging in to the server")
			return
		}
	}
	if email.String() == DefaultValue && password.String() == DefaultValue {
		//get email and password in order and separately from command line.
		if err := util.ReadStringFromCli("email", ); err != nil {
			return
		}
		if err := util.ReadStringFromCli("password", &password)
	}
}

func (h *Handler) ExecuteCommand() error {
	return h.command.Execute()
}

func (h *Handler) GetName() string {
	return Name
}

func (h *Handler) GetCommand() *cli.Command {
	return h.command
}
