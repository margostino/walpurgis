package shell

import (
	"github.com/margostino/griffin/pkg/griffin"
	"github.com/margostino/walpurgis/pkg/context"
)

var appShell *griffin.Shell

func Initialize() {
	appShell = griffin.New().
		SetPrompt(context.GetUsername() + "@walpurgis").
		SetSimpleActions(SimpleActionsMapping).
		SetMultiParamsActions(InputStringsActionMapping).
		SetConfiguration(context.CommandsConfiguration())
}

func Start() {
	appShell.Start()
}

func Help() {
	appShell.Help()
}
