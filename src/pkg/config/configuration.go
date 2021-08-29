package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Twitter struct {
	ApiKey       string `yaml:"api_key"`
	ApiSecret    string `yaml:"api_secret"`
	AccessKey    string `yaml:"access_key"`
	AccessSecret string `yaml:"access_secret"`
}

type Configuration struct {
	Twitter Twitter `yaml:"twitter"`
}

func GetConfigFile(filename string) string {
	var filePath string
	var configPath = os.Getenv("WALPURGIS_CONFIG_PATH")

	if configPath == "" {
		configPath = "../config/"
	}
	if strings.HasSuffix(configPath, "/") {
		filePath = configPath + filename
	} else {
		filePath = configPath + "/" + filename
	}
	return filePath
}

func GetConfiguration() *Configuration {
	var configuration Configuration
	yamlFile, err := ioutil.ReadFile(GetConfigFile("configuration.yaml"))

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