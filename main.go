package main

import (
	"github.com/truebluejason/p2e-background/internal/conf"
	"github.com/truebluejason/p2e-background/internal/poll"
	"github.com/truebluejason/p2e-background/internal/syncLog"
	"github.com/truebluejason/p2e-background/internal/timer"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if len(query["time"]) == 0 {
		syncLog.Println("[ERROR]: Time query is not set")
		w.Write([]byte("Please set a proper query string"))
		return
	}

	formattedTime := query["time"][0]
	syncLog.Println("received formattedTime: " + formattedTime)
	poll.PollBot(formattedTime)
	w.Write([]byte("Ping Received"))
}

func main() {
	timer.InitTimer()
	syncLog.Println("Running on port: " + conf.Configs.ServerPort)

	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":"+conf.Configs.ServerPort, nil); err != nil {
		panic(err)
	}
}
