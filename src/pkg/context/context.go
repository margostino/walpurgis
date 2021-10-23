package context

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/margostino/walpurgis/pkg/config"
)

var appContext *AppContext

type AppContext struct {
	Configuration *config.Configuration
	Twitter       *twitter.Client
}

func Initialize() {
	configuration := config.GetConfiguration()
	twitterConfig := oauth1.NewConfig(configuration.Twitter.ApiKey, configuration.Twitter.ApiSecret)
	token := oauth1.NewToken(configuration.Twitter.AccessKey, configuration.Twitter.AccessSecret)
	httpClient := twitterConfig.Client(oauth1.NoContext, token)
	appContext = &AppContext{
		Configuration: configuration,
		Twitter:       twitter.NewClient(httpClient),
	}
}

func GetTwitterUsername() string {
	return appContext.Configuration.Twitter.Username
}
