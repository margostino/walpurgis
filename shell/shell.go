package shell

import (
	"github.com/margostino/griffin/pkg/griffin"
	"github.com/margostino/walpurgis/config"
)

var appShell *griffin.Shell

func Initialize() {
	appShell = griffin.New().
		SetPrompt(config.GetUsername() + "@walpurgis").
		SetSimpleActions(SimpleActionsMapping).
		SetMultiParamsActions(InputStringsActionMapping).
		SetConfiguration(config.CommandsConfiguration())
}

func Start() {
	appShell.Start()
}

func Help() {
	appShell.Help()
}
