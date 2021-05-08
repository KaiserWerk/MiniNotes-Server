package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"mininotes-server/entity"
)

var (
	configFile = "config.yml"
)

func SetConfigFile(s string) {
	configFile = s
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
