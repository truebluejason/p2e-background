package conf

import (
	"io/ioutil"
	"encoding/json"
	"os"
)

type Settings struct {
	BotURL    	string `json"BotURL`
	DBHost    	string `json"DBHost`
	DBName    	string `json"DBName`
	DBPassword  string `json"DBPassword`
	DBUser    	string `json"DBUser`
	ServerAddr 	string `json:"ServerAddr"`
	ServerPort 	string `json:"ServerPort"`
}

var Configs Settings = initConfigs()

func initConfigs() (s Settings) {
	var file, env string

	env = os.Getenv("ENV")
	switch env {
	case "dev":
		file = "./config/dev.json"
	default:
		file = "/p2e-background/config/production.json"
	}

	raw, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(raw, &s)

	return s
}