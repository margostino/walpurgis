package config

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"math/rand"
	"os"
	"strings"
)

var Context *AppContext

type Data struct {
	Users []*User
}

type AppContext struct {
	Configuration *Configuration
	AppPath       string
	TwitterClient *twitter.Client
	Data          *Data
}

func Initialize() {
	appPath := getAppPath()
	configuration := GetConfiguration(appPath)
	twitterConfig := oauth1.NewConfig(configuration.Twitter.ApiKey, configuration.Twitter.ApiSecret)
	usersDataPath := appPath + "/" + configuration.Store.Users
	users := LoadUsersData(usersDataPath)
	token := oauth1.NewToken(configuration.Twitter.AccessKey, configuration.Twitter.AccessSecret)
	httpClient := twitterConfig.Client(oauth1.NoContext, token)
	Context = &AppContext{
		Configuration: configuration,
		AppPath:       appPath,
		TwitterClient: twitter.NewClient(httpClient),
		Data: &Data{
			Users: users,
		},
	}
}

func Welcome() {
	quotes := GetQuotes()
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Printf("\n")
	fmt.Printf("Welcome to Walpurgis %s!\n", GetTwitterUsername())
	fmt.Printf("\"%s\" - %s\n", quote.Quote, quote.Author)
	fmt.Printf("\n")
	fmt.Printf("Application path: %s\n", Context.AppPath)
	fmt.Printf("Configuration path: %s\n", GetConfigPath())
	fmt.Printf("Users store path: %s\n", GetUserStorePath())
	fmt.Printf("\n")
}

func GetTwitterUsername() string {
	return Context.Configuration.Twitter.Username
}

func GetUserStorePath() string {
	return Context.AppPath + "/" + Context.Configuration.Store.Users
}

func GetUsersData() []*User {
	return Context.Data.Users
}

func GetConfigPath() string {
	return Context.AppPath + "/config"
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
