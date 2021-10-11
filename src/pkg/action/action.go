package action

import (
	"github.com/margostino/griffin/pkg/griffin"
)

var PowerShell *griffin.Shell

var ActionMap = map[string]func(){
	"ExecuteRankUsers": ExecuteRankUsers,
	"ExecuteExit":      griffin.ExecuteExit,
	"ExecuteHelp":      ExecuteHelp,
}

func ExecuteHelp() {
	PowerShell.Help()
}
