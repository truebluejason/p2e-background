package conf

import (
	"io/ioutil"
	"encoding/json"
	"os"
)

type Settings struct {
	BotURL    	string `json"BotURL`
	DBName    	string `json"DBName`
	DBPassword  string `json"DBPassword`
	DBUser    	string `json"DBUser`
	ServerAddr 	string `json:"ServerAddr"`
	ServerPort 	string `json:"ServerPort"`
}

var defaultConfigs Settings = Settings{
	"http://127.0.0.1:3000/poll",
	"P2E",
	"path2enlightenment",
	"root",
	"http://127.0.0.1",
	"8080",
}

var Configs Settings = initConfigs()

func initConfigs() (s Settings) {
	var file, env string
	s = defaultConfigs

	env = os.Getenv("ENV")
	switch env {
	case "dev":
		file = "./configs/dev.json"
	default:
		return s
	}

	raw, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(raw, &s)

	return s
}