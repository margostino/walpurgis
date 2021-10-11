package main

import (
	"github.com/margostino/griffin/pkg/griffin"
	"github.com/margostino/walpurgis/pkg/action"
	"github.com/margostino/walpurgis/pkg/client"
)

func main() {
	client.NewClient()
	powershell := griffin.New().
		SetPrompt("walpurgis").
		SetActions(action.ActionMap).
		LoadConfiguration("../config/configuration.yml")
	action.PowerShell = powershell
	powershell.Start()
}
