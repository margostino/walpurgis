package context

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"math/rand"
	"os"
	"strings"
)

var appContext *AppContext

type AppContext struct {
	Configuration *Configuration
	AppPath       string
	TwitterClient *twitter.Client
}

func Initialize() {
	appPath := getAppPath()
	configuration := GetConfiguration(appPath + "/config/")
	twitterConfig := oauth1.NewConfig(configuration.Twitter.ApiKey, configuration.Twitter.ApiSecret)
	token := oauth1.NewToken(configuration.Twitter.AccessKey, configuration.Twitter.AccessSecret)
	httpClient := twitterConfig.Client(oauth1.NoContext, token)
	appContext = &AppContext{
		Configuration: configuration,
		AppPath:       appPath,
		TwitterClient: twitter.NewClient(httpClient),
	}
}

func Welcome() {
	quotes := GetQuotes()
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Printf("\n")
	fmt.Printf("Welcome to Walpurgis %s!\n", GetTwitterUsername())
	fmt.Printf("\"%s\" - %s\n", quote.Quote, quote.Author)
	fmt.Printf("\n")
	fmt.Printf("Application path: %s\n", appContext.AppPath)
	fmt.Printf("Configuration path: %s\n", GetConfigPath())
	fmt.Printf("Users store path: %s\n", GetUserStorePath())
	fmt.Printf("\n")
}

func GetTwitterUsername() string {
	return appContext.Configuration.Twitter.Username
}

func GetUserStorePath() string {
	return appContext.AppPath + "/" + appContext.Configuration.Store.Users
}

func GetConfigPath() string {
	return appContext.AppPath + "/config"
}

func getAppPath() string {
	appPath := os.Getenv("WALPURGIS_PATH")
	if appPath == "" {
		appPath = ".."
	} else if strings.HasSuffix(appPath, "/") {
		appPath = appPath[:len(appPath)-len("/")]
	}
	return appPath
}
