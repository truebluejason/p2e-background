package main

import (
	"fmt"
  "net/http"
  "./internal/poller"
  "./internal/readConf"
	//"database/sql"
)
//import _ "github.com/go-sql-driver/mysql"



func ping(w http.ResponseWriter, r *http.Request) {
  message := "Ping Received"
  w.Write([]byte(message))
}

func main() {
  settings, err := readConf.InitConfigs()
  if err != nil {
    panic(err)
  }

  poller.InitPoller(settings.ServerPort)

  fmt.Println("Running on port: " + settings.ServerPort)

  http.HandleFunc("/ping", ping)
  if err := http.ListenAndServe(":" + settings.ServerPort, nil); err != nil {
    panic(err)
  }
}