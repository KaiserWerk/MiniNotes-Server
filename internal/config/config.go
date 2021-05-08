package config

import (
	"github.com/KaiserWerk/mininotes-server/internal/entity"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	configFile = "config.yml"
)

func SetConfigFile(s string) {
	configFile = s
}

func GetConfigFile() string {
	return configFile
}

func GetConfig() *entity.UserConfig {
	cont, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic("Could not read config file: " + err.Error())
	}

	var cfg entity.UserConfig
	err = yaml.Unmarshal(cont, &cfg)
	if err != nil {
		panic("Could not parse config file: " + err.Error())
	}

	return &cfg
}
