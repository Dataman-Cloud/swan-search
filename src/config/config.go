package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

type Config struct {
	Clusters []map[string]string `json: "clusters"`
	Ip       string              `json: "ip"`
	Port     string              `json: "port"`
	Scheme   string              `json:"scheme"`
}

func LoadConfig(configFile string) Config {
	var searchConfig Config
	log.Debugf("configfile: ", configFile)
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Errorf("Failed to read config file %s: %s", configFile, err.Error())
		return searchConfig
	}
	err = json.Unmarshal(config, &searchConfig)
	if err != nil {
		log.Errorf("Failed to unmarshal configs from configFile %s:%s", configFile, err.Error())
	}
	return searchConfig
}
