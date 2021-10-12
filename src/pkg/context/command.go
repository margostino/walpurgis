package context

import (
	"github.com/margostino/griffin/pkg/griffin"
)

func CommandsConfiguration() *griffin.CommandsConfiguration {
	return &griffin.CommandsConfiguration{
		CommandList: appContext.Configuration.Commands,
	}
}
