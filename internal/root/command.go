package root

import (
	"github.com/eneskzlcn/dictionary-app-cli/cli"
	"github.com/eneskzlcn/dictionary-app-cli/config"
)

func NewRootCommand(run func(command *cli.Command, args []string)) *cli.Command {
	subCommands := make([]*cli.Command, 0)
	flags := make([]*cli.Flag, 0)
	return cli.NewCommand(config.AppConfig.Name, ShortDesc, LongDesc,subCommands, run, flags)
}

const (
	ShortDesc = "Wordict is an application to learn english"
	LongDesc = `Wordict designed as a long term english learning process.`
)
