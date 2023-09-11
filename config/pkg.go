package config

import (
	"github.com/tkanos/gonfig"
)

var (
	Config = Configuration{}
)

func init() {
	Config = GetConfig()
}

func GetConfig(params ...string) (configuration Configuration) {
	configFile := "config/config.json"
	if len(params) > 0 {
		configFile = params[0]
	}
	gonfig.GetConf(configFile, &configuration)
	return
}
