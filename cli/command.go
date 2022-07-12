package cli

import (
	"github.com/spf13/cobra"
)

type Command struct {
	command *cobra.Command
	flags map[string]*Flag
	subCommands map[string]*Command
}

type CommandRun func(*Command, []string)

func NewCommand(name, shortDesc, longDesc string, subCommands []*Command, handler CommandRun, flags []*Flag) *Command {
	cobraCmd := cobra.Command{
		Use:        name,
		Aliases:    []string{},
		SuggestFor: []string{},
		Short:      shortDesc,
		Long:       longDesc,
	}
	flagsMap := make(map[string]*Flag, 1)

	subCommandsMap := make(map[string]*Command, 1)
	for _, subCommand := range subCommands {
		cobraCmd.AddCommand(subCommand.command)
		subCommandsMap[subCommand.GetName()] = subCommand
	}
	command := &Command{command: &cobraCmd, flags: flagsMap, subCommands: subCommandsMap}
	command.command.Run = func(cmd *cobra.Command, args []string) {
		handler(command, args)
	}
	command.AddFlags(flags)
	return command
}
func (c *Command) AddCommand(command *Command) {
	c.command.AddCommand(command.command)
	c.subCommands[command.GetName()] = command
}
func (c *Command) AddCommands(commands []*Command) {
	for _, command := range commands {
		c.AddCommand(command)
	}
}
func (c *Command) GetName() string {
	return c.command.Use
}
func (c *Command) AddFlag(flag *Flag) {
	c.command.Flags().VarP(flag.Value,flag.Name,flag.Shorthand,flag.Usage)
	c.flags[flag.Name] = flag
	if flag.IsNoOption == true {
		c.SetFlagNoOptionDefaultValue(flag.Name, flag.NoOptionDefaultValue)
	}
}
func (c *Command) AddFlags(flags []*Flag) {
	for _, flag := range flags {
		c.AddFlag(flag)
	}
}
func (c *Command) GetFlag(name string) *Flag{
	return c.flags[name]
}
func (c *Command) SetFlagNoOptionDefaultValue(name ,value string) {
	c.command.Flags().Lookup(name).NoOptDefVal = value

}
func (c *Command) Execute() error {
	return c.command.Execute()
}