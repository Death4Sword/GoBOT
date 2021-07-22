package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config files ...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Printf("Failed to marshal config: %s\n", err.Error())
		return err
	}

	fmt.Println(string(file))

	json.Unmarshal(file, &config)

	if err != nil {
		fmt.Printf("Failed to marshal config: %s\n", err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
