package timer

import (
	"net/http"
	"io/ioutil"
	"strconv"
	"../syncLog"
	"time"
)

func InitTimer(port string) {
	go initTiming(port)
}

func initTiming(port string) {
	var now time.Time
	var formatted, recorded string

	for true {
		now = time.Now()
		formatted = strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute())

		if recorded != formatted {
			url := "http://127.0.0.1:" + port + "/ping?time=" + formatted
			syncLog.Println("Pinging " + url)

			resp, err := http.Get(url)
			if err != nil {
				syncLog.Println("[ERROR]: " + err.Error())
				pause(20)
				continue
			}

			defer resp.Body.Close()
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil || string(contents) != "Ping Received" {
				syncLog.Println("[ERROR]: Improper response received from ping")
				pause(20)
				continue
			}

			recorded = formatted
		}
		pause(20)
	}
}

func pause(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}