package readJson

import (
	"io/ioutil"
	"encoding/json"
	"os"
)

type Settings struct {
	BotAddr    string `json"BotAddr`
	ServerAddr string `json:"ServerAddr"`
	ServerPort string `json:"ServerPort"`
}

var settings Settings = Settings{"http://127.0.0.1:3000/poll", "http://127.0.0.1", "8080"}

func InitConfigs() (s Settings, err error) {
	var file, env string
	s = settings

	env = os.Getenv("ENV")
	switch env {
	case "dev":
		file = "./configs/dev.json"
	default:
		return s, err
	}

	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return s, err
	}
	json.Unmarshal(raw, &s)

	return s, err
}