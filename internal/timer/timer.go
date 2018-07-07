package timer

import (
	"github.com/truebluejason/p2e-background/internal/conf"
	"net/http"
	"io/ioutil"
	"strconv"
	"github.com/truebluejason/p2e-background/internal/syncLog"
	"time"
)

func InitTimer() {
	go initTiming()
}

func initTiming() {
	var now time.Time
	var formatted, recorded string
	var port string = conf.Configs.ServerPort

	for true {
		now = time.Now()

		var minute int = now.Minute()
		var minuteStr string
		if minute < 10 {
			minuteStr = "0" + strconv.Itoa(minute)
		} else {
			minuteStr = strconv.Itoa(minute)
		}

		formatted = strconv.Itoa(now.Hour()) + ":" + minuteStr

		if recorded != formatted {
			url := "http://127.0.0.1:" + port + "/ping?time=" + formatted
			syncLog.Println("Pinging " + url)

			resp, err := http.Get(url)
			if err != nil {
				syncLog.Println("[ERROR]: " + err.Error())
				pause(10)
				continue
			}

			defer resp.Body.Close()
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil || string(contents) != "Ping Received" {
				syncLog.Println("[ERROR]: Improper response received from ping")
				pause(10)
				continue
			}

			recorded = formatted
		}
		pause(10)
	}
}

func pause(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}