package main

import (
	"github.com/margostino/walpurgis/pkg/context"
	"github.com/margostino/walpurgis/pkg/shell"
)

func main() {
	context.Initialize()
	shell.Initialize()
	context.Welcome()
	shell.Start()
}
