package login

import "github.com/eneskzlcn/dictionary-app-cli/cli"

func NewLoginCommand(run func(command *cli.Command, args []string)) *cli.Command {
	flags := []*cli.Flag{
		cli.NewStringFlag(DefaultValue, EmailFlagName, EmailFlagShorthand, EmailFlagUsage, NoOptionDefaultValue,false),
		cli.NewStringFlag(DefaultValue, PasswordFlagName, PasswordFlagShorthand, PasswordFlagUsage, NoOptionDefaultValue,false),
	}
	command := cli.NewCommand(Name,ShortDesc,LongDesc,nil,run,flags)
	return command
}

const (
	ShortDesc = "Short desc of login"
	LongDesc = "Long desc of login"
	DefaultValue = "default"
	NoOptionDefaultValue= "no-option"
	EmailFlagName = "email"
	EmailFlagShorthand = "e"
	EmailFlagUsage = "Set your email"
	PasswordFlagName = "password"
	PasswordFlagShorthand = "p"
	PasswordFlagUsage = "Set your password"
)

