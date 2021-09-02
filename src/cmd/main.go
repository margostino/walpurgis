package main

import (
	"github.com/margostino/walpurgis/pkg/client"
	"github.com/margostino/walpurgis/pkg/commands"
	"github.com/margostino/walpurgis/pkg/shell"
)

func main() {
	var plan string
	client.NewClient()
	powershell := shell.NewShell()

	for {
		plan = powershell.Input()
		if plan == "rank creators" {
			commands.RankCreators()
		}
	}

}
