package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const ConfFile = "../conf.json"
const RequestString = "https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%x&boolean=true&apikey=%s"

type Config struct {
	ApiKey string `json:"api_key"`
	Port   string `json:"port"`
}

var Conf Config

func SetConfig(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	err = json.NewDecoder(file).Decode(&Conf)

	if err != nil {
		fmt.Println("error:", err)
		Conf.ApiKey = "YourApiKeyToken"
		Conf.Port = ":8181"
	}
}