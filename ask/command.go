package ask

import "github.com/eneskzlcn/dictionary-app-cli/cli"

func NewAskCommand(run func(command *cli.Command, args []string)) *cli.Command {
	flags := []*cli.Flag{
		cli.NewStringFlag(QuestionFlagNoOptionValue,QuestionTypeFlagName,"t",QuestionTypeFlagUsage,QuestionFlagNoOptionValue,true),
	}
	command := cli.NewCommand(Name,ShortDesc,LongDesc,nil, run,flags)
	return command
}

const (
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
