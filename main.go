package main

import (
	"github.com/margostino/walpurgis/config"
	"github.com/margostino/walpurgis/shell"
)

func main() {
	config.Initialize()
	shell.Initialize()
	config.Welcome()
	shell.Start()
}
