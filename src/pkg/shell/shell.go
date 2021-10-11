package shell

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/margostino/walpurgis/pkg/config"
	"strings"
)

type Shell struct {
	Suggestions []prompt.Suggest
}

var PowerShell *Shell

func (s *Shell) Prompt() string {
	username := config.GetConfiguration().Twitter.Username
	prefix := fmt.Sprintf("%s> ", username)
	return prompt.Input(strings.ToLower(prefix), Completer(s.Suggestions))
}

func (s *Shell) Input() string {
	return s.Prompt()
}

func Completer(suggestions []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}
}

func (s *Shell) GetOptions() []prompt.Suggest {
	return s.Suggestions
}

func NewShell() *Shell {
	suggestions := []prompt.Suggest{
		{Text: "rank users", Description: "Get friends ranking sorted by latest tweet"},
	}
	PowerShell = &Shell{Suggestions: suggestions}
	return PowerShell
}
