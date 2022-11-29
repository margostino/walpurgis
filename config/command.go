package config

import (
	"github.com/margostino/griffin/pkg/griffin"
)

func CommandsConfiguration() *griffin.CommandsConfiguration {
	return &griffin.CommandsConfiguration{
		CommandList: Context.Configuration.Commands,
	}
}
