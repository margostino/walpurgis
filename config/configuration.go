package config

import (
	"github.com/margostino/griffin/pkg/griffin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Twitter struct {
	Username     string `yaml:"username"`
	ApiKey       string `yaml:"api_key"`
	ApiSecret    string `yaml:"api_secret"`
	AccessKey    string `yaml:"access_key"`
	AccessSecret string `yaml:"access_secret"`
}

type Store struct {
	Users string `yaml:"users"`
}

type Quote struct {
	Author string `yaml:"author"`
	Quote  string `yaml:"quote"`
}

type Configuration struct {
	Store    Store                          `yaml:"store"`
	Twitter  Twitter                        `yaml:"social"`
	Commands []griffin.CommandConfiguration `yaml:"commands"`
	Quotes   []Quote                        `yaml:"quotes"`
}

func GetConfiguration(appPath string) *Configuration {
	var configuration Configuration
	yamlFile, err := ioutil.ReadFile(appPath + "/config.yml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	
	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))
	err = yaml.Unmarshal(yamlFile, &configuration)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &configuration
}

func GetUsername() string {
	return Context.Configuration.Twitter.Username
}

func GetQuotes() []Quote {
	return Context.Configuration.Quotes
}
