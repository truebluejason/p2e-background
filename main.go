package main

import (
  "net/http"
  "./internal/poll"
  "./internal/readJson"
  "./internal/syncLog"
  "./internal/timer"
)

settings, err := readJson.InitConfigs()
if err != nil {
  panic(err)
}

func ping(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  if len(query["time"]) == 0 {
    syncLog.Println("[ERROR]: Time query is not set")
    w.Write([]byte("Please set a proper query string"))
    return
  }

  formattedTime := query["time"][0]
  syncLog.Println("formattedTime is: " + formattedTime)
  poll.PollBot(formattedTime, settings.BotAddr)
  w.Write([]byte("Ping Received"))
}

func main() {
  timer.InitTimer(settings.ServerPort)
  syncLog.Println("Running on port: " + settings.ServerPort)

  http.HandleFunc("/ping", ping)
  if err := http.ListenAndServe(":" + settings.ServerPort, nil); err != nil {
    panic(err)
  }
}