package poller

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"time"
)

func InitPoller(port string) {
	go initPoll(port)
}

func initPoll(port string) {
	var now time.Time
	var formatted, recorded string

	for true {
		now = time.Now()
		// Legitimate Counter
		// formatted = strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute())
		formatted = strconv.Itoa(now.Minute()) + ":" + strconv.Itoa(now.Second())
		if recorded != formatted {
			url := "http://127.0.0.1:" + port + "/ping"

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Reason - " + err.Error())
				pause(5)
				continue
			}

			defer resp.Body.Close()

			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil || string(contents) != "Ping Received" {
				fmt.Println("Reason - " + err.Error())
				pause(5)
				continue
			}

			recorded = formatted
		}
	}
}

func pause(sec int) {
	fmt.Println("Pausing polling for " + strconv.Itoa(sec) + " seconds.")
	time.Sleep(time.Duration(sec) * time.Second)
}