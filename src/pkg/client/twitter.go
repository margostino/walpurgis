package client

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/margostino/walpurgis/pkg/config"
)

var Twitter *twitter.Client

func NewClient() {
	configuration := config.GetConfiguration()
	twitterConfig := oauth1.NewConfig(configuration.Twitter.ApiKey, configuration.Twitter.ApiSecret)
	token := oauth1.NewToken(configuration.Twitter.AccessKey, configuration.Twitter.AccessSecret)
	httpClient := twitterConfig.Client(oauth1.NoContext, token)
	Twitter = twitter.NewClient(httpClient)
}
