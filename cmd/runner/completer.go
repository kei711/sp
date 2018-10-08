package runner

import (
	"fmt"
	"github.com/kei711/sp/util"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"golang.org/x/sys/unix"
)

var suggestsCache []prompt.Suggest
var renderUsageCommandName string

func Completer(d prompt.Document) []prompt.Suggest {
	commands := strings.Split(d.TextBeforeCursor(), " ")

	if len(commands) > 1 && s.Contains(commands[0]) {
		commandName := commands[0]
		if renderUsageCommandName == "" || renderUsageCommandName != commandName {
			renderUsageCommandName = commandName
			command := s.GetCommand(commandName)

			ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
			if err != nil {
				panic(err)
			}
			fmt.Print("\n")
			fmt.Print("\r" + strings.Repeat(" ", int(ws.Col)))
			fmt.Println("\r" + command.Description)
			fmt.Print("\r" + strings.Repeat(" ", int(ws.Col)))
			fmt.Println("\r" + command.Usages.Usage[0])
		}
		return prompt.FilterHasPrefix(getOptionSuggests(commands), d.GetWordBeforeCursor(), true)
	}

	return util.FilterFuzzy(getCommandSuggests(), d.GetWordBeforeCursor())
}

func getCommandSuggests() []prompt.Suggest {
	if suggestsCache == nil {
		for id := range s.CommandMap {
			text := id
			suggestsCache = append(suggestsCache, prompt.Suggest{Text: text})
		}
		suggestsCache = append(suggestsCache, prompt.Suggest{Text: cmdQuit})
		suggestsCache = append(suggestsCache, prompt.Suggest{Text: cmdExit})
		suggestsCache = append(suggestsCache, prompt.Suggest{Text: cmdClearCache})
	}

	return suggestsCache
}

func clearSuggestsCache() {
	suggestsCache = nil
}

func getOptionSuggests(commands []string) []prompt.Suggest {
	var suggest []prompt.Suggest
	commandName := commands[0]
	for _, o := range s.GetCommand(commandName).Options.Option {
		if util.ArrayContains(commands, o.Name) {
			continue
		}
		suggest = append(suggest, prompt.Suggest{Text: o.Name})
		if o.Shortcut != "" {
			suggest = append(suggest, prompt.Suggest{Text: o.Shortcut})
		}
	}

	return suggest
}
