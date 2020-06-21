package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Img       string   `json:"img"`
	Envs      []string `json:"envs"`
	RunOption string   `json:"runOption"`
}

func parseConfigFile(path string) (Config, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var c Config
	if err := json.Unmarshal(raw, &c); err != nil {
		return Config{}, err
	}

	return c, nil
}
