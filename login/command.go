package login

import "github.com/eneskzlcn/dictionary-app-cli/cli"

func NewLoginCommand(run func(command *cli.Command, args []string)) *cli.Command {
	flags := []*cli.Flag{
		cli.NewStringFlag(DefaultValue,"email","e","Set your email",NoOptionDefaultValue,false),
		cli.NewStringFlag(DefaultValue,"password","p","Set your password",NoOptionDefaultValue,false),
	}
	command := cli.NewCommand(Name,ShortDesc,LongDesc,nil,run,flags)
	return command
}

const (
	ShortDesc = ""
	LongDesc = ""
	DefaultValue = "default"
	NoOptionDefaultValue= "no-option"
)

